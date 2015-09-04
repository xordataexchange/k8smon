package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/cactus/go-statsd-client/statsd"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	kSelector "k8s.io/kubernetes/pkg/labels"
)

var k8sclient *unversioned.Client
var statsdclient statsd.Statter

func main() {
	appname := path.Base(os.Args[0])
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"applicationname": appname,
	})
	var k8shost, k8sport, k8sproto string
	var statsdhost, statsdport string

	if k8shost = os.Getenv("KUBERNETES_SERVICE_HOST"); k8shost == "" {
		logger.Panic("Must Supply KUBERNETES_SERVICE_HOST")
	}
	if k8sport = os.Getenv("KUBERNETES_SERVICE_PORT"); k8sport == "" {
		logger.Panic("Must Supply KUBERNETES_SERVICE_PORT")
	}
	if k8sproto = os.Getenv("KUBERNETES_SERVICE_PROTO"); k8sproto == "" {
		logger.Panic("Must Supply KUBERNETES_SERVICE_PROTO (http/https)")
	}
	if statsdhost = os.Getenv("STATSD_SERVICE_HOST"); statsdhost == "" {
		logger.Panic("Must Supply STATSD_SERVICE_HOST")
	}
	if statsdport = os.Getenv("STATSD_SERVICE_PORT"); statsdport == "" {
		logger.Panic("Must Supply STATSD_SERVICE_PORT")
	}

	k8s := fmt.Sprintf("%s://%s:%s", k8sproto, k8shost, k8sport)
	logger.Infof("Connecting to Kubernetes Master: %s", k8s)
	config := unversioned.Config{
		Host:    k8s,
		Version: "v1",
	}
	var err error

	k8sclient, err = unversioned.New(&config)
	if err != nil {
		logger.Panic("Unable to connect to Kubernetes Master", err)
	}

	sd := fmt.Sprintf("%s:%s", statsdhost, statsdport)
	logger.Infof("Connecting to statsd: %s", sd)
	statsdclient, err = statsd.NewClient(sd, "k8smon")
	if err != nil {
		logger.Panic("Unable to connect to Kubernetes Master", err)
	}
	for {
		CountInstances()
		time.Sleep(time.Second * 5)
	}
}

func CountInstances() {
	list, _ := k8sclient.ReplicationControllers(api.NamespaceDefault).List(kSelector.Everything())
	for _, i := range list.Items {
		statsdclient.Gauge(i.ObjectMeta.Name, int64(i.Status.Replicas), 1.0)
	}

}
