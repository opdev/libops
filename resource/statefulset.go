package resource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetBuilder struct {
	target *appsv1.StatefulSet
}

type StatefulSetAnnotationBuilder struct {
	StatefulSetBuilder
}

func (s *StatefulSetBuilder) Named(name string) *StatefulSetBuilder {
	s.target.ObjectMeta.Name = name
	return s
}

func (s *StatefulSetBuilder) In(namespace string) *StatefulSetBuilder {
	s.target.ObjectMeta.Namespace = namespace
	return s
}

func (s *StatefulSetBuilder) OfSize(replicas int32) *StatefulSetBuilder {
	s.target.Spec.Replicas = &replicas
	return s
}

func (s *StatefulSetBuilder) Labeled(labels map[string]string) *StatefulSetBuilder {
	if s.target.ObjectMeta.Labels == nil {
		s.target.ObjectMeta.Labels = make(map[string]string)
	}

	if s.target.Spec.Template.Labels == nil {
		s.target.Spec.Template.Labels = make(map[string]string)
	}

	if s.target.Spec.Selector == nil {
		s.target.Spec.Selector = &metav1.LabelSelector{}
	}

	if s.target.Spec.Selector.MatchLabels == nil {
		s.target.Spec.Selector.MatchLabels = make(map[string]string)
	}

	for key, val := range labels {
		s.target.ObjectMeta.Labels[key] = val
		s.target.Spec.Selector.MatchLabels[key] = val
		s.target.Spec.Template.Labels[key] = val
	}
	return s
}

func (s *StatefulSetBuilder) WithAnnotations(annotations map[string]string) *StatefulSetBuilder {
	if s.target.ObjectMeta.Annotations == nil {
		s.target.ObjectMeta.Annotations = make(map[string]string)
	}
	for key, val := range s.target.ObjectMeta.Annotations {
		s.target.ObjectMeta.Annotations[key] = val
	}
	return s
}

func (s *StatefulSetBuilder) WithContainers(container ...corev1.Container) *StatefulSetBuilder {
	s.target.Spec.Template.Spec.Containers = append(
		s.target.Spec.Template.Spec.Containers,
		container...,
	)
	return s
}

func (s *StatefulSetBuilder) WithVolumes(volume ...corev1.Volume) *StatefulSetBuilder {
	s.target.Spec.Template.Spec.Volumes = append(
		s.target.Spec.Template.Spec.Volumes,
		volume...,
	)
	return s
}

func (s *StatefulSetBuilder) Annotate() *StatefulSetAnnotationBuilder {
	return &StatefulSetAnnotationBuilder{*s}
}

func (a *StatefulSetAnnotationBuilder) Version(version string) *StatefulSetAnnotationBuilder {
	if a.StatefulSetBuilder.target.ObjectMeta.Annotations == nil {
		a.StatefulSetBuilder.target.ObjectMeta.Annotations = make(map[string]string)
	}
	a.StatefulSetBuilder.target.ObjectMeta.Annotations["app.kubernetes.io/verison"] = version
	return a
}

func (a *StatefulSetAnnotationBuilder) Component(component string) *StatefulSetAnnotationBuilder {
	if a.StatefulSetBuilder.target.ObjectMeta.Annotations == nil {
		a.StatefulSetBuilder.target.ObjectMeta.Annotations = make(map[string]string)
	}
	a.StatefulSetBuilder.target.ObjectMeta.Annotations["app.kubernetes.io/component"] = component
	return a
}

func (s *StatefulSetBuilder) RunsAs(serviceAccount string) *StatefulSetBuilder {
	s.target.Spec.Template.Spec.ServiceAccountName = serviceAccount
	return s
}

func (s *StatefulSetBuilder) Build() *appsv1.StatefulSet {
	return s.target
}
