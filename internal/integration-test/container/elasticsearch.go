package container

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/moby/moby/api/types/network"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// Current major representatives

	// v7 minor/patch set
	ElasticsearchImage_v7_10_2  = "docker.elastic.co/elasticsearch/elasticsearch:7.10.2"
	ElasticsearchImage_v7_14_2  = "docker.elastic.co/elasticsearch/elasticsearch:7.14.2"
	ElasticsearchImage_v7_15_0  = "docker.elastic.co/elasticsearch/elasticsearch:7.15.0"
	ElasticsearchImage_v7_17_28 = "docker.elastic.co/elasticsearch/elasticsearch:7.17.28"

	// v8 minor/patch set
	ElasticsearchImage_v8_11_4  = "docker.elastic.co/elasticsearch/elasticsearch:8.11.4"
	ElasticsearchImage_v8_13_4  = "docker.elastic.co/elasticsearch/elasticsearch:8.13.4"
	ElasticsearchImage_v8_15_0  = "docker.elastic.co/elasticsearch/elasticsearch:8.15.0"
	ElasticsearchImage_v8_17_18 = "docker.elastic.co/elasticsearch/elasticsearch:8.17.10"

	// v9
	ElasticsearchImage_v9_0_0 = "docker.elastic.co/elasticsearch/elasticsearch:9.0.0"

	defaultPort = "9200/tcp"
)

type ElasticsearchContainer struct {
	containerContext context.Context
	containerRequest testcontainers.ContainerRequest
	container        testcontainers.Container
	address          string
	ip               string
	port             network.Port
}

func NewContainer(ctx context.Context, image string) *ElasticsearchContainer {
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{defaultPort},
		Env: map[string]string{
			"cluster.name":                    "testcontainers-go",
			"discovery.type":                  "single-node",
			"bootstrap.memory_lock":           "true",
			"xpack.security.enabled":          "false", // Disable security features (including TLS)
			"xpack.security.http.ssl.enabled": "false", // Disable HTTPS for the HTTP API
			"ES_JAVA_OPTS":                    "-Xms1g -Xmx1g",
		},
		WaitingFor: wait.ForHTTP("/").
			WithPort(defaultPort).
			WithStartupTimeout(2 * time.Minute),
	}
	return &ElasticsearchContainer{
		containerContext: ctx,
		containerRequest: req,
	}
}

func (c *ElasticsearchContainer) Run() (err error) {
	c.container, err = testcontainers.GenericContainer(c.containerContext, testcontainers.GenericContainerRequest{
		ContainerRequest: c.containerRequest,
		Started:          true,
	})
	if err != nil {
		return err
	}

	c.ip, err = c.container.Host(c.containerContext)
	if err != nil {
		return err
	}
	c.port, err = c.container.MappedPort(c.containerContext, defaultPort)
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
		return c.container.Terminate(c.containerContext)
	}

	return nil
}

func (c *ElasticsearchContainer) Host() string {
	return fmt.Sprintf("http://%s:%s", c.ip, c.port.Port())
}

func isRunningOnOSX() bool {
	return runtime.GOOS == "darwin"
}
