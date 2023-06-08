package resource

import (
	networkv1 "k8s.io/api/networking/v1"
)

type IngressBuilder struct {
	target *networkv1.Ingress
}

func (i *IngressBuilder) Named(name string) *IngressBuilder {
	i.target.Name = name
	return i
}

func (i *IngressBuilder) In(namespace string) *IngressBuilder {
	i.target.Namespace = namespace
	return i
}

func (i *IngressBuilder) Labeled(labels map[string]string) *IngressBuilder {
	i.target.ObjectMeta.Labels = labels
	return i
}

func (i *IngressBuilder) OfClass(ingressClass string) *IngressBuilder {
	i.target.Spec.IngressClassName = &ingressClass
	return i
}
