# libops

libops is an experimental golang module that aims to generate kubernetes resources in a consistent pattern.

## Example

```
Package main

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/opdev/libops/resource"
)

func nginxDeployment() *appsv1.Deployment {
	return resource.Create().
		Deployment().
		In("demo-ns").
		Named("nginx").
		Labeled(
			map[string]string{
				"app": "nginx",
			},
		).
		WithContainers(
			// define nginx container
			resource.Create().
				Container().
				Named("nginx").
				Image("nginx:latest").
				Exposes(
					corev1.ContainerPort{
						Name:          "http",
						Protocol:      corev1.ProtocolTCP,
						ContainerPort: 80,
					},
					corev1.ContainerPort{
						Name:          "https",
						Protocol:      corev1.ProtocolTCP,
						ContainerPort: 443,
					},
				).
				Build(),
		).
		Build()
}

func main() {
	fmt.Printf("%+v\n", nginxDeployment())
}
```
