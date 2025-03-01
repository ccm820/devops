package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"path/filepath"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// JobConfig holds the configuration for a Prometheus job
type JobConfig struct {
	JobName      string   `yaml:"job_name"`
	StaticConfig []Static `yaml:"static_configs"`
}

// Static holds static_configs for a Prometheus job
type Static struct {
	Targets []string          `yaml:"targets,flow"` // Use flow style for single-line output
	Labels  map[string]string `yaml:"labels,omitempty"`
}

// PrometheusConfig represents the entire Prometheus configuration
type PrometheusConfig struct {
	ScrapeConfigs []JobConfig `yaml:"scrape_configs"`
}

// getClientSet initializes a Kubernetes client
func getClientSet() (*kubernetes.Clientset, error) {
	// If running inside a cluster, use the in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		// Otherwise, use the kubeconfig file from the local filesystem
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("unable to load kubeconfig: %v", err)
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("unable to create Kubernetes client: %v", err)
	}
	return clientset, nil
}

// getNamespace retrieves the namespace from an environment variable or falls back to the namespace of the current pod
func getNamespace() (string, error) {
	if namespace := os.Getenv("NAMESPACE"); namespace != "" {
		log.Printf("Environment variable NAMESPACE is set: %s",namespace)
		return namespace, nil
	} 
	// Attempt to read the namespace of the current pod
	namespacePath := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	data, err := os.ReadFile(namespacePath)
	if err != nil {
		return "", fmt.Errorf("failed to read namespace file: %v", err)
	}
	return string(data), nil
}

// getConfigMap retrieves the specified ConfigMap from the given namespace
func getConfigMap(clientset *kubernetes.Clientset, namespace, configMapName string) (*v1.ConfigMap, error) {
	configMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ConfigMap: %v", err)
	}
	return configMap, nil
}

// updateConfigMap updates the specified ConfigMap in the given namespace, preserving non-updated fields
func updateConfigMap(clientset *kubernetes.Clientset, namespace, configMapName string, updatedData map[string]string) error {
	// Retrieve the existing ConfigMap
	configMap, err := getConfigMap(clientset, namespace, configMapName)
	if err != nil {
		return fmt.Errorf("failed to retrieve ConfigMap for updating: %v", err)
	}

	// Preserve original fields and update only the specified ones
	for key, value := range updatedData {
		configMap.Data[key] = value
	}

	// Apply the updated ConfigMap
	_, err = clientset.CoreV1().ConfigMaps(namespace).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update ConfigMap: %v", err)
	}
	return nil
}

// getPodIPs retrieves the IPs of running pods matching the specified label selector in the given namespace
func getPodIPs(clientset *kubernetes.Clientset, namespace, labelSelector string) ([]string, error) {
	log.Printf("retrieves the IPs of running pods matching the specified label selector %s", labelSelector)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve pods: %v", err)
	}
	var podIPs []string
	for _, pod := range pods.Items {
		if pod.Status.PodIP != "" && pod.Status.Phase == v1.PodRunning {
			podIPs = append(podIPs, pod.Status.PodIP)
		}
	}
	return podIPs, nil
}

