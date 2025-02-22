package tests_test

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"integration-tests/constants"
	"integration-tests/model"
	"integration-tests/repository"
	testdataprovider "integration-tests/test-data-provider"
	"os"
	"testing"

	"integration-tests/container"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	TestContext                    context.Context
	ElasticContainer               *container.ElasticsearchContainer
	ESClient                       *elasticsearch.Client
	FooElasticsearchRepository     repository.BaseGenericRepository[string, model.FooDocument]
	PokedexElasticsearchRepository repository.BaseGenericRepository[string, model.Pokemon]
}

func (s *testSuite) SetupSuite() {
	s.TestContext = context.Background()
	s.ElasticContainer = container.NewContainer(s.TestContext, container.ElasticsearchImage)
	err := s.ElasticContainer.Run()
	if err != nil {
		fmt.Printf("error starting elasticsearch container. err %s", err.Error())
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
	s.FooElasticsearchRepository = repository.NewFooElasticsearchRepository(s.ESClient)
	s.PokedexElasticsearchRepository = repository.NewPokedexElasticsearchRepository(s.ESClient)

	// create indicies
	s.createIndexRequest(constants.TestIndex, &model.FooDocument{})
	s.createIndexRequest(constants.PokemonIndex, &model.Pokemon{})

	// fill incidies
	pokemonData, err := testdataprovider.PokemonTestData()
	if err != nil {
		s.T().FailNow()
	}
	s.PokedexElasticsearchRepository.BulkInsert(s.TestContext, pokemonData...)
}

func (s *testSuite) TearDownSuite() {
	if err := s.ElasticContainer.TerminateContainer(); err != nil {
		s.T().FailNow()
	}
}

func (s *testSuite) createIndexRequest(indexName string, index repository.ElasticsearchIndex) {
	indexBody, err := repository.CreateIndexBody(index)
	if err != nil {
		fmt.Printf("#createIndexRequest - error generation index body. err %s\n", err.Error())
		s.T().FailNow()
	}
	indicesRequest := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  indexBody,
	}
	if _, err = indicesRequest.Do(s.TestContext, s.ESClient); err != nil {
		fmt.Printf("#createIndexRequest - error creating index. err %s\n", err.Error())
		s.T().FailNow()
	}
}
