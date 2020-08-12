package main

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"time"
//
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/rest"
//)
//
//// PodMetricsList : PodMetricsList
//type PodMetricsList struct {
//	Kind       string `json:"kind"`
//	APIVersion string `json:"apiVersion"`
//	Metadata   struct {
//		SelfLink string `json:"selfLink"`
//	} `json:"metadata"`
//	Items []struct {
//		Metadata struct {
//			Name              string    `json:"name"`
//			Namespace         string    `json:"namespace"`
//			SelfLink          string    `json:"selfLink"`
//			CreationTimestamp time.Time `json:"creationTimestamp"`
//		} `json:"metadata"`
//		Timestamp  time.Time `json:"timestamp"`
//		Window     string    `json:"window"`
//		Containers []struct {
//			Name  string `json:"name"`
//			Usage struct {
//				CPU    string `json:"cpu"`
//				Memory string `json:"memory"`
//			} `json:"usage"`
//		} `json:"containers"`
//	} `json:"items"`
//}
//
//func getMetrics(clientset *kubernetes.Clientset, pods *PodMetricsList) error {
//	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1/pods").DoRaw(context.TODO())
//	if err != nil {
//		return err
//	}
//	err = json.Unmarshal(data, &pods)
//	return err
//}
//
//func main() {
//	// creates the in-cluster config
//	// https://github.com/kubernetes/client-go/tree/master/examples#configuration
//	config, err := rest.InClusterConfig()
//	if err != nil {
//		panic(err.Error())
//	}
//	// creates the clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//	var pods PodMetricsList
//	for i := 0; i < 10; i++ {
//		err = getMetrics(clientset, &pods)
//		if err != nil {
//			panic(err.Error())
//		}
//		for _, m := range pods.Items {
//			fmt.Println(m.Metadata.Name, m.Metadata.Namespace, m.Timestamp.String())
//		}
//		time.Sleep(10 * time.Second)
//	}
//}

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"time"
)

const targetNamespaceEnv = "TARGET_NAMESPACE"

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("rest error")
		panic(err.Error())
	}
	clientset, err := metricsv.NewForConfig(config)
	if err != nil {
		fmt.Println("metricsv.NewForConfig(config) error")
		panic(err.Error())
	}
	ns := os.Getenv(targetNamespaceEnv)
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		fmt.Println("env: ", targetNamespaceEnv, " = ", ns)
		podMetricsList, err := clientset.MetricsV1beta1().PodMetricses(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("clientset error:\n", err)
		}
		fmt.Println(podMetricsList)
	}
}
