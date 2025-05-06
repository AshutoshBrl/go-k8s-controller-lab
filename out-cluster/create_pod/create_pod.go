package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

	// Step 3: Define a Pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-pod",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx:latest",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	// Step 4: Create the Pod
	fmt.Println("Creating Pod...")
	createdPod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create pod: %v", err)
	}
	fmt.Printf("Created Pod: %s\n", createdPod.Name)

	// Step 5: Get the Pod's details
	fetchedPod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx-pod", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get pod: %v", err)
	}
	fmt.Printf("Pod Status: %s\n", fetchedPod.Status.Phase)
}