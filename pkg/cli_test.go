package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource2Target(t *testing.T) {
	cli := &Cli{
		username:   "togettoyou",
		repository: "",
	}

	output, err := cli.Source2Target("", "")
	assert.Nil(t, output)

	source := "registry.k8s.io/kube-apiserver"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, source, output.Source)
	assert.Equal(t, "docker.io/togettoyou/registry.k8s.io.kube-apiserver", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, source, output.Source)
	assert.Equal(t, "docker.io/togettoyou/registry.k8s.io.kube-apiserver:v1.27.4", output.Target)

	source = "registry.k8s.io/kube-apiserver$kube-apiserver"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver", output.Source)
	assert.Equal(t, "docker.io/togettoyou/kube-apiserver", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4$kube-apiserver"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver:v1.27.4", output.Source)
	assert.Equal(t, "docker.io/togettoyou/kube-apiserver:v1.27.4", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4$kube-apiserver:mytag"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver:v1.27.4", output.Source)
	assert.Equal(t, "docker.io/togettoyou/kube-apiserver:mytag", output.Target)

	source = "nginx@sha256:123456$nginx"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "nginx@sha256:123456", output.Source)
	assert.Equal(t, "docker.io/togettoyou/nginx:123456", output.Target)

	source = "nginx@sha256:123456$nginx:mytag"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "nginx@sha256:123456", output.Source)
	assert.Equal(t, "docker.io/togettoyou/nginx:mytag", output.Target)

	source = "golang"
	output, err = cli.Source2Target(source, "linux/arm64/v8")
	assert.Nil(t, err)
	assert.Equal(t, "golang", output.Source)
	assert.Equal(t, "docker.io/togettoyou/golang-linux-arm64-v8", output.Target)

	source = "golang:1.21.6"
	output, err = cli.Source2Target(source, "linux/arm64/v8")
	assert.Nil(t, err)
	assert.Equal(t, "golang:1.21.6", output.Source)
	assert.Equal(t, "docker.io/togettoyou/golang-linux-arm64-v8:1.21.6", output.Target)

	source = "golang:1.21.6$mygolang"
	output, err = cli.Source2Target(source, "linux/arm64/v8")
	assert.Nil(t, err)
	assert.Equal(t, "golang:1.21.6", output.Source)
	assert.Equal(t, "docker.io/togettoyou/mygolang-linux-arm64-v8:1.21.6", output.Target)

	source = "golang:1.21.6$mygolang:1.21.6arm64"
	output, err = cli.Source2Target(source, "linux/arm64/v8")
	assert.Nil(t, err)
	assert.Equal(t, "golang:1.21.6", output.Source)
	assert.Equal(t, "docker.io/togettoyou/mygolang-linux-arm64-v8:1.21.6arm64", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4"
	output, err = cli.Source2Target(source, "arm64")
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver:v1.27.4", output.Source)
	assert.Equal(t, "docker.io/togettoyou/registry.k8s.io.kube-apiserver-arm64:v1.27.4", output.Target)

	source = "public.ecr.aws/docker/library/golang:1.27.1-trixie@sha256:433790e515d27dc6003e847e644cc0af956985cf315c1c58a3b73ee2dd305183"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, source, output.Source)
	assert.Equal(t, "docker.io/togettoyou/public.ecr.aws.docker.library.golang:1.27.1-trixie", output.Target)

	source = "public.ecr.aws/docker/library/golang:1.27.1-trixie@sha256:433790e515d27dc6003e847e644cc0af956985cf315c1c58a3b73ee2dd305183$golang:1.27.1-trixie"
	output, err = cli.Source2Target(source, "")
	assert.Nil(t, err)
	assert.Equal(t, "public.ecr.aws/docker/library/golang:1.27.1-trixie@sha256:433790e515d27dc6003e847e644cc0af956985cf315c1c58a3b73ee2dd305183", output.Source)
	assert.Equal(t, "docker.io/togettoyou/golang:1.27.1-trixie", output.Target)
}
