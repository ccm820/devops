from kubernetes import client, config
import yaml
import json
import time


def get_prometheus_configmap(namespace, configmap_name):
    """
    获取 Prometheus 的 ConfigMap 内容。
    """
    config.load_kube_config()
    v1 = client.CoreV1Api()
    try:
        configmap = v1.read_namespaced_config_map(configmap_name, namespace)
        return configmap
    except Exception as e:
        print(f"获取 ConfigMap 失败: {e}")
        return None


def update_prometheus_configmap(namespace, configmap_name, new_data):
    """
    更新 Prometheus 的 ConfigMap 内容。
    """
    config.load_kube_config()
    v1 = client.CoreV1Api()
    try:
        # 更新 ConfigMap 数据
        body = client.V1ConfigMap(data=new_data)
        v1.patch_namespaced_config_map(configmap_name, namespace, body)
        print("成功更新 ConfigMap!")
    except Exception as e:
        print(f"更新 ConfigMap 失败: {e}")


def fetch_pod_ips(namespace, label_selector):
    """
    获取指定 namespace 和 label_selector 的 Pod IP 列表。
    """
    config.load_kube_config()
    v1 = client.CoreV1Api()
    pod_ips = []
    try:
        pods = v1.list_namespaced_pod(namespace=namespace, label_selector=label_selector)
        for pod in pods.items:
            if pod.status.pod_ip and pod.status.phase == "Running":
                pod_ips.append(pod.status.pod_ip)
    except Exception as e:
        print(f"获取 Pod 信息失败: {e}")
    return pod_ips


def update_job_targets(config_data, job_name, new_targets):
    """
    更新 Prometheus 配置中指定 job_name 的 targets。
    """
    config_dict = yaml.safe_load(config_data)
    updated = False

    for scrape_config in config_dict.get("scrape_configs", []):
        if scrape_config.get("job_name") == job_name:
            scrape_config["static_configs"] = [{"targets": new_targets}]
            updated = True

    if updated:
        return yaml.dump(config_dict)
    else:
        print(f"未找到 job_name: {job_name}")
        return None


def extract_jobs_from_config(config_data):
    """
    从 prometheus.yml 配置文件中提取所有的 job 信息。
    """
    config_dict = yaml.safe_load(config_data)
    jobs = []

    for scrape_config in config_dict.get("scrape_configs", []):
        job = {}
        job["job_name"] = scrape_config.get("job_name")
        if "static_configs" in scrape_config:
            for config in scrape_config["static_configs"]:
                if "targets" in config:
                    job["targets"] = config["targets"]
                    # 如果有指定 label_selector，提取出来
                    if "labels" in config:
                        job["label_selector"] = config["labels"]
                    else:
                        job["label_selector"] = None
        jobs.append(job)

    return jobs


def monitor_and_update(namespace, configmap_name, interval):
    """
    监控 Pod IP 并动态更新 Prometheus 的 ConfigMap。
    """
    while True:
        try:
            # 获取当前 ConfigMap 数据
            configmap = get_prometheus_configmap(namespace, configmap_name)
            if not configmap or not configmap.data or "prometheus.yml" not in configmap.data:
                print("未找到 Prometheus 配置文件")
                continue

            # 加载 Prometheus 配置
            prometheus_config = configmap.data["prometheus.yml"]

            # 提取所有的 job 配置信息
            jobs = extract_jobs_from_config(prometheus_config)

            for job in jobs:
                job_name = job['job_name']
                label_selector = job['label_selector']
                if label_selector:
                    # 获取 Pod IP
                    pod_ips = fetch_pod_ips(namespace, label_selector)
                    new_targets = [f"{ip}:9100" for ip in pod_ips]  # 默认端口为9100

                    # 更新 Prometheus 配置
                    updated_config = update_job_targets(prometheus_config, job_name, new_targets)
                    if updated_config:
                        configmap.data["prometheus.yml"] = updated_config

            # 更新 ConfigMap
            update_prometheus_configmap(namespace, configmap_name, configmap.data)

        except Exception as e:
            print(f"监控和更新任务失败: {e}")
        time.sleep(interval)


if __name__ == "__main__":
    # 配置参数
    NAMESPACE = "monitoring"  # Prometheus ConfigMap 所在的命名空间
    CONFIGMAP_NAME = "prometheus-config"  # Prometheus ConfigMap 名称
    UPDATE_INTERVAL = 30  # 更新间隔（秒）

    # 启动监控和更新任务
    monitor_and_update(NAMESPACE, CONFIGMAP_NAME, UPDATE_INTERVAL)
