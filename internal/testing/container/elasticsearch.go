package container

import (
	"context"
	"fmt"
	"runtime"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	ElasticsearchImage = "docker.elastic.co/elasticsearch/elasticsearch:8.15.0"
	defaultPort        = "9200"
)

type ElasticsearchContainer struct {
	address          string
	ip               string
	port             nat.Port
	container        testcontainers.Container
	containerRequest testcontainers.ContainerRequest
}

func NewContainer(image string) *ElasticsearchContainer {
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{fmt.Sprintf("%s:%s", defaultPort, defaultPort)},
		Env: map[string]string{
			"cluster.name":                    "testcontainers-go",
			"discovery.type":                  "single-node",
			"bootstrap.memory_lock":           "true",
			"xpack.security.enabled":          "false", // Disable security features (including TLS)
			"xpack.security.http.ssl.enabled": "false", // Disable HTTPS for the HTTP API
			"ES_JAVA_OPTS":                    "-Xms1g -Xmx1g",
		},
		WaitingFor: wait.ForLog("up and running"),
	}
	return &ElasticsearchContainer{
		containerRequest: req,
	}
}

func (c *ElasticsearchContainer) Run() (err error) {
	c.container, err = testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: c.containerRequest,
		Started:          true,
	})
	if err != nil {
		return err
	}

	c.ip, err = c.container.Host(context.Background())
	if err != nil {
		return err
	}
	c.port, err = c.container.MappedPort(context.Background(), defaultPort)
	if err != nil {
		return err
	}

	if isRunningOnOSX() {
		c.ip = "127.0.0.1"
	}

	return nil
}

func (c *ElasticsearchContainer) TerminateContainer() (err error) {
	if c.container != nil {
		return c.container.Terminate(context.Background())
	}

	return nil
}

func (c *ElasticsearchContainer) Host() string {
	return fmt.Sprintf("http://%s:%s", c.ip, c.port.Port())
}

func isRunningOnOSX() bool {
	return runtime.GOOS == "darwin"
}
