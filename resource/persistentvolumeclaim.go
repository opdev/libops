package resource

import (
	corev1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaimBuilder struct {
	target *corev1.PersistentVolumeClaim
}

func (c *PersistentVolumeClaimBuilder) Named(name string) *PersistentVolumeClaimBuilder {
	c.target.Name = name
	return c
}

func (c *PersistentVolumeClaimBuilder) In(namespace string) *PersistentVolumeClaimBuilder {
	c.target.Namespace = namespace
	return c
}

func (c *PersistentVolumeClaimBuilder) Labeled(labels map[string]string) *PersistentVolumeClaimBuilder {
	if c.target.ObjectMeta.Labels == nil {
		c.target.ObjectMeta.Labels = make(map[string]string)
	}

	for key, value := range labels {
		c.target.ObjectMeta.Labels[key] = value
	}
	return c
}

func (c *PersistentVolumeClaimBuilder) AccessModes(mode ...corev1.PersistentVolumeAccessMode) *PersistentVolumeClaimBuilder {
	c.target.Spec.AccessModes = append(c.target.Spec.AccessModes, mode...)
	return c
}

func (c *PersistentVolumeClaimBuilder) ResourceRequests(request corev1.ResourceList) *PersistentVolumeClaimBuilder {
	c.target.Spec.Resources.Requests = request
	return c
}

func (c *PersistentVolumeClaimBuilder) ResourceLimits(limit corev1.ResourceList) *PersistentVolumeClaimBuilder {
	c.target.Spec.Resources.Limits = limit
	return c
}

func (c *PersistentVolumeClaimBuilder) StorageClass(name string) *PersistentVolumeClaimBuilder {
	c.target.Spec.StorageClassName = &name
	return c
}

func (c *PersistentVolumeClaimBuilder) Build() *corev1.PersistentVolumeClaim {
	return c.target
}
