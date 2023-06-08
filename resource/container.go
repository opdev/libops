package resource

import corev1 "k8s.io/api/core/v1"

type ContainerBuilder struct {
	target *corev1.Container
}

func (c *ContainerBuilder) Named(name string) *ContainerBuilder {
	c.target.Name = name
	return c
}

func (c *ContainerBuilder) Image(image string) *ContainerBuilder {
	c.target.Image = image
	return c
}

func (c *ContainerBuilder) Command(cmd []string) *ContainerBuilder {
	c.target.Command = cmd
	return c
}

func (c *ContainerBuilder) Args(args []string) *ContainerBuilder {
	c.target.Args = args
	return c
}

func (c *ContainerBuilder) EnvironmentVars(envvar ...corev1.EnvVar) *ContainerBuilder {
	c.target.Env = append(c.target.Env, envvar...)
	return c
}

func (c *ContainerBuilder) EnvironmentsFrom(envSource ...corev1.EnvFromSource) *ContainerBuilder {
	c.target.EnvFrom = append(c.target.EnvFrom, envSource...)
	return c
}

func (c *ContainerBuilder) VolumeMounts(mount ...corev1.VolumeMount) *ContainerBuilder {
	c.target.VolumeMounts = append(c.target.VolumeMounts, mount...)
	return c
}

func (c *ContainerBuilder) LivensesProbe(liveProbe *corev1.Probe) *ContainerBuilder {
	c.target.LivenessProbe = liveProbe
	return c
}

func (c *ContainerBuilder) ReadinessProbe(readyProbe *corev1.Probe) *ContainerBuilder {
	c.target.ReadinessProbe = readyProbe
	return c
}

func (c *ContainerBuilder) Exposes(port ...corev1.ContainerPort) *ContainerBuilder {
	c.target.Ports = append(c.target.Ports, port...)
	return c
}

func (c *ContainerBuilder) Build() corev1.Container {
	return *c.target
}
