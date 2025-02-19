package tests_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"integration-tests"
	"integration-tests/constants"
	"integration-tests/container"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	TestContext             context.Context
	ElasticContainer        *container.ElasticsearchContainer
	ESClient                *elasticsearch.Client
	ElasticsearchRepository integrationtest.ElasticsearchRepository
}

func (s *testSuite) SetupSuite() {
	s.TestContext = context.Background()
	s.ElasticContainer = container.NewContainer(s.TestContext, container.ElasticsearchImage)
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
	s.ElasticsearchRepository = integrationtest.NewElasticsearchRepository(s.ESClient)

	indicesRequest := esapi.IndicesCreateRequest{
		Index: constants.TestIndex,
		Body:  strings.NewReader(testIndexBody()),
	}

	if _, err = indicesRequest.Do(s.TestContext, s.ESClient); err != nil {
		fmt.Printf("error creating index. err %v\n", err)
		s.T().FailNow()
	}
}

func (s *testSuite) TearDownSuite() {
	if err := s.ElasticContainer.TerminateContainer(); err != nil {
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
