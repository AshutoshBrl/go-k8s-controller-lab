package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	// "k8s.io/client-go/tools/cache"
)

func main() {
	// Step 1: Load kubeconfig
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	// Step 2: Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	// Step 3: Set up context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Step 4: Create a watcher for Pods
	watcher, err := clientset.CoreV1().Pods("default").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to start watcher: %v", err)
	}

	// Step 5: Process watch events
	fmt.Println("Watching Pods in the default namespace...")
	for event := range watcher.ResultChan() {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			log.Println("Unexpected object type")
			continue
		}
		fmt.Printf("Event: %s, Pod: %s\n", event.Type, pod.Name)
	}

	fmt.Println("Watcher stopped")
}