#### Note

Container probes: kubelet is provide some methods for us to check container health. LivenessProbe and ReadinessProbe

##### LivenessProbe
It is used for checking container health, if not it will restart the container.

##### ReadinessProbe
It is difference with LivenessProbe. If it sees a container is not healthy, it will remove this container from network.
For example: we have 3 pods, one of them crashes and we send some requests to server. Service will forward the request to random pod whether 
pod is crashing. If we have readinessProbe, it will remove the pod from service, so the pod crashing will never receive any request.