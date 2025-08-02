class ESQueryParser {
    constructor() {
        this.indent = 0;
    }

    parse(esQuery) {
        try {
            let code;
            this.indent = 1;

            // Eğer sadece aggregation varsa NewAggs() kullan
            if ((esQuery.aggs || esQuery.aggregations) && !esQuery.query) {
                code = 'es.NewAggs(\n';
                const aggs = esQuery.aggs || esQuery.aggregations;
                for (const [name, agg] of Object.entries(aggs)) {
                    code += `${this.getIndent()}es.Agg("${name}", ${this.parseAggregation(agg)}),\n`;
                }
                this.indent--;
                code += ')';

                // Top level parametreler (aggregation'ları hariç tut)
                code += this.parseTopLevelParams(esQuery, true);
            } else {
                // Query varsa veya hem query hem aggregation varsa NewQuery() kullan
                code = 'es.NewQuery(\n';
                // Ana query gövdesi
                if (esQuery.query) {
                    code += `${this.getIndent()}${this.parseQueryBody(esQuery.query)},\n`;
                }
                this.indent--;
                code += ')';

                // Top level parametreler (aggregation'lar dahil)
                code += this.parseTopLevelParams(esQuery, false);
            }

            return code;
        } catch (error) {
            console.error('Parsing error:', error);
            throw new Error(`Query parsing failed: ${error.message}`);
        }
    }

    // Check if query has top-level parameters (sort, size, from, etc.)
    hasTopLevelParams(query, skipAggs = false) {
        return !!(query.sort || query.size !== undefined || query.from !== undefined ||
          query._source !== undefined || query.track_total_hits !== undefined ||
          query.min_score !== undefined ||
          (!skipAggs && (query.aggs || query.aggregations)));
    }

    // Check if query is simple (non-bool, single clause)
    isSimpleQuery(query) {
        if (!query) return true;

        // Bool queries are never simple
        if (query.bool) return false;

        // Check if it's a single clause query
        const queryTypes = Object.keys(query);
        return queryTypes.length === 1;
    }

    parseTopLevelParams(query, skipAggs = false) {
        let code = '';

        // Sort
        if (query.sort) {
            code += '.Sort(\n';
            this.indent++;
            const sortParts = query.sort.map(sortItem => {
                const field = Object.keys(sortItem)[0];
                const options = sortItem[field];
                let sortCode = `${this.getIndent()}es.Sort("${field}")`;
                if (options.order) {
                    sortCode += `.Order(order.${this.capitalize(options.order)})`;
                }
                if (options.mode) {
                    sortCode += `.Mode("${options.mode}")`;
                }
                return sortCode;
            });
            code += sortParts.join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // Aggregations - sadece skipAggs false ise ekle
        if (!skipAggs && (query.aggs || query.aggregations)) {
            const aggs = query.aggs || query.aggregations;
            code += '.Aggs(\n';
            this.indent++;
            const aggParts = Object.entries(aggs).map(([name, agg]) => {
                return `${this.getIndent()}es.Agg("${name}", ${this.parseAggregation(agg)})`;
            });
            code += aggParts.join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // Source Filtering
        if (query._source !== undefined) {
            if (Array.isArray(query._source)) {
                code += `.\n${this.getIndent()}SourceIncludes(${query._source.map(field => `"${field}"`).join(', ')})`;
            } else if (typeof query._source === 'boolean') {
                code += `.\n${this.getIndent()}Source${query._source ? 'True' : 'False'}()`;
            } else if (typeof query._source === 'object') {
                if (query._source.includes) {
                    code += `.\n${this.getIndent()}SourceIncludes(${query._source.includes.map(field => `"${field}"`).join(', ')})`;
                }
                if (query._source.excludes) {
                    code += `.\n${this.getIndent()}SourceExcludes(${query._source.excludes.map(field => `"${field}"`).join(', ')})`;
                }
            }
        }

        // Pagination
        if (query.size !== undefined) {
            code += `.\n${this.getIndent()}Size(${query.size})`;
        }
        if (query.from !== undefined) {
            code += `.\n${this.getIndent()}From(${query.from})`;
        }

        // Other Parameters
        if (query.track_total_hits !== undefined) {
            code += `.\n${this.getIndent()}TrackTotalHits(${query.track_total_hits})`;
        }
        if (query.min_score !== undefined) {
            code += `.\n${this.getIndent()}MinScore(${query.min_score})`;
        }

        return code;
    }

    parseQueryBody(query) {
        if (query.bool) return this.parseBoolQuery(query.bool);
        if (query.term) return this.parseTerm(query.term);
        if (query.terms) return this.parseTerms(query.terms);
        if (query.match) return this.parseMatch(query.match);
        if (query.match_all) return this.parseMatchAll(query.match_all);
        if (query.match_none) return this.parseMatchNone(query.match_none);
        if (query.match_bool_prefix) return this.parseMatchBoolPrefix(query.match_bool_prefix);
        if (query.match_phrase) return this.parseMatchPhrase(query.match_phrase);
        if (query.match_phrase_prefix) return this.parseMatchPhrasePrefix(query.match_phrase_prefix);
        if (query.range) return this.parseRange(query.range);
        if (query.exists) return this.parseExists(query.exists);
        if (query.nested) return this.parseNested(query.nested);
        if (query.query_string) return this.parseQueryString(query.query_string);
        if (query.simple_query_string) return this.parseSimpleQueryString(query.simple_query_string);
        if (query.multi_match) return this.parseMultiMatch(query.multi_match);

        throw new Error(`Unsupported query type: ${Object.keys(query)}`);
    }

    parseBoolQuery(bool) {
        let code = `${this.getIndent()}es.Bool()`;

        // Must
        if (bool.must) {
            code += '.Must(\n';
            this.indent++;
            code += bool.must.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // Should
        if (bool.should) {
            code += '.Should(\n';
            this.indent++;
            code += bool.should.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // Filter
        if (bool.filter) {
            code += '.Filter(\n';
            this.indent++;
            code += bool.filter.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // Must Not
        if (bool.must_not) {
            code += '.MustNot(\n';
            this.indent++;
            code += bool.must_not.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        // MinimumShouldMatch
        if (bool.minimum_should_match !== undefined) {
            code += `.MinimumShouldMatch(${bool.minimum_should_match})`;
        }

        // Boost
        if (bool.boost) {
            code += `.Boost(${bool.boost})`;
        }

        return code;
    }

    parseTerm(term) {
        const field = Object.keys(term)[0];
        const value = term[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.Term("${field}", ${this.formatValue(value.value)})`;
            if (value.boost) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Term("${field}", ${this.formatValue(value)})`;
    }

    formatValue(value) {
        if (typeof value === 'string') {
            return `"${value}"`;
        }
        return value;
    }

    parseTerms(terms) {
        const field = Object.keys(terms)[0];
        const values = terms[field];

        // Go library individual parameters kullanıyor: es.Terms("field", "value1", "value2")
        if (Array.isArray(values)) {
            const formattedValues = values.map(item => {
                if (typeof item === 'string') {
                    return `"${item}"`;
                }
                return item;
            }).join(', ');

            return `${this.getIndent()}es.Terms("${field}", ${formattedValues})`;
        }

        // Single value case
        return `${this.getIndent()}es.Terms("${field}", ${this.formatValue(values)})`;
    }

    formatArray(arr) {
        if (!Array.isArray(arr)) {
            return JSON.stringify(arr);
        }

        const values = arr.map(item => {
            if (typeof item === 'string') {
                return `"${item}"`;
            }
            return item;
        }).join(', ');

        return `{${values}}`;
    }

    determineArrayType(arr) {
        if (!Array.isArray(arr) || arr.length === 0) {
            return '';
        }

        // Tüm elemanların tipini kontrol et
        const types = new Set(arr.map(item => typeof item));

        // Tek tip varsa
        if (types.size === 1) {
            const type = types.values().next().value;
            switch (type) {
                case 'string':
                    return '[]string';
                case 'number':
                    // Tüm sayıların tam sayı olup olmadığını kontrol et
                    if (arr.every(num => Number.isInteger(num))) {
                        return '[]int';
                    }
                    return '[]float64';
                case 'boolean':
                    return '[]bool';
                default:
                    return '';
            }
        }

        // Karışık tipler varsa
        return '';
    }

    parseMatch(match) {
        const field = Object.keys(match)[0];
        const value = match[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.Match("${field}", "${value.query}")`;
            if (value.operator) {
                code += `.Operator("${value.operator.toLowerCase()}")`;
            }
            if (value.boost) {
                code += `.Boost(${value.boost})`;
            }
            if (value.fuzziness) {
                code += `.Fuzziness("${value.fuzziness}")`;
            }
            if (value.zero_terms_query) {
                code += `.ZeroTermsQuery("${value.zero_terms_query}")`;
            }
            if (value.auto_generate_synonyms_phrase_query !== undefined) {
                code += `.AutoGenerateSynonymsPhraseQuery(${value.auto_generate_synonyms_phrase_query})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Match("${field}", "${value}")`;
    }

    parseMatchBoolPrefix(matchBoolPrefix) {
        const field = Object.keys(matchBoolPrefix)[0];
        const value = matchBoolPrefix[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.MatchBoolPrefix("${field}", "${value.query}")`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.minimum_should_match) {
                code += `.MinimumShouldMatch("${value.minimum_should_match}")`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchBoolPrefix("${field}", "${value}")`;
    }

    parseMatchPhrase(matchPhrase) {
        const field = Object.keys(matchPhrase)[0];
        const value = matchPhrase[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.MatchPhrase("${field}", "${value.query}")`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.boost) {
                code += `.Boost(${value.boost})`;
            }
            if (value.zero_terms_query) {
                code += `.ZeroTermsQuery("${value.zero_terms_query}")`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchPhrase("${field}", "${value}")`;
    }

    parseMatchPhrasePrefix(matchPhrasePrefix) {
        const field = Object.keys(matchPhrasePrefix)[0];
        const value = matchPhrasePrefix[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.MatchPhrasePrefix("${field}", "${value.query}")`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.max_expansions) {
                code += `.MaxExpansions(${value.max_expansions})`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchPhrasePrefix("${field}", "${value}")`;
    }

    parseMultiMatch(multiMatch) {
        let code = `${this.getIndent()}es.MultiMatch("${multiMatch.query}")`;

        if (multiMatch.fields) {
            const fieldsStr = multiMatch.fields.map(field => `"${field}"`).join(', ');
            code += `.Fields(${fieldsStr})`;
        }
        if (multiMatch.type) {
            code += `.Type("${multiMatch.type}")`;
        }
        if (multiMatch.operator) {
            code += `.Operator("${multiMatch.operator.toLowerCase()}")`;
        }
        if (multiMatch.minimum_should_match) {
            code += `.MinimumShouldMatch("${multiMatch.minimum_should_match}")`;
        }
        if (multiMatch.tie_breaker !== undefined) {
            code += `.TieBreaker(${multiMatch.tie_breaker})`;
        }
        if (multiMatch.analyzer) {
            code += `.Analyzer("${multiMatch.analyzer}")`;
        }
        if (multiMatch.boost) {
            code += `.Boost(${multiMatch.boost})`;
        }
        if (multiMatch.fuzziness) {
            code += `.Fuzziness("${multiMatch.fuzziness}")`;
        }
        if (multiMatch.fuzzy_transpositions !== undefined) {
            code += `.FuzzyTranspositions(${multiMatch.fuzzy_transpositions})`;
        }
        if (multiMatch.lenient !== undefined) {
            code += `.Lenient(${multiMatch.lenient})`;
        }
        if (multiMatch.prefix_length !== undefined) {
            code += `.PrefixLength(${multiMatch.prefix_length})`;
        }
        if (multiMatch.max_expansions !== undefined) {
            code += `.MaxExpansions(${multiMatch.max_expansions})`;
        }
        if (multiMatch.zero_terms_query) {
            code += `.ZeroTermsQuery("${multiMatch.zero_terms_query}")`;
        }
        if (multiMatch.auto_generate_synonyms_phrase_query !== undefined) {
            code += `.AutoGenerateSynonymsPhraseQuery(${multiMatch.auto_generate_synonyms_phrase_query})`;
        }

        return code;
    }

    parseRange(range) {
        const field = Object.keys(range)[0];
        const conditions = range[field];
        let code = `${this.getIndent()}es.Range("${field}")`;

        if (conditions.gt !== undefined) {
            code += `.GreaterThan(${JSON.stringify(conditions.gt)})`;
        }
        if (conditions.gte !== undefined) {
            code += `.GreaterThanOrEqual(${JSON.stringify(conditions.gte)})`;
        }
        if (conditions.lt !== undefined) {
            code += `.LessThan(${JSON.stringify(conditions.lt)})`;
        }
        if (conditions.lte !== undefined) {
            code += `.LessThanOrEqual(${JSON.stringify(conditions.lte)})`;
        }
        if (conditions.boost !== undefined) {
            code += `.Boost(${conditions.boost})`;
        }
        if (conditions.relation) {
            code += `.Relation(RangeRelation.${this.capitalize(conditions.relation)})`;
        }
        if (conditions.time_zone) {
            code += `.TimeZone("${conditions.time_zone}")`;
        }
        if (conditions.format) {
            code += `.Format("${conditions.format}")`;
        }

        return code;
    }

    parseExists(exists) {
        return `${this.getIndent()}es.Exists("${exists.field}")`;
    }

    parseNested(nested) {
        let code = `${this.getIndent()}es.Nested("${nested.path}", ${this.parseQueryBody(nested.query)})`;

        if (nested.score_mode) {
            code += `.ScoreMode(ScoreMode.${this.capitalize(nested.score_mode)})`;
        }
        if (nested.inner_hits) {
            code += `.InnerHits(${this.parseInnerHits(nested.inner_hits)})`;
        }

        return code;
    }

    parseInnerHits(innerHits) {
        let code = 'es.InnerHits()';

        if (innerHits.from !== undefined) {
            code += `.From(${innerHits.from})`;
        }
        if (innerHits.size !== undefined) {
            code += `.Size(${innerHits.size})`;
        }
        if (innerHits.name) {
            code += `.Name("${innerHits.name}")`;
        }
        if (innerHits.sort) {
            code += `.Sort(${this.parseSort(innerHits.sort)})`;
        }

        return code;
    }

    parseQueryString(qs) {
        let code = `${this.getIndent()}es.QueryString("${qs.query}")`;

        if (qs.default_field) {
            code += `.DefaultField("${qs.default_field}")`;
        }
        if (qs.fields) {
            code += `.Fields(${JSON.stringify(qs.fields)})`;
        }
        if (qs.type) {
            code += `.Type("${qs.type}")`;
        }
        if (qs.tie_breaker !== undefined) {
            code += `.TieBreaker(${qs.tie_breaker})`;
        }
        if (qs.default_operator) {
            code += `.DefaultOperator("${qs.default_operator.toUpperCase()}")`;
        }
        if (qs.allow_leading_wildcard !== undefined) {
            code += `.AllowLeadingWildcard(${qs.allow_leading_wildcard})`;
        }
        if (qs.fuzzy_max_expansions !== undefined) {
            code += `.FuzzyMaxExpansions(${qs.fuzzy_max_expansions})`;
        }
        if (qs.fuzziness) {
            code += `.Fuzziness("${qs.fuzziness}")`;
        }
        if (qs.fuzzy_transpositions !== undefined) {
            code += `.FuzzyTranspositions(${qs.fuzzy_transpositions})`;
        }
        if (qs.lenient !== undefined) {
            code += `.Lenient(${qs.lenient})`;
        }
        if (qs.max_determinized_states !== undefined) {
            code += `.MaxDeterminizedStates(${qs.max_determinized_states})`;
        }
        if (qs.minimum_should_match) {
            code += `.MinimumShouldMatch("${qs.minimum_should_match}")`;
        }
        if (qs.quote_field_suffix) {
            code += `.QuoteFieldSuffix("${qs.quote_field_suffix}")`;
        }
        if (qs.phrase_slop !== undefined) {
            code += `.PhraseSlop(${qs.phrase_slop})`;
        }
        if (qs.boost !== undefined) {
            code += `.Boost(${qs.boost})`;
        }
        if (qs.auto_generate_synonyms_phrase_query !== undefined) {
            code += `.AutoGenerateSynonymsPhraseQuery(${qs.auto_generate_synonyms_phrase_query})`;
        }
        if (qs.analyze_wildcard !== undefined) {
            code += `.AnalyzeWildcard(${qs.analyze_wildcard})`;
        }
        if (qs.rewrite) {
            code += `.Rewrite("${qs.rewrite}")`;
        }
        if (qs.fuzzy_rewrite) {
            code += `.FuzzyRewrite("${qs.fuzzy_rewrite}")`;
        }

        return code;
    }

    parseSimpleQueryString(sqs) {
        let code = `${this.getIndent()}es.SimpleQueryString("${sqs.query}")`;

        if (sqs.fields) {
            const fieldsStr = sqs.fields.map(field => `"${field}"`).join(', ');
            code += `.Fields([]string{${fieldsStr}})`;
        }
        if (sqs.default_operator) {
            code += `.DefaultOperator("${sqs.default_operator.toUpperCase()}")`;
        }
        if (sqs.analyzer) {
            code += `.Analyzer("${sqs.analyzer}")`;
        }
        if (sqs.flags) {
            code += `.Flags("${sqs.flags}")`;
        }
        if (sqs.fuzzy_max_expansions !== undefined) {
            code += `.FuzzyMaxExpansions(${sqs.fuzzy_max_expansions})`;
        }
        if (sqs.fuzzy_prefix_length !== undefined) {
            code += `.FuzzyPrefixLength(${sqs.fuzzy_prefix_length})`;
        }
        if (sqs.fuzzy_transpositions !== undefined) {
            code += `.FuzzyTranspositions(${sqs.fuzzy_transpositions})`;
        }
        if (sqs.lenient !== undefined) {
            code += `.Lenient(${sqs.lenient})`;
        }
        if (sqs.minimum_should_match) {
            code += `.MinimumShouldMatch("${sqs.minimum_should_match}")`;
        }
        if (sqs.quote_field_suffix) {
            code += `.QuoteFieldSuffix("${sqs.quote_field_suffix}")`;
        }
        if (sqs.analyze_wildcard !== undefined) {
            code += `.AnalyzeWildcard(${sqs.analyze_wildcard})`;
        }
        if (sqs.auto_generate_synonyms_phrase_query !== undefined) {
            code += `.AutoGenerateSynonymsPhraseQuery(${sqs.auto_generate_synonyms_phrase_query})`;
        }

        return code;
    }

    parseMatchAll(matchAll) {
        let code = `${this.getIndent()}es.MatchAll()`;
        if (matchAll && matchAll.boost !== undefined) {
            code += `.Boost(${matchAll.boost})`;
        }
        return code;
    }

    parseMatchNone(matchNone) {
        if (!matchNone || Object.keys(matchNone).length === 0) {
            return `${this.getIndent()}es.MatchNone()`;
        }

        // Get the field and value from match_none structure
        const field = Object.keys(matchNone)[0];
        const value = matchNone[field];

        if (value && value.query !== undefined) {
            return `${this.getIndent()}es.MatchNone("${field}", ${this.formatValue(value.query)})`;
        }

        return `${this.getIndent()}es.MatchNone("${field}", ${this.formatValue(value)})`;
    }

    parseAggregation(agg) {
        // Terms aggregation
        if (agg.terms) {
            let code = `es.TermsAgg("${agg.terms.field}")`;
            if (agg.terms.size !== undefined) {
                code += `.\n${this.getIndent()}Size(${agg.terms.size})`;
            }
            if (agg.terms.shard_size !== undefined) {
                code += `.\n${this.getIndent()}ShardSize(${agg.terms.shard_size})`;
            }
            if (agg.terms.min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}MinDocCount(${agg.terms.min_doc_count})`;
            }
            if (agg.terms.shard_min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}ShardMinDocCount(${agg.terms.shard_min_doc_count})`;
            }
            if (agg.terms.show_term_doc_count_error !== undefined) {
                code += `.\n${this.getIndent()}ShowTermDocCountError(${agg.terms.show_term_doc_count_error})`;
            }
            if (agg.terms.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.terms.missing)})`;
            }
            if (agg.terms.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.terms.script)})`;
            }
            if (agg.terms.format) {
                code += `.\n${this.getIndent()}Format("${agg.terms.format}")`;
            }
            if (agg.terms.include) {
                if (Array.isArray(agg.terms.include)) {
                    code += `.\n${this.getIndent()}Include(${agg.terms.include.map(v => `"${v}"`).join(", ")})`;
                } else {
                    code += `.\n${this.getIndent()}Include("${agg.terms.include}")`;
                }
            }
            if (agg.terms.exclude) {
                if (Array.isArray(agg.terms.exclude)) {
                    code += `.\n${this.getIndent()}Exclude(${agg.terms.exclude.map(v => `"${v}"`).join(", ")})`;
                } else {
                    code += `.\n${this.getIndent()}Exclude("${agg.terms.exclude}")`;
                }
            }
            if (agg.terms.execution_hint) {
                code += `.\n${this.getIndent()}ExecutionHint(ExecutionHint.${this.capitalize(agg.terms.execution_hint)})`;
            }
            if (agg.terms.collect_mode) {
                code += `.\n${this.getIndent()}CollectMode(CollectMode.${this.capitalize(agg.terms.collect_mode)})`;
            }
            if (agg.terms.meta) {
                for (const [key, value] of Object.entries(agg.terms.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            if (agg.terms.order) {
                for (const [field, order] of Object.entries(agg.terms.order)) {
                    code += `.\n${this.getIndent()}Order("${field}", Order.${this.capitalize(order)})`;
                }
            }
            // İç içe aggregation'ları ekle
            if (agg.aggs || agg.aggregations) {
                code += `.\n${this.getIndent()}Aggs(\n`;
                this.indent++;
                const subAggs = agg.aggs || agg.aggregations;
                const aggParts = Object.entries(subAggs).map(([name, subAgg]) => {
                    return `${this.getIndent()}es.Agg("${name}", ${this.parseAggregation(subAgg)})`;
                });
                code += aggParts.join(',\n');
                code += ',\n';
                this.indent--;
                code += `${this.getIndent()})`;
            }
            return code;
        }

        // Stats aggregation
        if (agg.stats) {
            let code = `es.StatsAgg("${agg.stats.field}")`;
            if (agg.stats.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.stats.missing)})`;
            }
            if (agg.stats.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.stats.script)})`;
            }
            if (agg.stats.format) {
                code += `.\n${this.getIndent()}Format("${agg.stats.format}")`;
            }
            if (agg.stats.meta) {
                for (const [key, value] of Object.entries(agg.stats.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Extended Stats aggregation
        if (agg.extended_stats) {
            let code = `es.ExtendedStatsAgg("${agg.extended_stats.field}")`;
            if (agg.extended_stats.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.extended_stats.missing)})`;
            }
            if (agg.extended_stats.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.extended_stats.script)})`;
            }
            if (agg.extended_stats.format) {
                code += `.\n${this.getIndent()}Format("${agg.extended_stats.format}")`;
            }
            if (agg.extended_stats.meta) {
                for (const [key, value] of Object.entries(agg.extended_stats.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Min aggregation
        if (agg.min) {
            let code = `es.MinAgg("${agg.min.field}")`;
            if (agg.min.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.min.missing)})`;
            }
            if (agg.min.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.min.script)})`;
            }
            if (agg.min.format) {
                code += `.\n${this.getIndent()}Format("${agg.min.format}")`;
            }
            if (agg.min.meta) {
                for (const [key, value] of Object.entries(agg.min.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Max aggregation
        if (agg.max) {
            let code = `es.MaxAgg("${agg.max.field}")`;
            if (agg.max.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.max.missing)})`;
            }
            if (agg.max.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.max.script)})`;
            }
            if (agg.max.format) {
                code += `.\n${this.getIndent()}Format("${agg.max.format}")`;
            }
            if (agg.max.meta) {
                for (const [key, value] of Object.entries(agg.max.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Avg aggregation
        if (agg.avg) {
            let code = `es.AvgAgg("${agg.avg.field}")`;
            if (agg.avg.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.avg.missing)})`;
            }
            if (agg.avg.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.avg.script)})`;
            }
            if (agg.avg.format) {
                code += `.\n${this.getIndent()}Format("${agg.avg.format}")`;
            }
            if (agg.avg.meta) {
                for (const [key, value] of Object.entries(agg.avg.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Sum aggregation
        if (agg.sum) {
            let code = `es.SumAgg("${agg.sum.field}")`;
            if (agg.sum.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.sum.missing)})`;
            }
            if (agg.sum.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.sum.script)})`;
            }
            if (agg.sum.format) {
                code += `.\n${this.getIndent()}Format("${agg.sum.format}")`;
            }
            if (agg.sum.meta) {
                for (const [key, value] of Object.entries(agg.sum.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Cardinality aggregation
        if (agg.cardinality) {
            let code = `es.CardinalityAgg("${agg.cardinality.field}")`;
            if (agg.cardinality.precision_threshold !== undefined) {
                code += `.\n${this.getIndent()}PrecisionThreshold(${agg.cardinality.precision_threshold})`;
            }
            if (agg.cardinality.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.cardinality.missing)})`;
            }
            if (agg.cardinality.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.cardinality.script)})`;
            }
            if (agg.cardinality.meta) {
                for (const [key, value] of Object.entries(agg.cardinality.meta)) {
                    code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
                }
            }
            return code;
        }

        // Multi Terms aggregation
        if (agg.multi_terms) {
            let code = `es.MultiTermsAgg(${agg.multi_terms.terms.map(term => `es.TermAgg("${term.field}")`).join(", ")})`;
            if (agg.multi_terms.size !== undefined) {
                code += `.\n${this.getIndent()}Size(${agg.multi_terms.size})`;
            }
            if (agg.multi_terms.shard_size !== undefined) {
                code += `.\n${this.getIndent()}ShardSize(${agg.multi_terms.shard_size})`;
            }
            if (agg.multi_terms.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.multi_terms.missing)})`;
            }
            if (agg.multi_terms.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.multi_terms.script)})`;
            }
            if (agg.multi_terms.ignore_unmapped !== undefined) {
                code += `.\n${this.getIndent()}IgnoreUnmapped(${agg.multi_terms.ignore_unmapped})`;
            }
            if (agg.multi_terms.include) {
                if (Array.isArray(agg.multi_terms.include)) {
                    code += `.\n${this.getIndent()}Include(${agg.multi_terms.include.map(v => `"${v}"`).join(", ")})`;
                } else {
                    code += `.\n${this.getIndent()}Include("${agg.multi_terms.include}")`;
                }
            }
            if (agg.multi_terms.exclude) {
                if (Array.isArray(agg.multi_terms.exclude)) {
                    code += `.\n${this.getIndent()}Exclude(${agg.multi_terms.exclude.map(v => `"${v}"`).join(", ")})`;
                } else {
                    code += `.\n${this.getIndent()}Exclude("${agg.multi_terms.exclude}")`;
                }
            }
            if (agg.multi_terms.min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}MinDocCount(${agg.multi_terms.min_doc_count})`;
            }
            if (agg.multi_terms.execution_hint) {
                code += `.\n${this.getIndent()}ExecutionHint(ExecutionHint.${this.capitalize(agg.multi_terms.execution_hint)})`;
            }
            if (agg.multi_terms.collect_mode) {
                code += `.\n${this.getIndent()}CollectMode(CollectMode.${this.capitalize(agg.multi_terms.collect_mode)})`;
            }
            if (agg.multi_terms.order) {
                for (const [field, order] of Object.entries(agg.multi_terms.order)) {
                    code += `.\n${this.getIndent()}Order("${field}", Order.${this.capitalize(order)})`;
                }
            }
            return code;
        }

        // Nested aggregation
        if (agg.nested) {
            let code = `es.NestedAgg("${agg.nested.path}")`;
            if (agg.aggs || agg.aggregations) {
                code += `.\n${this.getIndent()}Aggs(\n`;
                this.indent++;
                const subAggs = agg.aggs || agg.aggregations;
                const aggParts = Object.entries(subAggs).map(([name, subAgg]) => {
                    return `${this.getIndent()}es.Agg("${name}", ${this.parseAggregation(subAgg)})`;
                });
                code += aggParts.join(',\n');
                code += ',\n';
                this.indent--;
                code += `${this.getIndent()})`;
            }
            return code;
        }

        throw new Error(`Unsupported aggregation type: ${Object.keys(agg)}`);
    }

    getIndent() {
        return '\t'.repeat(this.indent);
    }

    capitalize(str) {
        return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
    }

    // Helper method to format values correctly
    formatValue(value) {
        if (typeof value === 'string') {
            return `"${value}"`;
        }
        return value;
    }

    // Helper method to parse script objects
    parseScript(script) {
        if (typeof script === 'string') {
            return `es.ScriptSource("${script}", ScriptLanguage.Painless)`;
        }

        let scriptCode = '';
        if (script.source) {
            scriptCode = `es.ScriptSource("${script.source}"`;
            if (script.lang) {
                scriptCode += `, ScriptLanguage.${this.capitalize(script.lang)}`;
            }
            scriptCode += ')';
        } else if (script.id) {
            scriptCode = `es.ScriptId("${script.id}"`;
            if (script.lang) {
                scriptCode += `, ScriptLanguage.${this.capitalize(script.lang)}`;
            }
            scriptCode += ')';
        }

        if (script.params) {
            scriptCode += `.Params(${JSON.stringify(script.params)})`;
        }

        return scriptCode;
    }
}

// Node.js export
if (typeof module !== 'undefined' && module.exports) {
    module.exports = ESQueryParser;
}