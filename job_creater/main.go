package main

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "job-test"

	c := clientset.BatchV1().Jobs(namespace)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pi",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "pi",
							Image: "perl",
							Command: []string{
								"perl",
								"-Mbignum=bpi",
								"-wle",
								"print bpi(2000)",
							},
							Resources: v1.ResourceRequirements{
								Limits: v1.ResourceList{
									v1.ResourceMemory: resource.MustParse("500Mi"),
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: int32Ptr(4),
		},
	}

	fmt.Println("Creating job...")
	result, err := c.Create(job)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Job %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
