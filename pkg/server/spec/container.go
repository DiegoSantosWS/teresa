package spec

import (
	"fmt"
	"strconv"

	"github.com/luizalabs/teresa/pkg/server/app"
	"github.com/luizalabs/teresa/pkg/server/storage"
)

const (
	nginxConfTmplDir = "/etc/nginx/template/"
	nginxConfDir     = "/etc/nginx/"
	nginxConfFile    = "nginx.conf"
	nginxArgTmpl     = "envsubst < %s%s > %s%s && nginx -g 'daemon off;'"
	nginxBackendTmpl = "http://localhost:%d"
)

type ContainerLimits struct {
	CPU    string
	Memory string
}

type VolumeMounts struct {
	Name      string
	MountPath string
	ReadOnly  bool
}

type Port struct {
	Name          string
	ContainerPort int32
}

type Container struct {
	Name            string
	Image           string
	ContainerLimits *ContainerLimits
	Env             map[string]string
	VolumeMounts    []*VolumeMounts
	Command         []string
	Args            []string
	Ports           []Port
	Secrets         []string
}

func newSlugVolumeMount() *VolumeMounts {
	return &VolumeMounts{
		Name:      slugVolumeName,
		MountPath: slugVolumeMountPath,
	}
}

func newStorageKeyVolumeMount() *VolumeMounts {
	return &VolumeMounts{
		Name:      "storage-keys",
		MountPath: "/var/run/secrets/deis/objectstore/creds",
		ReadOnly:  true,
	}
}

func newInitContainers(slugURL, image string, a *app.App, fs storage.Storage) []*Container {
	ic := &Container{
		Name:  "slugstore",
		Image: image,
		Env: map[string]string{
			"BUILDER_STORAGE": fs.Type(),
			"SLUG_URL":        slugURL,
			"SLUG_DIR":        slugVolumeMountPath,
		},
		VolumeMounts: []*VolumeMounts{
			newStorageKeyVolumeMount(),
			newSlugVolumeMount(),
		},
	}
	for k, v := range fs.PodEnvVars() {
		ic.Env[k] = v
	}
	return []*Container{ic}
}

func newNginxVolumeMount() *VolumeMounts {
	return &VolumeMounts{
		Name:      "nginx-conf",
		MountPath: nginxConfTmplDir,
		ReadOnly:  true,
	}
}

func newNginxContainer(image string) *Container {
	args := fmt.Sprintf(nginxArgTmpl, nginxConfTmplDir, nginxConfFile, nginxConfDir, nginxConfFile)
	port := strconv.Itoa(DefaultPort)
	backend := fmt.Sprintf(nginxBackendTmpl, secondaryPort)

	return &Container{
		Name:    "nginx",
		Image:   image,
		Command: []string{"/bin/sh"},
		Args:    []string{"-c", args},
		Ports: []Port{{
			Name:          "nginx",
			ContainerPort: int32(DefaultPort),
		}},
		Env: map[string]string{
			"NGINX_PORT":    port,
			"NGINX_BACKEND": backend,
		},
		VolumeMounts: []*VolumeMounts{
			newNginxVolumeMount(),
		},
	}
}

func newAppContainer(name, image string, envVars map[string]string, port int, secrets []string) *Container {
	return &Container{
		Name:    name,
		Image:   image,
		Env:     envVars,
		Secrets: secrets,
		Ports: []Port{{
			Name:          "app",
			ContainerPort: int32(port),
		}},
	}
}
