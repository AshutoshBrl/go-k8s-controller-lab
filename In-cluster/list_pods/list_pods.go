package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	// Step 1: Load the kubeconfig file
	// Typically located at ~/.kube/config
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	// Step 2: Create a Kubernetes clientset
	// The clientset provides access to Kubernetes API groups
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	// Step 3: List Pods in the default namespace
	// Use the CoreV1 API group to interact with Pods
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list pods: %v", err)
	}

	// Step 4: Print the names of all Pods
	fmt.Println("Pods in the default namespace:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}