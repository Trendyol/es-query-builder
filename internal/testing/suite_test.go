package testing

import (
	"context"
	"fmt"
	"integration-tests/container"
	"os"
	"strings"
	"testing"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/stretchr/testify/suite"
)

const (
	testIndexName = "foo-index"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	ElasticContainer        *container.ElasticsearchContainer
	ESClient                *elasticsearch.Client
	ElasticsearchRepository ElasticsearchRepository
}

func (s *testSuite) SetupSuite() {
	s.ElasticContainer = container.NewContainer(container.ElasticsearchImage)
	err := s.ElasticContainer.Run()
	if err != nil {
		fmt.Println(fmt.Printf("error starting elasticsearch container. err %v", err))
		s.T().FailNow()
	}

	cfg := elasticsearch.Config{
		Addresses: []string{s.ElasticContainer.Host()},
		Logger: &elastictransport.ColorLogger{
			Output: os.Stdout,
		},
		DiscoverNodesOnStart: false,
	}
	s.ESClient, err = elasticsearch.NewClient(cfg)
	s.ElasticsearchRepository = NewElasticsearchRepository(s.ESClient)

	indicesRequest := esapi.IndicesCreateRequest{
		Index: testIndexName,
		Body:  strings.NewReader(testIndexBody()),
	}

	_, err = indicesRequest.Do(context.Background(), s.ESClient)
	if err != nil {
		fmt.Println(fmt.Printf("error creating index. err %v", err))
		s.T().FailNow()
	}
}

func (s *testSuite) TearDownSuite() {
	err := s.ElasticContainer.TerminateContainer()
	if err != nil {
		s.T().FailNow()
	}
}

func testIndexBody() string {
	return `{
			"settings": {
				"index": {
					"refresh_interval": "1ms",
					"number_of_shards": "1",
					"number_of_replicas": "1",
					"max_result_window": "50000",
					"max_terms_count": "100"
				}
			},
			"mappings": {
				"properties": {
					"foo": {
						"type": "keyword"
					},
					"meta": {
						"properties": {
							"id": {
								"type": "keyword"
							}
						}
					}
				}
			}
		}`
}
