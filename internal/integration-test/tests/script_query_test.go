package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	scriptlanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_script_query() {
	// Given - Weight'i height'tan büyük olan pokemonlar
	query := es.NewQuery(
		es.ScriptQuery(
			es.ScriptSource("doc['weight'].value > doc['height'].value * 10", scriptlanguage.Painless),
		),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar weight > height * 10 koşulunu sağlamalı
	for _, pokemon := range result {
		assert.Greater(s.T(), pokemon.Weight, pokemon.Height*10,
			"Pokemon %s için weight (%d) > height*10 (%d) koşulu sağlanmıyor",
			pokemon.Name, pokemon.Weight, pokemon.Height*10)
	}
}

func (s *testSuite) Test_it_should_return_heavy_pokemons_with_script_query() {
	// Given - Weight ve height toplamı 100'den büyük
	script := es.ScriptSource("doc['weight'].value + doc['height'].value > params.threshold", scriptlanguage.Painless).
		Parameter("threshold", 100)

	query := es.NewQuery(
		es.ScriptQuery(script),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar weight + height > 100 koşulunu sağlamalı
	for _, pokemon := range result {
		total := pokemon.Weight + pokemon.Height
		assert.Greater(s.T(), total, uint16(100),
			"Pokemon %s için weight+height (%d) > 100 koşulu sağlanmıyor",
			pokemon.Name, total)
	}
}
