package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// JobConfig holds the configuration for a Prometheus job
type JobConfig struct {
	JobName      string   `yaml:"job_name"`
	StaticConfig []Static `yaml:"static_configs"`
}

// Static holds static_configs for a Prometheus job
type Static struct {
	Targets []string          `yaml:"targets"`
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

// updateConfigMap updates the specified ConfigMap in the given namespace
func updateConfigMap(clientset *kubernetes.Clientset, namespace, configMapName string, data map[string]string) error {
	_, err := clientset.CoreV1().ConfigMaps(namespace).Update(context.TODO(), &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
		Data: data,
	}, metav1.UpdateOptions{})
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

// updateJobTargets updates the targets for a specific job in the Prometheus configuration
func updateJobTargets(config *PrometheusConfig, jobName string, newTargets []string) bool {
	for i, job := range config.ScrapeConfigs {
		if job.JobName == jobName {
			config.ScrapeConfigs[i].StaticConfig = []Static{
				{Targets: newTargets},
			}
			return true
		}
	}
	return false
}

func main() {
	// Configuration
	configMapName := os.Getenv("CONFIG_MAP_NAME")
	if configMapName == "" {
		configMapName = "prometheus-config"
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

	log.Printf("to retrieve ConfigMap: %s", configMapName)
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
		err = yaml.Unmarshal([]byte(configMap.Data["prometheus.yml"]), &prometheusConfig)
		if err != nil {
			log.Printf("Failed to parse Prometheus configuration: %v", err)
			time.Sleep(interval)
			continue
		}

		// Update targets for each job
		for _, job := range prometheusConfig.ScrapeConfigs {
			labelSelector := job.StaticConfig[0].Labels["app"]
			// log.Printf("labelSelector: app=%s",labelSelector)
			podIPs, err := getPodIPs(clientset, namespace, fmt.Sprintf("app=%s", labelSelector))
			if err != nil {
				log.Printf("Failed to retrieve pods: %v", err)
				continue
			}
			newTargets := []string{}
			for _, ip := range podIPs {
				newTargets = append(newTargets, fmt.Sprintf("%s:9100", ip))
			}
			if updateJobTargets(&prometheusConfig, job.JobName, newTargets) {
				log.Printf("Updated targets for job [%s] successfully", job.JobName)
			}
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
			"prometheus.yml": string(newConfigData),
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
