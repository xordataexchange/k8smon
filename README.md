# k8smon
k8smon is a simple kubernetes to statsd bridge.  It reads the Kubernetes API and publishes the number of running pods for each replication controller it finds every five seconds.  We built it so we could monitor the number of running pods for each ReplicationController and trigger alerts (elsewhere) when that number was lower than our threshold.


## Configuration
k8smon expects a few environment variables:
```
STATSD_SERVICE_HOST=127.0.0.1 
STATSD_SERVICE_PORT=8125 
KUBERNETES_SERVICE_HOST=10.0.10.10
KUBERNETES_SERVICE_PORT=8080 
KUBERNETES_SERVICE_PROTO=http 
STATSD_PREFIX=k8smon
```
## Running
Standalone:
```
STATSD_SERVICE_HOST=127.0.0.1 STATSD_SERVICE_PORT=8125 KUBERNETES_SERVICE_HOST=10.0.10.10 KUBERNETES_SERVICE_PORT=8080 KUBERNETES_SERVICE_PROTO=http /path/to/k8smon
```
Docker:
```
docker run -d -e STATSD_SERVICE_HOST=127.0.0.1 -e STATSD_SERVICE_PORT=8125 -e KUBERNETES_SERVICE_HOST=10.0.10.10 -e KUBERNETES_SERVICE_PORT=8080 -e KUBERNETES_SERVICE_PROTO=http bketelsen/k8smon
```
## Building
k8smon requires [gb](http://getgb.io)

Clone the repository, and type `gb build`



