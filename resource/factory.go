package resource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
)

type ResourceBuilder struct{}

func Create() *ResourceBuilder {
	return &ResourceBuilder{}
}

func (r *ResourceBuilder) Pod() *PodBuilder {
	return &PodBuilder{
		target: &corev1.Pod{},
	}
}

func (r *ResourceBuilder) StatefulSet() *StatefulSetBuilder {
	return &StatefulSetBuilder{
		target: &appsv1.StatefulSet{},
	}
}

func (r *ResourceBuilder) Deployment() *DeploymentBuilder {
	return &DeploymentBuilder{
		target: &appsv1.Deployment{},
	}
}

func (r *ResourceBuilder) Service() *ServiceBuilder {
	return &ServiceBuilder{
		target: &corev1.Service{},
	}
}

func (r *ResourceBuilder) PersistentVolumeClaim() *PersistentVolumeClaimBuilder {
	return &PersistentVolumeClaimBuilder{
		target: &corev1.PersistentVolumeClaim{},
	}
}

func (r *ResourceBuilder) Ingress() *IngressBuilder {
	return &IngressBuilder{
		target: &networkv1.Ingress{},
	}
}

func (r *ResourceBuilder) Container() *ContainerBuilder {
	return &ContainerBuilder{
		target: &corev1.Container{},
	}
}
