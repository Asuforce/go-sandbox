package main

import (
	"fmt"
	"math/rand"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rs   = "abcdefghijklmnopqrstuvwxyz"
	bits = 6
	mask = 1<<bits - 1
	max  = 63 / bits
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

	jobName := "pl-" + randString(10)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  jobName,
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

func randString(n int) string {
	b := make([]byte, n)
	c, r := randSrc.Int63(), max
	for i := n - 1; i >= 0; {
		if r == 0 {
			c, r = randSrc.Int63(), max
		}
		idx := int(c & mask)
		if idx < len(rs) {
			b[i] = rs[idx]
			i--
		}
		c >>= bits
		r--
	}

	return string(b)
}
