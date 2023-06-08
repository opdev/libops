package resource

import (
	corev1 "k8s.io/api/core/v1"
)

type PodBuilder struct {
	target *corev1.Pod
}

func (p *PodBuilder) Named(name string) *PodBuilder {
	p.target.ObjectMeta.Name = name
	return p
}

func (p *PodBuilder) In(namespace string) *PodBuilder {
	p.target.ObjectMeta.Namespace = namespace
	return p
}

func (p *PodBuilder) Labeled(labels map[string]string) *PodBuilder {
	if p.target.ObjectMeta.Labels == nil {
		p.target.ObjectMeta.Labels = make(map[string]string)
	}

	for key, val := range labels {
		p.target.ObjectMeta.Labels[key] = val
	}
	return p
}

func (p *PodBuilder) WithAnnotations(annotations map[string]string) *PodBuilder {
	if p.target.ObjectMeta.Annotations == nil {
		p.target.ObjectMeta.Annotations = make(map[string]string)
	}
	for key, val := range p.target.ObjectMeta.Annotations {
		p.target.ObjectMeta.Annotations[key] = val
	}
	return p
}

func (p *PodBuilder) WithContainers(container ...corev1.Container) *PodBuilder {
	p.target.Spec.Containers = append(
		p.target.Spec.Containers,
		container...,
	)
	return p
}

func (p *PodBuilder) WithVolumes(volume ...corev1.Volume) *PodBuilder {
	p.target.Spec.Volumes = append(
		p.target.Spec.Volumes,
		volume...,
	)
	return p
}

func (p *PodBuilder) RunsAs(serviceAccount string) *PodBuilder {
	p.target.Spec.ServiceAccountName = serviceAccount
	return p
}

func (p *PodBuilder) Build() *corev1.Pod {
	return p.target
}
