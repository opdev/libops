package resource

import (
	corev1 "k8s.io/api/core/v1"
)

type ServiceBuilder struct {
	target *corev1.Service
}

func (s *ServiceBuilder) Named(name string) *ServiceBuilder {
	s.target.Name = name
	return s
}

func (s *ServiceBuilder) In(namespace string) *ServiceBuilder {
	s.target.Namespace = namespace
	return s
}

func (s *ServiceBuilder) Labeled(labels map[string]string) *ServiceBuilder {
	if s.target.ObjectMeta.Labels == nil {
		s.target.ObjectMeta.Labels = make(map[string]string)
	}

	for key, value := range labels {
		s.target.ObjectMeta.Labels[key] = value
	}

	return s
}

func (s *ServiceBuilder) WithSelectors(selectors map[string]string) *ServiceBuilder {
	if s.target.Spec.Selector == nil {
		s.target.Spec.Selector = make(map[string]string)
	}

	for key, value := range selectors {
		s.target.Spec.Selector[key] = value
	}
	return s
}

func (s *ServiceBuilder) Ports(svc ...corev1.ServicePort) *ServiceBuilder {
	s.target.Spec.Ports = append(s.target.Spec.Ports, svc...)
	return s
}

func (s *ServiceBuilder) OfType(svcType corev1.ServiceType) *ServiceBuilder {
	s.target.Spec.Type = svcType
	return s
}

func (s *ServiceBuilder) Build() *corev1.Service {
	return s.target
}
