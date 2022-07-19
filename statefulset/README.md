##### Note:
In order to interact with pod (pod of statefulset), we need to create a headless service (ClusterIP: None).
The reason means when we create a service with ClusterIP, k8s will create a virtual IP and DNS stand before pods.
So the pod will be called random by LB. So when we set `ClusterIP: None`, we can call the pod directly by dns of this pod.