// hasRequiredLabels checks if the given job's static config contains the required labels
func hasRequiredLabels(staticConfigs []Static, requiredLabels map[string]string) bool {
	for _, staticConfig := range staticConfigs {
		matches := true
		for key, value := range requiredLabels {
			if staticConfig.Labels[key] != value {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

// updateJobTargetsPreservingFieldsWithPorts updates only the IPs in the targets of a specific job,
// preserving other fields and labels, and returns true if any change was made.
func updateJobTargetsPreservingFieldsWithPorts(config *PrometheusConfig, jobName string, newIPs []string) bool {
	updated := false
	for i, job := range config.ScrapeConfigs {
		if job.JobName == jobName {
			if len(job.StaticConfig) == 0 {
				continue
			}

			existingTargets := job.StaticConfig[0].Targets
			port := extractPortFromTarget(existingTargets)
			if port == "" {
				log.Printf("Failed to extract port for job [%s], skipping update", jobName)
				continue
			}

			// Generate new targets with the same port
			newTargets := generateTargetsWithPort(newIPs, port)

			// Compare existing targets with new targets
			if !isTargetsEqual(existingTargets, newTargets) {
				log.Printf("Job [%s]: targets IPs changed from %s to %s", jobName,existingTargets,newTargets)
				config.ScrapeConfigs[i].StaticConfig[0].Targets = newTargets
				updated = true
			} else {
				log.Printf("Job [%s]: No IPs changed for targets %s", jobName,existingTargets)
			}
		}
	}
	return updated
}

// extractPortFromTarget extracts the port from an existing target (e.g., "192.168.1.1:9090").
func extractPortFromTarget(targets []string) string {
	if len(targets) > 0 {
		parts := strings.Split(targets[0], ":")
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return ""
}

// generateTargetsWithPort appends a fixed port to a list of IPs to generate new targets.
func generateTargetsWithPort(ips []string, port string) []string {
	var targets []string
	for _, ip := range ips {
		targets = append(targets, fmt.Sprintf("%s:%s", ip, port))
	}
	return targets
}

// isTargetsEqual compares two target lists for equality.
func isTargetsEqual(existing, new []string) bool {
	if len(existing) != len(new) {
		return false
	}

	existingMap := make(map[string]struct{}, len(existing))
	for _, target := range existing {
		existingMap[target] = struct{}{}
	}

	for _, target := range new {
		if _, exists := existingMap[target]; !exists {
			return false
		}
	}
	return true
}

func main() {
	// Configuration
	configMapName := os.Getenv("CONFIG_MAP_NAME")
	if configMapName == "" {
		configMapName = "prometheus-config"
	}

	configMapDataFile := os.Getenv("CONFIG_MAP_DATA_FILE")
	if configMapDataFile == "" {
		configMapDataFile = "prometheus.yml"
	}

	interval := 30 * time.Second

	clientset, err := getClientSet()
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}

	namespace, err := getNamespace()
	if err != nil {
		log.Fatalf("Failed to determine namespace: %v", err)
	}

	// Define the required labels for updating a job
	requiredLabels := map[string]string{
		"monitor": "true", // Replace "example" with your actual label value
	}
	
	for {
		// Retrieve the ConfigMap
		configMap, err := getConfigMap(clientset, namespace, configMapName)
		if err != nil {
			log.Printf("Failed to retrieve ConfigMap: %v", err)
			time.Sleep(interval)
			continue
		}

		// Parse Prometheus configuration
		prometheusConfig := PrometheusConfig{}
		
		err = yaml.Unmarshal([]byte(configMap.Data[configMapDataFile]), &prometheusConfig)
		if err != nil {
			log.Printf("Failed to parse Prometheus configuration: %v", err)
			time.Sleep(interval)
			continue
		}

		// Update targets for jobs with required labels
		updated := false
		updated_jobs := 0	
		// Update targets for each job
		for _, job := range prometheusConfig.ScrapeConfigs {
			if hasRequiredLabels(job.StaticConfig, requiredLabels) {
				labelSelector := job.StaticConfig[0].Labels["app"]
				podIPs, err := getPodIPs(clientset, namespace, fmt.Sprintf("app=%s", labelSelector))
				if err != nil {
					log.Printf("Failed to retrieve pods: %v", err)
					continue
				}
	            updated_jobs = updated_jobs + 1
				if updateJobTargetsPreservingFieldsWithPorts(&prometheusConfig, job.JobName, podIPs) {
					log.Printf("Updated targets for job [%s] with IPs %s successfully", job.JobName, podIPs)
					updated = true
				}
			}
		}
		if updated_jobs == 0 {
			log.Println("No jobs met the required label criteria skipping ConfigMap update")
			time.Sleep(interval)
			continue
		} 
		if !updated {
			log.Printf("No changes for %d jobs which met the required label criteria", updated_jobs)
			time.Sleep(interval)
			continue
		}
		// Serialize updated configuration
		newConfigData, err := yaml.Marshal(&prometheusConfig)
		if err != nil {
			log.Printf("Failed to serialize Prometheus configuration: %v", err)
			time.Sleep(interval)
			continue
		}

		// Update the ConfigMap
		err = updateConfigMap(clientset, namespace, configMapName, map[string]string{
			configMapDataFile: string(newConfigData),
		})
		if err != nil {
			log.Printf("Failed to update ConfigMap: %v", err)
			time.Sleep(interval)
			continue
		}

		log.Println("Successfully updated Prometheus configuration")
		time.Sleep(interval)
	}
}
