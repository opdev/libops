package resource

import (
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentBuilder struct {
	target *appsv1.Deployment
}

type DeploymentAnnotationBuilder struct {
	DeploymentBuilder
}

func (d *DeploymentBuilder) Named(name string) *DeploymentBuilder {
	d.target.ObjectMeta.Name = name
	return d
}

func (d *DeploymentBuilder) In(namespace string) *DeploymentBuilder {
	d.target.ObjectMeta.Namespace = namespace
	return d
}

func (d *DeploymentBuilder) OfSize(replicas int32) *DeploymentBuilder {
	d.target.Spec.Replicas = &replicas
	return d
}

func (d *DeploymentBuilder) Labeled(labels map[string]string) *DeploymentBuilder {
	if d.target.ObjectMeta.Labels == nil {
		d.target.ObjectMeta.Labels = make(map[string]string)
	}

	if d.target.Spec.Template.Labels == nil {
		d.target.Spec.Template.Labels = make(map[string]string)
	}

	if d.target.Spec.Selector == nil {
		d.target.Spec.Selector = &metav1.LabelSelector{}
	}

	if d.target.Spec.Selector.MatchLabels == nil {
		d.target.Spec.Selector.MatchLabels = make(map[string]string)
	}

	for key, val := range labels {
		d.target.ObjectMeta.Labels[key] = val
		d.target.Spec.Selector.MatchLabels[key] = val
		d.target.Spec.Template.Labels[key] = val
	}
	return d
}

func (d *DeploymentBuilder) WithAnnotations(annotations map[string]string) *DeploymentBuilder {
	if d.target.ObjectMeta.Annotations == nil {
		d.target.ObjectMeta.Annotations = make(map[string]string)
	}
	for key, val := range d.target.ObjectMeta.Annotations {
		d.target.ObjectMeta.Annotations[key] = val
	}
	return d
}

func (d *DeploymentBuilder) WithContainers(container ...corev1.Container) *DeploymentBuilder {
	d.target.Spec.Template.Spec.Containers = append(
		d.target.Spec.Template.Spec.Containers,
		container...,
	)
	return d
}

func (d *DeploymentBuilder) WithVolumes(volume ...corev1.Volume) *DeploymentBuilder {
	d.target.Spec.Template.Spec.Volumes = append(
		d.target.Spec.Template.Spec.Volumes,
		volume...,
	)
	return d
}

func (d *DeploymentBuilder) Annotate() *DeploymentAnnotationBuilder {
	return &DeploymentAnnotationBuilder{*d}
}

func (a *DeploymentAnnotationBuilder) Version(version string) *DeploymentAnnotationBuilder {
	if a.DeploymentBuilder.target.ObjectMeta.Annotations == nil {
		a.DeploymentBuilder.target.ObjectMeta.Annotations = make(map[string]string)
	}
	a.DeploymentBuilder.target.ObjectMeta.Annotations["app.kubernetes.io/version"] = version
	return a
}

func (a *DeploymentAnnotationBuilder) Component(component string) *DeploymentAnnotationBuilder {
	if a.DeploymentBuilder.target.ObjectMeta.Annotations == nil {
		a.DeploymentBuilder.target.ObjectMeta.Annotations = make(map[string]string)
	}
	a.DeploymentBuilder.target.ObjectMeta.Annotations["app.kubernetes.io/component"] = component
	return a
}

func (d *DeploymentBuilder) RunsAs(serviceAccount string) *DeploymentBuilder {
	d.target.Spec.Template.Spec.ServiceAccountName = serviceAccount
	return d
}

func (d *DeploymentBuilder) Build() *appsv1.Deployment {
	return d.target
}

func (d *DeploymentBuilder) Validate() error {
	if d.target.Name == "" {
		return errored("missing name; Use .Named() method to define name of deployment")
	}

	if d.target.Namespace == "" {
		return errored("missing namespace; Use .In() method to specify namespace ")
	}

	if d.target.ObjectMeta.Labels == nil {
		return errored("no labels provided; Use .Labeled() method to provided labels")
	}

	return nil
}

func errored(msg string) error {
	return errors.New(msg)
}
