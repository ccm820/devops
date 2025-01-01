package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// JobConfig represents job configuration in Prometheus
type JobConfig struct {
	JobName      string   `yaml:"job_name"`
	StaticConfig []Static `yaml:"static_configs"`
}

// Static represents static_configs in Prometheus configuration
type Static struct {
	Targets []string          `yaml:"targets"`
	Labels  map[string]string `yaml:"labels,omitempty"`
}

// PrometheusConfig represents the complete Prometheus configuration
type PrometheusConfig struct {
	ScrapeConfigs []JobConfig `yaml:"scrape_configs"`
}

// getClientSet initializes and returns a Kubernetes clientset
func getClientSet() (*kubernetes.Clientset, error) {
	// Use in-cluster configuration if available, otherwise fall back to kubeconfig
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load kubeconfig: %v", err)
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}
	return clientset, nil
}

// getConfigMap retrieves the specified ConfigMap from Kubernetes
func getConfigMap(clientset *kubernetes.Clientset, namespace, configMapName string) (*v1.ConfigMap, error) {
	configMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ConfigMap: %v", err)
	}
	return configMap, nil
}

// updateConfigMap updates the specified ConfigMap with new data
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

// getPodIPs retrieves the IPs of running Pods based on the provided label selector
func getPodIPs(clientset *kubernetes.Clientset, namespace, labelSelector string) ([]string, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Pods: %v", err)
	}
	var podIPs []string
	for _, pod := range pods.Items {
		if pod.Status.PodIP != "" && pod.Status.Phase == v1.PodRunning {
			podIPs = append(podIPs, pod.Status.PodIP)
		}
	}
	return podIPs, nil
}

// updateJobTargets updates the targets of the specified job in the Prometheus configuration
func updateJobTargets(config *PrometheusConfig, jobName string, newTargets []string) bool {
	for i, job := range config.ScrapeConfigs {
		if job.JobName == jobName {
			if len(job.StaticConfig) > 0 && reflect.DeepEqual(job.StaticConfig[0].Targets, newTargets) {
				log.Printf("Targets for job [%s] have not changed, skipping update.", jobName)
				return false
			}
			config.ScrapeConfigs[i].StaticConfig = []Static{
				{Targets: newTargets},
			}
			return true
		}
	}
	return false
}

func main() {
	configMapName := os.Getenv("CONFIGMAP_NAME")
	if configMapName == "" {
		log.Fatalf("CONFIGMAP_NAME environment variable is not set.")
	}

	interval := 30 * time.Second

	clientset, err := getClientSet()
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}

	// Retrieve namespace and ConfigMap name from environment variables
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		log.Printf("NAMESPACE environment variable is not set.")
	} else {
		namespace, err := getNamespace()
		if err != nil {
			log.Fatalf("Failed to determine namespace: %v", err)
		} else {
			log.Printf("Current running NAMESPACE is : %s", namespace)
		}
	}

	for {
		// Retrieve the ConfigMap
		configMap, err := getConfigMap(clientset, namespace, configMapName)
		if err != nil {
			log.Printf("Failed to get ConfigMap: %v", err)
			time.Sleep(interval)
			continue
		}

		// Parse the Prometheus configuration
		prometheusConfig := PrometheusConfig{}
		err = yaml.Unmarshal([]byte(configMap.Data["prometheus.yml"]), &prometheusConfig)
		if err != nil {
			log.Printf("Failed to parse Prometheus configuration: %v", err)
			time.Sleep(interval)
			continue
		}

		// Update targets for each job
		for _, job := range prometheusConfig.ScrapeConfigs {
			if len(job.StaticConfig) == 0 || len(job.StaticConfig[0].Labels) == 0 {
				log.Printf("Job [%s] has no static_configs or labels, skipping...", job.JobName)
				continue
			}

			labelSelector := job.StaticConfig[0].Labels["app"]
			podIPs, err := getPodIPs(clientset, namespace, fmt.Sprintf("app=%s", labelSelector))
			if err != nil {
				log.Printf("Failed to get Pods for job [%s]: %v", job.JobName, err)
				continue
			}

			var newTargets []string
			for _, ip := range podIPs {
				newTargets = append(newTargets, fmt.Sprintf("%s:9100", ip))
			}

			if updateJobTargets(&prometheusConfig, job.JobName, newTargets) {
				log.Printf("Successfully updated targets for job [%s]", job.JobName)
			}
		}

		// Serialize the updated configuration
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

		log.Println("Successfully updated Prometheus configuration.")
		time.Sleep(interval)
	}
}
