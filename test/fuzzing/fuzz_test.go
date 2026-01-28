package fuzzing

import (
	"encoding/json"
	"testing"

	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"

	"github.com/Trendyol/es-query-builder/es"
)

func Fuzz_Term_Query(f *testing.F) {
	// Add seed corpus
	f.Add("field_name", "value")
	f.Add("status", "active")
	f.Add("category.keyword", "electronics")
	f.Add("", "")
	f.Add("field", "special!@#$%^&*()")
	f.Add("nested.deep.field", "unicode: 日本語 🎉")

	f.Fuzz(func(t *testing.T, field, value string) {
		// Build query - should not panic
		query := es.Term(field, value)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(query)
		if err != nil {
			return // Invalid JSON is acceptable, panics are not
		}

		// Unmarshal back - should produce valid JSON structure
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Match_Query(f *testing.F) {
	// Add seed corpus
	f.Add("title", "elasticsearch query")
	f.Add("description", "full text search")
	f.Add("content", "")
	f.Add("", "empty field")
	f.Add("text.analyzed", "special chars: <>&\"'")

	f.Fuzz(func(t *testing.T, field, query string) {
		// Build query - should not panic
		matchQuery := es.Match(field, query)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(matchQuery)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Bool_Query(f *testing.F) {
	// Add seed corpus
	f.Add("field1", "value1", "field2", "value2", 1)
	f.Add("status", "active", "category", "books", 2)
	f.Add("", "", "", "", 0)

	f.Fuzz(func(t *testing.T, field1, value1, field2, value2 string, minShouldMatch int) {
		// Build complex bool query - should not panic
		boolQuery := es.Bool().
			Must(es.Term(field1, value1)).
			Filter(es.Term(field2, value2)).
			Should(es.Match(field1, value2)).
			MustNot(es.Term(field2, value1)).
			MinimumShouldMatch(minShouldMatch)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(boolQuery)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Range_Query(f *testing.F) {
	// Add seed corpus
	f.Add("age", int64(18), int64(65))
	f.Add("price", int64(0), int64(1000))
	f.Add("timestamp", int64(-1000), int64(1000))
	f.Add("", int64(0), int64(0))

	f.Fuzz(func(t *testing.T, field string, gte, lte int64) {
		// Build range query - should not panic
		rangeQuery := es.Range(field).
			GreaterThanOrEqual(gte).
			LessThanOrEqual(lte)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(rangeQuery)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_New_Query(f *testing.F) {
	// Add seed corpus
	f.Add("title", "search term", 10, 0)
	f.Add("content", "full text", 100, 50)
	f.Add("", "", 0, 0)

	f.Fuzz(func(t *testing.T, field, value string, size, from int) {
		// Build full query - should not panic
		query := es.NewQuery(
			es.Bool().
				Must(es.Match(field, value)),
		).Size(size).From(from)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Complex_Query(f *testing.F) {
	// Add seed corpus
	f.Add("status", "active", "category", "electronics", int64(100), int64(500), true)
	f.Add("type", "product", "brand", "apple", int64(0), int64(1000), false)
	f.Add("", "", "", "", int64(0), int64(0), true)

	f.Fuzz(func(t *testing.T, field1, value1, field2, value2 string, priceMin, priceMax int64, trackHits bool) {
		// Build complex query with multiple features - should not panic
		query := es.NewQuery(
			es.Bool().
				Must(
					es.Term(field1, value1),
					es.Match(field2, value2),
				).
				Filter(
					es.Range("price").
						GreaterThanOrEqual(priceMin).
						LessThanOrEqual(priceMax),
				).
				Should(
					es.Term(field2, value1),
				).
				MinimumShouldMatch(1),
		).
			Size(10).
			From(0).
			TrackTotalHits(trackHits)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Term_Query_With_Options(f *testing.F) {
	// Add seed corpus
	f.Add("field", "value", 1.5, true)
	f.Add("status.keyword", "ACTIVE", 2.0, false)
	f.Add("", "", 0.0, true)

	f.Fuzz(func(t *testing.T, field, value string, boost float64, caseInsensitive bool) {
		// Build term query with options - should not panic
		query := es.Term(field, value).
			Boost(boost).
			CaseInsensitive(caseInsensitive)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Source_Includes(f *testing.F) {
	// Add seed corpus
	f.Add("title", "author", "date")
	f.Add("field1", "field2", "field3")
	f.Add("", "", "")
	f.Add("nested.field", "another.nested", "deep.nested.field")

	f.Fuzz(func(t *testing.T, include1, include2, exclude string) {
		// Build query with source filtering - should not panic
		query := es.NewQuery(es.Bool()).
			SourceIncludes(include1, include2).
			SourceExcludes(exclude)

		// Marshal to JSON - should not panic
		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		// Unmarshal back
		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Exists_Query(f *testing.F) {
	f.Add("field_name")
	f.Add("nested.field")
	f.Add("")
	f.Add("special!@#$%")

	f.Fuzz(func(t *testing.T, field string) {
		query := es.Exists(field).Boost(1.5)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Terms_Query(f *testing.F) {
	f.Add("category", "books", "electronics", "clothing")
	f.Add("status", "active", "pending", "")
	f.Add("", "", "", "")

	f.Fuzz(func(t *testing.T, field, val1, val2, val3 string) {
		query := es.Terms(field, val1, val2, val3).Boost(1.0)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Nested_Query(f *testing.F) {
	f.Add("comments", "user", "admin", true)
	f.Add("items", "status", "active", false)
	f.Add("", "", "", true)

	f.Fuzz(func(t *testing.T, path, field, value string, ignoreUnmapped bool) {
		query := es.Nested(path, es.Term(field, value)).
			ScoreMode(ScoreMode.Avg).
			IgnoreUnmapped(ignoreUnmapped)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Query_String(f *testing.F) {
	f.Add("foo AND bar", "title", true)
	f.Add("status:active", "content", false)
	f.Add("", "", true)
	f.Add("special!@#$% OR test", "field.keyword", false)

	f.Fuzz(func(t *testing.T, queryStr, defaultField string, analyzeWildcard bool) {
		query := es.QueryString(queryStr).
			DefaultField(defaultField).
			AnalyzeWildcard(analyzeWildcard).
			Boost(1.5)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_IDs_Query(f *testing.F) {
	f.Add("doc1", "doc2", "doc3")
	f.Add("123", "456", "789")
	f.Add("", "", "")

	f.Fuzz(func(t *testing.T, id1, id2, id3 string) {
		query := es.IDs(id1, id2, id3).Boost(1.0)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Regexp_Query(f *testing.F) {
	f.Add("endpoint", "/api/.*", true)
	f.Add("path", "^/users/[0-9]+$", false)
	f.Add("", "", true)

	f.Fuzz(func(t *testing.T, field, pattern string, caseInsensitive bool) {
		query := es.Regexp(field, pattern).
			CaseInsensitive(caseInsensitive).
			Boost(1.0)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Range_Query_With_Format(f *testing.F) {
	f.Add("date", "2024-01-01", "2024-12-31", "yyyy-MM-dd")
	f.Add("timestamp", "now-1d", "now", "epoch_millis")
	f.Add("", "", "", "")

	f.Fuzz(func(t *testing.T, field, gte, lte, format string) {
		query := es.Range(field).
			GreaterThanOrEqual(gte).
			LessThanOrEqual(lte).
			Format(format).
			Boost(1.0)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Aggregation_Terms(f *testing.F) {
	f.Add("category.keyword", 10, 5)
	f.Add("status", 100, 1)
	f.Add("", 0, 0)

	f.Fuzz(func(t *testing.T, field string, size, minDocCount int) {
		agg := es.TermsAgg(field).
			Size(size).
			MinDocCount(minDocCount)

		result, err := json.Marshal(agg)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Match_Query_With_Options(f *testing.F) {
	f.Add("title", "search query", "AUTO", 50)
	f.Add("content", "full text search", "2", 100)
	f.Add("", "", "", 0)

	f.Fuzz(func(t *testing.T, field, query, fuzziness string, maxExpansions int) {
		matchQuery := es.Match(field, query).
			Fuzziness(fuzziness).
			MaxExpansions(maxExpansions).
			Lenient(true)

		result, err := json.Marshal(matchQuery)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Bool_Query_With_Boost(f *testing.F) {
	f.Add("field1", "value1", 1.5, true)
	f.Add("status", "active", 2.0, false)
	f.Add("", "", 0.0, true)

	f.Fuzz(func(t *testing.T, field, value string, boost float64, adjustPureNegative bool) {
		query := es.Bool().
			Must(es.Term(field, value)).
			Boost(boost).
			AdjustPureNegative(adjustPureNegative)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Deeply_Nested_Bool_Query(f *testing.F) {
	f.Add("f1", "v1", "f2", "v2", "f3", "v3")
	f.Add("status", "active", "type", "product", "category", "electronics")
	f.Add("", "", "", "", "", "")

	f.Fuzz(func(t *testing.T, f1, v1, f2, v2, f3, v3 string) {
		query := es.Bool().
			Must(
				es.Bool().
					Should(
						es.Term(f1, v1),
						es.Match(f2, v2),
					).
					MinimumShouldMatch(1),
			).
			Filter(
				es.Bool().
					Must(es.Term(f3, v3)).
					MustNot(es.Exists(f1)),
			)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}

func Fuzz_Full_Search_Query(f *testing.F) {
	f.Add("title", "search", "status", "active", 10, 0, true)
	f.Add("content", "query builder", "type", "document", 100, 50, false)
	f.Add("", "", "", "", 0, 0, true)

	f.Fuzz(func(t *testing.T, searchField, searchValue, filterField, filterValue string, size, from int, trackHits bool) {
		query := es.NewQuery(
			es.Bool().
				Must(es.Match(searchField, searchValue)).
				Filter(es.Term(filterField, filterValue)),
		).
			Size(size).
			From(from).
			TrackTotalHits(trackHits).
			SourceIncludes(searchField, filterField)

		result, err := json.Marshal(query)
		if err != nil {
			return
		}

		var unmarshaled map[string]any
		if err = json.Unmarshal(result, &unmarshaled); err != nil {
			t.Errorf("produced invalid JSON: %v", err)
		}
	})
}
