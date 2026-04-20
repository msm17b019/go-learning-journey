package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(
		os.Getenv("HOME"),
		".kube",
		"config",
	)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().
		Pods("").
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	namespaceCount := make(map[string]int)

	for _, pod := range pods.Items {

		ns := pod.Namespace

		namespaceCount[ns]++
	}

	fmt.Println("Pods per namespace:")

	for ns, count := range namespaceCount {

		fmt.Printf(
			"Namespace: %s -> Pods: %d\n",
			ns,
			count,
		)
	}
}
