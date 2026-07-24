class ESQueryParser {
    constructor() {
        this.indent = 0;

        // JSON value (lower/snake) → Go enum const name
        this.enumMaps = {
            Operator: {
                and: 'And',
                or: 'Or',
            },
            ZeroTermsQuery: {
                all: 'All',
                none: 'None',
            },
            TextQueryType: {
                best_fields: 'Bestfields',
                most_fields: 'Mostfields',
                cross_fields: 'Crossfields',
                phrase: 'Phrase',
                phrase_prefix: 'Phraseprefix',
                bool_prefix: 'Boolprefix',
            },
            RangeRelation: {
                within: 'Within',
                contains: 'Contains',
                intersects: 'Intersects',
            },
            ScoreMode: {
                avg: 'Avg',
                max: 'Max',
                min: 'Min',
                none: 'None',
                sum: 'Sum',
            },
            BoostMode: {
                multiply: 'Multiply',
                replace: 'Replace',
                sum: 'Sum',
                avg: 'Avg',
                max: 'Max',
                min: 'Min',
            },
            Modifier: {
                none: 'None',
                log: 'Log',
                log1p: 'Log1p',
                log2p: 'Log2p',
                ln: 'Ln',
                ln1p: 'Ln1p',
                ln2p: 'Ln2p',
                square: 'Square',
                sqrt: 'Sqrt',
                reciprocal: 'Reciprocal',
            },
            MultiValuesMode: {
                min: 'Min',
                max: 'Max',
                avg: 'Avg',
                sum: 'Sum',
            },
            ScriptLanguage: {
                painless: 'Painless',
                expression: 'Expression',
                mustache: 'Mustache',
                java: 'Java',
            },
            Order: {
                asc: 'Asc',
                desc: 'Desc',
                _default: 'Default',
            },
            Mode: {
                min: 'Min',
                max: 'Max',
                sum: 'Sum',
                avg: 'Avg',
                median: 'Median',
                _default: 'Default',
            },
            ExecutionHint: {
                map: 'Map',
                global_ordinals: 'GlobalOrdinals',
                fielddata: 'FieldData',
            },
            CollectMode: {
                breadth_first: 'BreadthFirst',
                depth_first: 'DepthFirst',
            },
            HighlighterType: {
                unified: 'Unified',
                plain: 'Plain',
                fvh: 'Fvh',
            },
            Fragmenter: {
                simple: 'Simple',
                span: 'Span',
            },
            BoundaryScanner: {
                chars: 'Chars',
                sentence: 'Sentence',
                word: 'Word',
            },
            DistanceType: {
                arc: 'Arc',
                plane: 'Plane',
            },
            ValidationMethod: {
                strict: 'Strict',
                ignore_malformed: 'IgnoreMalformed',
                coerce: 'Coerce',
            },
        };
    }

    parse(esQuery) {
        try {
            let code;
            this.indent = 1;

            if ((esQuery.aggs || esQuery.aggregations) && !esQuery.query) {
                code = 'es.NewAggs(\n';
                const aggs = esQuery.aggs || esQuery.aggregations;
                for (const [name, agg] of Object.entries(aggs)) {
                    code += `${this.getIndent()}es.Agg("${name}", ${this.parseAggregation(agg)}),\n`;
                }
                this.indent--;
                code += ')';
                code += this.parseTopLevelParams(esQuery, true);
            } else {
                code = 'es.NewQuery(\n';
                if (esQuery.query) {
                    code += `${this.getIndent()}${this.parseQueryBody(esQuery.query)},\n`;
                }
                this.indent--;
                code += ')';
                code += this.parseTopLevelParams(esQuery, false);
            }

            return code;
        } catch (error) {
            throw new Error(`Query parsing failed: ${error.message}`);
        }
    }

    formatEnum(pkg, value) {
        if (value === undefined || value === null) {
            throw new Error(`Missing enum value for ${pkg}`);
        }
        const map = this.enumMaps[pkg];
        if (!map) {
            throw new Error(`Unknown enum package: ${pkg}`);
        }
        const key = String(value).toLowerCase();
        const constName = map[key];
        if (!constName) {
            throw new Error(`Unsupported ${pkg} value: ${value}`);
        }
        return `${pkg}.${constName}`;
    }

    formatValue(value) {
        if (typeof value === 'string') {
            return `"${value.replace(/\\/g, '\\\\').replace(/"/g, '\\"')}"`;
        }
        if (value === null) {
            return 'nil';
        }
        if (typeof value === 'boolean' || typeof value === 'number') {
            return String(value);
        }
        return JSON.stringify(value);
    }

    // Go methods that take string but ES JSON may send number (e.g. minimum_should_match: 2)
    formatStringParam(value) {
        if (typeof value === 'string') {
            return this.formatValue(value);
        }
        return `"${String(value)}"`;
    }

    parseGeoPoint(point) {
        if (point === undefined || point === null) {
            throw new Error('Missing geo point');
        }
        if (typeof point === 'object' && !Array.isArray(point)) {
            if (point.lat === undefined || point.lon === undefined) {
                throw new Error('Geo point object requires lat and lon');
            }
            return { lat: point.lat, lon: point.lon };
        }
        if (typeof point === 'string') {
            const parts = point.split(',').map(p => p.trim());
            if (parts.length !== 2) {
                throw new Error(`Unsupported geo point string: ${point}`);
            }
            const lat = parseFloat(parts[0]);
            const lon = parseFloat(parts[1]);
            if (Number.isNaN(lat) || Number.isNaN(lon)) {
                throw new Error(`Unsupported geo point string: ${point}`);
            }
            return { lat, lon };
        }
        if (Array.isArray(point) && point.length === 2) {
            // GeoJSON order: [lon, lat]
            return { lon: point[0], lat: point[1] };
        }
        throw new Error(`Unsupported geo point shape: ${JSON.stringify(point)}`);
    }

    parseSortItem(sortItem) {
        if (typeof sortItem === 'string') {
            return `es.Sort("${sortItem}")`;
        }
        const field = Object.keys(sortItem)[0];
        const options = sortItem[field];
        if (typeof options === 'string') {
            return `es.Sort("${field}").Order(${this.formatEnum('Order', options)})`;
        }
        let sortCode = `es.Sort("${field}")`;
        if (options && typeof options === 'object') {
            if (options.order) {
                sortCode += `.Order(${this.formatEnum('Order', options.order)})`;
            }
            if (options.mode) {
                sortCode += `.Mode(${this.formatEnum('Mode', options.mode)})`;
            }
            if (options.nested) {
                sortCode += `.Nested(${this.parseNestedSort(options.nested)})`;
            }
        }
        return sortCode;
    }

    parseNestedSort(nested) {
        let code = `es.NestedSort("${nested.path}")`;
        if (nested.filter) {
            code += `.Filter(${this.parseQueryBody(nested.filter)})`;
        }
        if (nested.max_children !== undefined) {
            code += `.MaxChildren(${nested.max_children})`;
        }
        if (nested.nested) {
            code += `.Nested(${this.parseNestedSort(nested.nested)})`;
        }
        return code;
    }

    appendSubAggs(code, agg) {
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

    appendMeta(code, meta) {
        if (!meta) return code;
        for (const [key, value] of Object.entries(meta)) {
            code += `.\n${this.getIndent()}Meta("${key}", ${this.formatValue(value)})`;
        }
        return code;
    }

    parseTopLevelParams(query, skipAggs = false) {
        let code = '';

        if (query.sort) {
            code += '.Sort(\n';
            this.indent++;
            const sortParts = query.sort.map(sortItem => {
                return `${this.getIndent()}${this.parseSortItem(sortItem)}`;
            });
            code += sortParts.join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

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

        if (query._source !== undefined) {
            if (Array.isArray(query._source)) {
                code += `.\n${this.getIndent()}SourceIncludes(${query._source.map(field => `"${field}"`).join(', ')})`;
            } else if (typeof query._source === 'boolean') {
                if (!query._source) {
                    code += `.\n${this.getIndent()}SourceFalse()`;
                }
            } else if (typeof query._source === 'object') {
                if (query._source.includes) {
                    code += `.\n${this.getIndent()}SourceIncludes(${query._source.includes.map(field => `"${field}"`).join(', ')})`;
                }
                if (query._source.excludes) {
                    code += `.\n${this.getIndent()}SourceExcludes(${query._source.excludes.map(field => `"${field}"`).join(', ')})`;
                }
            }
        }

        if (query.size !== undefined) {
            code += `.\n${this.getIndent()}Size(${query.size})`;
        }
        if (query.from !== undefined) {
            code += `.\n${this.getIndent()}From(${query.from})`;
        }

        if (query.track_total_hits !== undefined) {
            code += `.\n${this.getIndent()}TrackTotalHits(${query.track_total_hits})`;
        }

        if (query.post_filter) {
            code += `.\n${this.getIndent()}PostFilter(${this.parseQueryBody(query.post_filter)})`;
        }

        if (query.search_after) {
            const values = (Array.isArray(query.search_after) ? query.search_after : [query.search_after])
                .map(v => this.formatValue(v))
                .join(', ');
            code += `.\n${this.getIndent()}SearchAfter(${values})`;
        }

        if (query.highlight) {
            code += `.\n${this.getIndent()}Highlight(\n`;
            this.indent++;
            code += `${this.getIndent()}${this.parseHighlight(query.highlight)},\n`;
            this.indent--;
            code += `${this.getIndent()})`;
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
        if (query.constant_score) return this.parseConstantScore(query.constant_score);
        if (query.function_score) return this.parseFunctionScore(query.function_score);
        if (query.ids) return this.parseIds(query.ids);
        if (query.regexp) return this.parseRegexp(query.regexp);
        if (query.script) return this.parseScriptQuery(query.script);
        if (query.terms_set) return this.parseTermsSet(query.terms_set);
        if (query.fuzzy) return this.parseFuzzy(query.fuzzy);
        if (query.prefix) return this.parsePrefix(query.prefix);
        if (query.wildcard) return this.parseWildcard(query.wildcard);
        if (query.dis_max) return this.parseDisMax(query.dis_max);
        if (query.geo_distance) return this.parseGeoDistance(query.geo_distance);
        if (query.geo_bounding_box) return this.parseGeoBoundingBox(query.geo_bounding_box);

        throw new Error(`Unsupported query type: ${Object.keys(query)}`);
    }

    parseBoolQuery(bool) {
        let code = `${this.getIndent()}es.Bool()`;

        if (bool.must) {
            const clauses = Array.isArray(bool.must) ? bool.must : [bool.must];
            code += '.Must(\n';
            this.indent++;
            code += clauses.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        if (bool.should) {
            const clauses = Array.isArray(bool.should) ? bool.should : [bool.should];
            code += '.Should(\n';
            this.indent++;
            code += clauses.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        if (bool.filter) {
            const clauses = Array.isArray(bool.filter) ? bool.filter : [bool.filter];
            code += '.Filter(\n';
            this.indent++;
            code += clauses.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        if (bool.must_not) {
            const clauses = Array.isArray(bool.must_not) ? bool.must_not : [bool.must_not];
            code += '.MustNot(\n';
            this.indent++;
            code += clauses.map(q => this.parseQueryBody(q)).join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        if (bool.minimum_should_match !== undefined) {
            code += `.MinimumShouldMatch(${this.formatValue(bool.minimum_should_match)})`;
        }

        if (bool.adjust_pure_negative !== undefined) {
            code += `.AdjustPureNegative(${bool.adjust_pure_negative})`;
        }

        if (bool.boost !== undefined) {
            code += `.Boost(${bool.boost})`;
        }

        return code;
    }

    parseTerm(term) {
        const field = Object.keys(term)[0];
        const value = term[field];
        if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
            let code = `${this.getIndent()}es.Term("${field}", ${this.formatValue(value.value)})`;
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            if (value.case_insensitive !== undefined) {
                code += `.CaseInsensitive(${value.case_insensitive})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Term("${field}", ${this.formatValue(value)})`;
    }

    parseTerms(terms) {
        const field = Object.keys(terms)[0];
        const values = terms[field];

        if (Array.isArray(values)) {
            const formattedValues = values.map(item => this.formatValue(item)).join(', ');
            return `${this.getIndent()}es.Terms("${field}", ${formattedValues})`;
        }

        if (typeof values === 'object' && values !== null) {
            const list = values.value !== undefined
                ? (Array.isArray(values.value) ? values.value : [values.value])
                : null;
            if (!list) {
                throw new Error(`Unsupported terms format for field ${field}`);
            }
            const formattedValues = list.map(item => this.formatValue(item)).join(', ');
            let code = `${this.getIndent()}es.Terms("${field}", ${formattedValues})`;
            if (values.boost !== undefined) {
                code += `.Boost(${values.boost})`;
            }
            return code;
        }

        return `${this.getIndent()}es.Terms("${field}", ${this.formatValue(values)})`;
    }

    parseMatch(match) {
        const field = Object.keys(match)[0];
        const value = match[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.Match("${field}", ${this.formatValue(value.query)})`;
            if (value.operator) {
                code += `.Operator(${this.formatEnum('Operator', value.operator)})`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            if (value.cutoff_frequency !== undefined) {
                code += `.CutoffFrequency(${value.cutoff_frequency})`;
            }
            if (value.fuzziness !== undefined) {
                code += `.Fuzziness(${this.formatValue(value.fuzziness)})`;
            }
            if (value.fuzzy_rewrite) {
                code += `.FuzzyRewrite("${value.fuzzy_rewrite}")`;
            }
            if (value.fuzzy_transpositions !== undefined) {
                code += `.FuzzyTranspositions(${value.fuzzy_transpositions})`;
            }
            if (value.lenient !== undefined) {
                code += `.Lenient(${value.lenient})`;
            }
            if (value.max_expansions !== undefined) {
                code += `.MaxExpansions(${value.max_expansions})`;
            }
            if (value.prefix_length !== undefined) {
                code += `.PrefixLength(${value.prefix_length})`;
            }
            if (value.zero_terms_query) {
                code += `.ZeroTermsQuery(${this.formatEnum('ZeroTermsQuery', value.zero_terms_query)})`;
            }
            if (value.auto_generate_synonyms_phrase_query !== undefined) {
                code += `.AutoGenerateSynonymsPhraseQuery(${value.auto_generate_synonyms_phrase_query})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Match("${field}", ${this.formatValue(value)})`;
    }

    parseMatchBoolPrefix(matchBoolPrefix) {
        const field = Object.keys(matchBoolPrefix)[0];
        const value = matchBoolPrefix[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.MatchBoolPrefix("${field}", ${this.formatValue(value.query)})`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.minimum_should_match !== undefined) {
                code += `.MinimumShouldMatch(${this.formatValue(value.minimum_should_match)})`;
            }
            if (value.operator) {
                code += `.Operator(${this.formatEnum('Operator', value.operator)})`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            if (value.fuzziness !== undefined) {
                code += `.Fuzziness(${this.formatValue(value.fuzziness)})`;
            }
            if (value.fuzzy_rewrite) {
                code += `.FuzzyRewrite("${value.fuzzy_rewrite}")`;
            }
            if (value.fuzzy_transpositions !== undefined) {
                code += `.FuzzyTranspositions(${value.fuzzy_transpositions})`;
            }
            if (value.max_expansions !== undefined) {
                code += `.MaxExpansions(${value.max_expansions})`;
            }
            if (value.prefix_length !== undefined) {
                code += `.PrefixLength(${value.prefix_length})`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchBoolPrefix("${field}", ${this.formatValue(value)})`;
    }

    parseMatchPhrase(matchPhrase) {
        const field = Object.keys(matchPhrase)[0];
        const value = matchPhrase[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.MatchPhrase("${field}", ${this.formatValue(value.query)})`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            if (value.zero_terms_query) {
                code += `.ZeroTermsQuery(${this.formatEnum('ZeroTermsQuery', value.zero_terms_query)})`;
            }
            if (value.slop !== undefined) {
                code += `.Slop(${value.slop})`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchPhrase("${field}", ${this.formatValue(value)})`;
    }

    parseMatchPhrasePrefix(matchPhrasePrefix) {
        const field = Object.keys(matchPhrasePrefix)[0];
        const value = matchPhrasePrefix[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.MatchPhrasePrefix("${field}", ${this.formatValue(value.query)})`;
            if (value.analyzer) {
                code += `.Analyzer("${value.analyzer}")`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            if (value.max_expansions !== undefined) {
                code += `.MaxExpansions(${value.max_expansions})`;
            }
            if (value.zero_terms_query) {
                code += `.ZeroTermsQuery(${this.formatEnum('ZeroTermsQuery', value.zero_terms_query)})`;
            }
            if (value.slop !== undefined) {
                code += `.Slop(${value.slop})`;
            }
            return code;
        }
        return `${this.getIndent()}es.MatchPhrasePrefix("${field}", ${this.formatValue(value)})`;
    }

    parseMultiMatch(multiMatch) {
        let code = `${this.getIndent()}es.MultiMatch(${this.formatValue(multiMatch.query)})`;

        if (multiMatch.fields) {
            const fieldsStr = multiMatch.fields.map(field => `"${field}"`).join(', ');
            code += `.Fields(${fieldsStr})`;
        }
        if (multiMatch.type) {
            code += `.Type(${this.formatEnum('TextQueryType', multiMatch.type)})`;
        }
        if (multiMatch.operator) {
            code += `.Operator(${this.formatEnum('Operator', multiMatch.operator)})`;
        }
        if (multiMatch.minimum_should_match !== undefined) {
            code += `.MinimumShouldMatch(${this.formatValue(multiMatch.minimum_should_match)})`;
        }
        if (multiMatch.tie_breaker !== undefined) {
            code += `.TieBreaker(${multiMatch.tie_breaker})`;
        }
        if (multiMatch.analyzer) {
            code += `.Analyzer("${multiMatch.analyzer}")`;
        }
        if (multiMatch.boost !== undefined) {
            code += `.Boost(${multiMatch.boost})`;
        }
        if (multiMatch.cutoff_frequency !== undefined) {
            code += `.CutoffFrequency(${multiMatch.cutoff_frequency})`;
        }
        if (multiMatch.fuzziness !== undefined) {
            code += `.Fuzziness(${this.formatValue(multiMatch.fuzziness)})`;
        }
        if (multiMatch.fuzzy_rewrite) {
            code += `.FuzzyRewrite("${multiMatch.fuzzy_rewrite}")`;
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
        if (multiMatch.slop !== undefined) {
            code += `.Slop(${multiMatch.slop})`;
        }
        if (multiMatch.zero_terms_query) {
            code += `.ZeroTermsQuery(${this.formatEnum('ZeroTermsQuery', multiMatch.zero_terms_query)})`;
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
            code += `.GreaterThan(${this.formatValue(conditions.gt)})`;
        }
        if (conditions.gte !== undefined) {
            code += `.GreaterThanOrEqual(${this.formatValue(conditions.gte)})`;
        }
        if (conditions.lt !== undefined) {
            code += `.LessThan(${this.formatValue(conditions.lt)})`;
        }
        if (conditions.lte !== undefined) {
            code += `.LessThanOrEqual(${this.formatValue(conditions.lte)})`;
        }
        if (conditions.from !== undefined) {
            code += `.From(${this.formatValue(conditions.from)})`;
        }
        if (conditions.to !== undefined) {
            code += `.To(${this.formatValue(conditions.to)})`;
        }
        if (conditions.boost !== undefined) {
            code += `.Boost(${conditions.boost})`;
        }
        if (conditions.relation) {
            code += `.Relation(${this.formatEnum('RangeRelation', conditions.relation)})`;
        }
        if (conditions.format) {
            code += `.Format("${conditions.format}")`;
        }

        return code;
    }

    parseExists(exists) {
        let code = `${this.getIndent()}es.Exists("${exists.field}")`;
        if (exists.boost !== undefined) {
            code += `.Boost(${exists.boost})`;
        }
        return code;
    }

    parseNested(nested) {
        let code = `${this.getIndent()}es.Nested("${nested.path}", ${this.parseQueryBody(nested.query)})`;

        if (nested.score_mode) {
            code += `.ScoreMode(${this.formatEnum('ScoreMode', nested.score_mode)})`;
        }
        if (nested.inner_hits) {
            code += `.InnerHits(${this.parseInnerHits(nested.inner_hits)})`;
        }
        if (nested.boost !== undefined) {
            code += `.Boost(${nested.boost})`;
        }
        if (nested.ignore_unmapped !== undefined) {
            code += `.IgnoreUnmapped(${nested.ignore_unmapped})`;
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
            const sortParts = innerHits.sort.map(sortItem => this.parseSortItem(sortItem));
            code += `.Sort(${sortParts.join(', ')})`;
        }
        if (innerHits.explain !== undefined) {
            code += `.Explain(${innerHits.explain})`;
        }
        if (innerHits.version !== undefined) {
            code += `.Version(${innerHits.version})`;
        }
        if (innerHits.seq_no_primary_term !== undefined) {
            code += `.SeqNoPrimaryTerm(${innerHits.seq_no_primary_term})`;
        }
        if (innerHits.track_scores !== undefined) {
            code += `.TrackScores(${innerHits.track_scores})`;
        }
        if (innerHits.ignore_unmapped !== undefined) {
            code += `.IgnoreUnmapped(${innerHits.ignore_unmapped})`;
        }
        if (innerHits.fields) {
            code += `.Fields(${innerHits.fields.map(f => `"${f}"`).join(', ')})`;
        }
        if (innerHits.stored_fields) {
            code += `.StoredFields(${innerHits.stored_fields.map(f => `"${f}"`).join(', ')})`;
        }
        if (innerHits._source !== undefined) {
            if (innerHits._source === false) {
                code += `.SourceFalse()`;
            } else if (Array.isArray(innerHits._source)) {
                code += `.SourceIncludes(${innerHits._source.map(f => `"${f}"`).join(', ')})`;
            } else if (typeof innerHits._source === 'object') {
                if (innerHits._source.includes) {
                    code += `.SourceIncludes(${innerHits._source.includes.map(f => `"${f}"`).join(', ')})`;
                }
                if (innerHits._source.excludes) {
                    code += `.SourceExcludes(${innerHits._source.excludes.map(f => `"${f}"`).join(', ')})`;
                }
            }
        }

        return code;
    }

    parseQueryString(qs) {
        let code = `${this.getIndent()}es.QueryString(${this.formatValue(qs.query)})`;

        if (qs.default_field) {
            code += `.DefaultField("${qs.default_field}")`;
        }
        if (qs.fields) {
            const fieldsStr = qs.fields.map(field => `"${field}"`).join(', ');
            code += `.Fields([]string{${fieldsStr}})`;
        }
        if (qs.type) {
            code += `.Type(${this.formatEnum('TextQueryType', qs.type)})`;
        }
        if (qs.tie_breaker !== undefined) {
            code += `.TieBreaker(${qs.tie_breaker})`;
        }
        if (qs.default_operator) {
            code += `.DefaultOperator(${this.formatEnum('Operator', qs.default_operator)})`;
        }
        if (qs.analyzer) {
            code += `.Analyzer("${qs.analyzer}")`;
        }
        if (qs.quote_analyzer) {
            code += `.QuoteAnalyzer("${qs.quote_analyzer}")`;
        }
        if (qs.allow_leading_wildcard !== undefined) {
            code += `.AllowLeadingWildcard(${qs.allow_leading_wildcard})`;
        }
        if (qs.enable_position_increments !== undefined) {
            code += `.EnablePositionIncrements(${qs.enable_position_increments})`;
        }
        if (qs.fuzzy_max_expansions !== undefined) {
            code += `.FuzzyMaxExpansions(${qs.fuzzy_max_expansions})`;
        }
        if (qs.fuzzy_prefix_length !== undefined) {
            code += `.FuzzyPrefixLength(${qs.fuzzy_prefix_length})`;
        }
        if (qs.fuzziness !== undefined) {
            code += `.Fuzziness(${this.formatStringParam(qs.fuzziness)})`;
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
        if (qs.minimum_should_match !== undefined) {
            code += `.MinimumShouldMatch(${this.formatStringParam(qs.minimum_should_match)})`;
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
        if (qs.time_zone) {
            code += `.TimeZone("${qs.time_zone}")`;
        }
        if (qs.escape !== undefined) {
            code += `.Escape(${qs.escape})`;
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
        let code = `${this.getIndent()}es.SimpleQueryString(${this.formatValue(sqs.query)})`;

        if (sqs.fields) {
            const fieldsStr = sqs.fields.map(field => `"${field}"`).join(', ');
            code += `.Fields([]string{${fieldsStr}})`;
        }
        if (sqs.default_operator) {
            code += `.DefaultOperator(${this.formatEnum('Operator', sqs.default_operator)})`;
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
        if (sqs.minimum_should_match !== undefined) {
            code += `.MinimumShouldMatch(${this.formatStringParam(sqs.minimum_should_match)})`;
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
        if (sqs.boost !== undefined) {
            code += `.Boost(${sqs.boost})`;
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
        let code = `${this.getIndent()}es.MatchNone()`;
        if (matchNone && matchNone.boost !== undefined) {
            code += `.Boost(${matchNone.boost})`;
        }
        return code;
    }

    parseConstantScore(constantScore) {
        let code = `${this.getIndent()}es.ConstantScore(${this.parseQueryBody(constantScore.filter)})`;
        if (constantScore.boost !== undefined) {
            code += `.Boost(${constantScore.boost})`;
        }
        if (constantScore._name) {
            code += `.Name("${constantScore._name}")`;
        }
        return code;
    }

    parseFunctionScore(fs) {
        let code = `${this.getIndent()}es.FunctionScore(`;
        if (fs.query) {
            code += `${this.parseQueryBody(fs.query)}`;
        } else {
            code += `es.MatchAll()`;
        }
        code += `)`;

        if (fs.boost !== undefined) {
            code += `.Boost(${fs.boost})`;
        }
        if (fs.max_boost !== undefined) {
            code += `.MaxBoost(${fs.max_boost})`;
        }
        if (fs.score_mode) {
            code += `.ScoreMode(${this.formatEnum('ScoreMode', fs.score_mode)})`;
        }
        if (fs.boost_mode) {
            code += `.BoostMode(${this.formatEnum('BoostMode', fs.boost_mode)})`;
        }
        if (fs.min_score !== undefined) {
            code += `.MinScore(${fs.min_score})`;
        }
        if (fs.weight !== undefined) {
            code += `.Weight(${fs.weight})`;
        }
        if (fs.random_score) {
            code += `.RandomScore(${fs.random_score.seed || 0}, "${fs.random_score.field || '_seq_no'}")`;
        }
        if (fs.script_score) {
            code += `.ScriptScore(${this.parseScript(fs.script_score.script)})`;
        }
        if (fs.field_value_factor) {
            code += `.FieldValueFactor(${this.parseFieldValueFactor(fs.field_value_factor)})`;
        }
        if (fs.functions) {
            code += `.Functions(\n`;
            this.indent++;
            const funcParts = fs.functions.map(fn => {
                return `${this.getIndent()}${this.parseFunctionScoreFunction(fn)}`;
            });
            code += funcParts.join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }

        return code;
    }

    parseFunctionScoreFunction(fn) {
        let code = '';
        if (fn.script_score) {
            code = `es.ScriptScoreFunction(${this.parseScript(fn.script_score.script)})`;
        } else if (fn.random_score !== undefined) {
            code = `es.RandomScoreFunction()`;
            if (fn.random_score.seed !== undefined) {
                code += `.Seed(${fn.random_score.seed})`;
            }
            if (fn.random_score.field) {
                code += `.Field("${fn.random_score.field}")`;
            }
        } else if (fn.field_value_factor) {
            code = `es.FieldValueFactorFunction(${this.parseFieldValueFactor(fn.field_value_factor)})`;
        } else if (fn.gauss || fn.linear || fn.exp) {
            const decayType = fn.gauss ? 'gauss' : fn.linear ? 'linear' : 'exp';
            const decay = fn[decayType];
            const field = Object.keys(decay)[0];
            const params = decay[field];
            let decayCode = `es.Decay("${field}")`;
            if (params.origin !== undefined) {
                decayCode += `.Origin(${this.formatValue(params.origin)})`;
            }
            if (params.scale !== undefined) {
                decayCode += `.Scale(${this.formatValue(params.scale)})`;
            }
            if (params.offset !== undefined) {
                decayCode += `.Offset(${this.formatValue(params.offset)})`;
            }
            if (params.decay !== undefined) {
                decayCode += `.DecayValue(${params.decay})`;
            }
            if (params.multi_value_mode) {
                decayCode += `.MultiValueMode(${this.formatEnum('MultiValuesMode', params.multi_value_mode)})`;
            }
            code = `es.DecayFunction("${decayType}", ${decayCode})`;
        } else if (fn.weight !== undefined) {
            code = `es.WeightFunction(${fn.weight})`;
        }

        if (fn.filter) {
            code += `.Filter(${this.parseQueryBody(fn.filter)})`;
        }
        if (fn.weight !== undefined && code.indexOf('WeightFunction') === -1) {
            code += `.Weight(${fn.weight})`;
        }

        return code;
    }

    parseFieldValueFactor(fvf) {
        let code = `es.FieldValueFactor("${fvf.field}")`;
        if (fvf.factor !== undefined) {
            code += `.Factor(${fvf.factor})`;
        }
        if (fvf.modifier) {
            code += `.Modifier(${this.formatEnum('Modifier', fvf.modifier)})`;
        }
        if (fvf.missing !== undefined) {
            code += `.Missing(${fvf.missing})`;
        }
        return code;
    }

    parseIds(ids) {
        const values = ids.values || [];
        const formattedValues = values.map(v => this.formatValue(v)).join(', ');
        let code = `${this.getIndent()}es.IDs(${formattedValues})`;
        if (ids.boost !== undefined) {
            code += `.Boost(${ids.boost})`;
        }
        if (ids._name) {
            code += `.Name("${ids._name}")`;
        }
        return code;
    }

    parseRegexp(regexp) {
        const field = Object.keys(regexp)[0];
        const value = regexp[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.Regexp("${field}", "${value.value}")`;
            if (value.flags) {
                code += `.Flags("${value.flags}")`;
            }
            if (value.case_insensitive !== undefined) {
                code += `.CaseInsensitive(${value.case_insensitive})`;
            }
            if (value.max_determinized_states !== undefined) {
                code += `.MaxDeterminizedStates(${value.max_determinized_states})`;
            }
            if (value.rewrite) {
                code += `.Rewrite("${value.rewrite}")`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Regexp("${field}", "${value}")`;
    }

    parseScriptQuery(script) {
        let code = `${this.getIndent()}es.ScriptQuery(${this.parseScript(script.script)})`;
        if (script.boost !== undefined) {
            code += `.Boost(${script.boost})`;
        }
        if (script._name) {
            code += `.Name("${script._name}")`;
        }
        return code;
    }

    parseTermsSet(termsSet) {
        const field = Object.keys(termsSet)[0];
        const value = termsSet[field];
        const terms = value.terms || [];
        const formattedTerms = terms.map(t => this.formatValue(t)).join(', ');
        let code = `${this.getIndent()}es.TermsSet("${field}", ${formattedTerms})`;
        if (value.minimum_should_match_field) {
            code += `.MinimumShouldMatchField("${value.minimum_should_match_field}")`;
        }
        if (value.minimum_should_match_script) {
            code += `.MinimumShouldMatchScript(${this.parseScript(value.minimum_should_match_script)})`;
        }
        if (value.boost !== undefined) {
            code += `.Boost(${value.boost})`;
        }
        return code;
    }

    parseFuzzy(fuzzy) {
        const field = Object.keys(fuzzy)[0];
        const value = fuzzy[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.Fuzzy("${field}", ${this.formatValue(value.value)})`;
            if (value.fuzziness !== undefined) {
                code += `.Fuzziness(${this.formatValue(value.fuzziness)})`;
            }
            if (value.max_expansions !== undefined) {
                code += `.MaxExpansions(${value.max_expansions})`;
            }
            if (value.prefix_length !== undefined) {
                code += `.PrefixLength(${value.prefix_length})`;
            }
            if (value.transpositions !== undefined) {
                code += `.Transpositions(${value.transpositions})`;
            }
            if (value.rewrite) {
                code += `.Rewrite("${value.rewrite}")`;
            }
            if (value.case_insensitive !== undefined) {
                code += `.CaseInsensitive(${value.case_insensitive})`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Fuzzy("${field}", ${this.formatValue(value)})`;
    }

    parsePrefix(prefix) {
        const field = Object.keys(prefix)[0];
        const value = prefix[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.Prefix("${field}", ${this.formatValue(value.value)})`;
            if (value.case_insensitive !== undefined) {
                code += `.CaseInsensitive(${value.case_insensitive})`;
            }
            if (value.rewrite) {
                code += `.Rewrite("${value.rewrite}")`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Prefix("${field}", ${this.formatValue(value)})`;
    }

    parseWildcard(wildcard) {
        const field = Object.keys(wildcard)[0];
        const value = wildcard[field];
        if (typeof value === 'object' && value !== null) {
            let code = `${this.getIndent()}es.Wildcard("${field}", ${this.formatValue(value.value)})`;
            if (value.case_insensitive !== undefined) {
                code += `.CaseInsensitive(${value.case_insensitive})`;
            }
            if (value.rewrite) {
                code += `.Rewrite("${value.rewrite}")`;
            }
            if (value.boost !== undefined) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Wildcard("${field}", ${this.formatValue(value)})`;
    }

    parseDisMax(disMax) {
        const queries = disMax.queries || [];
        this.indent++;
        const queryParts = queries.map(q => this.parseQueryBody(q));
        this.indent--;
        let code = `${this.getIndent()}es.DisMax(\n`;
        this.indent++;
        code += queryParts.join(',\n');
        code += ',\n';
        this.indent--;
        code += `${this.getIndent()})`;
        if (disMax.tie_breaker !== undefined) {
            code += `.TieBreaker(${disMax.tie_breaker})`;
        }
        if (disMax.boost !== undefined) {
            code += `.Boost(${disMax.boost})`;
        }
        if (disMax._name) {
            code += `.Name("${disMax._name}")`;
        }
        return code;
    }

    parseGeoDistance(geoDistance) {
        const reserved = new Set([
            'distance', 'distance_type', 'validation_method',
            'ignore_unmapped', 'boost', '_name',
        ]);
        const field = Object.keys(geoDistance).find(k => !reserved.has(k));
        if (!field) {
            throw new Error('geo_distance requires a field');
        }
        if (geoDistance.distance === undefined) {
            throw new Error('geo_distance requires distance');
        }
        const point = this.parseGeoPoint(geoDistance[field]);
        let code = `${this.getIndent()}es.GeoDistance("${field}", ${point.lat}, ${point.lon}, ${this.formatValue(geoDistance.distance)})`;
        if (geoDistance.distance_type) {
            code += `.DistanceType(${this.formatEnum('DistanceType', geoDistance.distance_type)})`;
        }
        if (geoDistance.validation_method) {
            code += `.ValidationMethod(${this.formatEnum('ValidationMethod', geoDistance.validation_method)})`;
        }
        if (geoDistance.ignore_unmapped !== undefined) {
            code += `.IgnoreUnmapped(${geoDistance.ignore_unmapped})`;
        }
        if (geoDistance.boost !== undefined) {
            code += `.Boost(${geoDistance.boost})`;
        }
        if (geoDistance._name) {
            code += `.Name("${geoDistance._name}")`;
        }
        return code;
    }

    parseGeoBoundingBox(geoBoundingBox) {
        const reserved = new Set([
            'validation_method', 'ignore_unmapped', 'boost', '_name',
            'type',
        ]);
        const field = Object.keys(geoBoundingBox).find(k => !reserved.has(k));
        if (!field) {
            throw new Error('geo_bounding_box requires a field');
        }
        const box = geoBoundingBox[field];
        if (!box || box.top_left === undefined || box.bottom_right === undefined) {
            throw new Error('geo_bounding_box requires top_left and bottom_right');
        }
        const topLeft = this.parseGeoPoint(box.top_left);
        const bottomRight = this.parseGeoPoint(box.bottom_right);
        let code = `${this.getIndent()}es.GeoBoundingBox("${field}", ${topLeft.lat}, ${topLeft.lon}, ${bottomRight.lat}, ${bottomRight.lon})`;
        if (geoBoundingBox.validation_method) {
            code += `.ValidationMethod(${this.formatEnum('ValidationMethod', geoBoundingBox.validation_method)})`;
        }
        if (geoBoundingBox.ignore_unmapped !== undefined) {
            code += `.IgnoreUnmapped(${geoBoundingBox.ignore_unmapped})`;
        }
        if (geoBoundingBox.boost !== undefined) {
            code += `.Boost(${geoBoundingBox.boost})`;
        }
        if (geoBoundingBox._name) {
            code += `.Name("${geoBoundingBox._name}")`;
        }
        return code;
    }

    parseHighlight(highlight) {
        let code = `es.Highlight()`;
        if (highlight.pre_tags) {
            const tags = highlight.pre_tags.map(t => `"${t}"`).join(', ');
            code += `.PreTags(${tags})`;
        }
        if (highlight.post_tags) {
            const tags = highlight.post_tags.map(t => `"${t}"`).join(', ');
            code += `.PostTags(${tags})`;
        }
        if (highlight.type) {
            code += `.Type(${this.formatEnum('HighlighterType', highlight.type)})`;
        }
        if (highlight.order) {
            code += `.Order("${highlight.order}")`;
        }
        if (highlight.encoder) {
            code += `.Encoder("${highlight.encoder}")`;
        }
        if (highlight.require_field_match !== undefined) {
            code += `.RequireFieldMatch(${highlight.require_field_match})`;
        }
        if (highlight.fragment_size !== undefined) {
            code += `.FragmentSize(${highlight.fragment_size})`;
        }
        if (highlight.number_of_fragments !== undefined) {
            code += `.NumberOfFragments(${highlight.number_of_fragments})`;
        }
        if (highlight.no_match_size !== undefined) {
            code += `.NoMatchSize(${highlight.no_match_size})`;
        }
        if (highlight.boundary_scanner) {
            code += `.BoundaryScanner(${this.formatEnum('BoundaryScanner', highlight.boundary_scanner)})`;
        }
        if (highlight.boundary_chars) {
            code += `.BoundaryChars("${highlight.boundary_chars}")`;
        }
        if (highlight.boundary_max_scan !== undefined) {
            code += `.BoundaryMaxScan(${highlight.boundary_max_scan})`;
        }
        if (highlight.boundary_scanner_locale) {
            code += `.BoundaryScannerLocale("${highlight.boundary_scanner_locale}")`;
        }
        if (highlight.fragmenter) {
            code += `.Fragmenter(${this.formatEnum('Fragmenter', highlight.fragmenter)})`;
        }
        if (highlight.fragment_offset !== undefined) {
            code += `.FragmentOffset(${highlight.fragment_offset})`;
        }
        if (highlight.max_fragment_length !== undefined) {
            code += `.MaxFragmentLength(${highlight.max_fragment_length})`;
        }
        if (highlight.max_analyzed_offset !== undefined) {
            code += `.MaxAnalyzedOffset(${highlight.max_analyzed_offset})`;
        }
        if (highlight.highlight_query) {
            code += `.HighlightQuery(${this.parseQueryBody(highlight.highlight_query)})`;
        }
        if (highlight.tags_schema) {
            code += `.TagsSchema("${highlight.tags_schema}")`;
        }
        if (highlight.fields) {
            for (const [fieldName, fieldOpts] of Object.entries(highlight.fields)) {
                let fieldCode = `es.HighlightField("${fieldName}")`;
                if (fieldOpts && typeof fieldOpts === 'object') {
                    if (fieldOpts.fragment_size !== undefined) {
                        fieldCode += `.FragmentSize(${fieldOpts.fragment_size})`;
                    }
                    if (fieldOpts.number_of_fragments !== undefined) {
                        fieldCode += `.NumberOfFragments(${fieldOpts.number_of_fragments})`;
                    }
                    if (fieldOpts.pre_tags) {
                        const tags = fieldOpts.pre_tags.map(t => `"${t}"`).join(', ');
                        fieldCode += `.PreTags(${tags})`;
                    }
                    if (fieldOpts.post_tags) {
                        const tags = fieldOpts.post_tags.map(t => `"${t}"`).join(', ');
                        fieldCode += `.PostTags(${tags})`;
                    }
                    if (fieldOpts.type) {
                        fieldCode += `.Type(${this.formatEnum('HighlighterType', fieldOpts.type)})`;
                    }
                    if (fieldOpts.no_match_size !== undefined) {
                        fieldCode += `.NoMatchSize(${fieldOpts.no_match_size})`;
                    }
                    if (fieldOpts.require_field_match !== undefined) {
                        fieldCode += `.RequireFieldMatch(${fieldOpts.require_field_match})`;
                    }
                    if (fieldOpts.order) {
                        fieldCode += `.Order("${fieldOpts.order}")`;
                    }
                    if (fieldOpts.fragmenter) {
                        fieldCode += `.Fragmenter(${this.formatEnum('Fragmenter', fieldOpts.fragmenter)})`;
                    }
                    if (fieldOpts.highlight_query) {
                        fieldCode += `.HighlightQuery(${this.parseQueryBody(fieldOpts.highlight_query)})`;
                    }
                    if (fieldOpts.matched_fields) {
                        const matched = fieldOpts.matched_fields.map(f => `"${f}"`).join(', ');
                        fieldCode += `.MatchedFields(${matched})`;
                    }
                }
                code += `.\n${this.getIndent()}Field(${fieldCode})`;
            }
        }
        return code;
    }

    parseAggregation(agg) {
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
                    code += `.\n${this.getIndent()}Include(${agg.terms.include.map(v => `"${v}"`).join(', ')})`;
                } else {
                    code += `.\n${this.getIndent()}Include("${agg.terms.include}")`;
                }
            }
            if (agg.terms.exclude) {
                if (Array.isArray(agg.terms.exclude)) {
                    code += `.\n${this.getIndent()}Exclude(${agg.terms.exclude.map(v => `"${v}"`).join(', ')})`;
                } else {
                    code += `.\n${this.getIndent()}Exclude("${agg.terms.exclude}")`;
                }
            }
            if (agg.terms.execution_hint) {
                code += `.\n${this.getIndent()}ExecutionHint(${this.formatEnum('ExecutionHint', agg.terms.execution_hint)})`;
            }
            if (agg.terms.collect_mode) {
                code += `.\n${this.getIndent()}CollectMode(${this.formatEnum('CollectMode', agg.terms.collect_mode)})`;
            }
            code = this.appendMeta(code, agg.terms.meta);
            if (agg.terms.order) {
                for (const [field, order] of Object.entries(agg.terms.order)) {
                    code += `.\n${this.getIndent()}Order("${field}", ${this.formatEnum('Order', order)})`;
                }
            }
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.stats.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.extended_stats.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.min.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.max.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.avg.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.sum.meta);
            return this.appendSubAggs(code, agg);
        }

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
            code = this.appendMeta(code, agg.cardinality.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.value_count) {
            let code = `es.ValueCountAgg("${agg.value_count.field}")`;
            if (agg.value_count.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.value_count.missing)})`;
            }
            if (agg.value_count.script) {
                code += `.\n${this.getIndent()}Script(${this.parseScript(agg.value_count.script)})`;
            }
            code = this.appendMeta(code, agg.value_count.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.top_hits) {
            return this.parseTopHitsAgg(agg.top_hits, agg);
        }

        if (agg.multi_terms) {
            let code = `es.MultiTermsAgg(${agg.multi_terms.terms.map(term => `es.TermAgg("${term.field}")`).join(', ')})`;
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
                    code += `.\n${this.getIndent()}Include(${agg.multi_terms.include.map(v => `"${v}"`).join(', ')})`;
                } else {
                    code += `.\n${this.getIndent()}Include("${agg.multi_terms.include}")`;
                }
            }
            if (agg.multi_terms.exclude) {
                if (Array.isArray(agg.multi_terms.exclude)) {
                    code += `.\n${this.getIndent()}Exclude(${agg.multi_terms.exclude.map(v => `"${v}"`).join(', ')})`;
                } else {
                    code += `.\n${this.getIndent()}Exclude("${agg.multi_terms.exclude}")`;
                }
            }
            if (agg.multi_terms.min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}MinDocCount(${agg.multi_terms.min_doc_count})`;
            }
            if (agg.multi_terms.execution_hint) {
                code += `.\n${this.getIndent()}ExecutionHint(${this.formatEnum('ExecutionHint', agg.multi_terms.execution_hint)})`;
            }
            if (agg.multi_terms.collect_mode) {
                code += `.\n${this.getIndent()}CollectMode(${this.formatEnum('CollectMode', agg.multi_terms.collect_mode)})`;
            }
            if (agg.multi_terms.order) {
                for (const [field, order] of Object.entries(agg.multi_terms.order)) {
                    code += `.\n${this.getIndent()}Order("${field}", ${this.formatEnum('Order', order)})`;
                }
            }
            return this.appendSubAggs(code, agg);
        }

        if (agg.nested) {
            let code = `es.NestedAgg("${agg.nested.path}")`;
            return this.appendSubAggs(code, agg);
        }

        if (agg.reverse_nested !== undefined) {
            let code = `es.ReverseNestedAgg()`;
            if (agg.reverse_nested.path) {
                code += `.\n${this.getIndent()}Path("${agg.reverse_nested.path}")`;
            }
            return this.appendSubAggs(code, agg);
        }

        if (agg.date_histogram) {
            let code = `es.DateHistogramAgg("${agg.date_histogram.field}")`;
            if (agg.date_histogram.calendar_interval) {
                code += `.\n${this.getIndent()}CalendarInterval("${agg.date_histogram.calendar_interval}")`;
            }
            if (agg.date_histogram.fixed_interval) {
                code += `.\n${this.getIndent()}FixedInterval("${agg.date_histogram.fixed_interval}")`;
            }
            if (agg.date_histogram.format) {
                code += `.\n${this.getIndent()}Format("${agg.date_histogram.format}")`;
            }
            if (agg.date_histogram.time_zone) {
                code += `.\n${this.getIndent()}TimeZone("${agg.date_histogram.time_zone}")`;
            }
            if (agg.date_histogram.offset) {
                code += `.\n${this.getIndent()}Offset("${agg.date_histogram.offset}")`;
            }
            if (agg.date_histogram.min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}MinDocCount(${agg.date_histogram.min_doc_count})`;
            }
            if (agg.date_histogram.keyed !== undefined) {
                code += `.\n${this.getIndent()}Keyed(${agg.date_histogram.keyed})`;
            }
            if (agg.date_histogram.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.date_histogram.missing)})`;
            }
            if (agg.date_histogram.extended_bounds) {
                code += `.\n${this.getIndent()}ExtendedBounds(${this.formatValue(agg.date_histogram.extended_bounds.min)}, ${this.formatValue(agg.date_histogram.extended_bounds.max)})`;
            }
            if (agg.date_histogram.hard_bounds) {
                code += `.\n${this.getIndent()}HardBounds(${this.formatValue(agg.date_histogram.hard_bounds.min)}, ${this.formatValue(agg.date_histogram.hard_bounds.max)})`;
            }
            if (agg.date_histogram.order && typeof agg.date_histogram.order === 'object') {
                for (const [field, order] of Object.entries(agg.date_histogram.order)) {
                    code += `.\n${this.getIndent()}Order(es.AggOrder("${field}", ${this.formatEnum('Order', order)}))`;
                }
            }
            code = this.appendMeta(code, agg.date_histogram.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.date_range) {
            let code = `es.DateRangeAgg("${agg.date_range.field}")`;
            if (agg.date_range.format) {
                code += `.\n${this.getIndent()}Format("${agg.date_range.format}")`;
            }
            if (agg.date_range.time_zone) {
                code += `.\n${this.getIndent()}TimeZone("${agg.date_range.time_zone}")`;
            }
            if (agg.date_range.keyed !== undefined) {
                code += `.\n${this.getIndent()}Keyed(${agg.date_range.keyed})`;
            }
            if (agg.date_range.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.date_range.missing)})`;
            }
            if (agg.date_range.ranges) {
                for (const range of agg.date_range.ranges) {
                    let entryCode = `es.DateRangeEntry()`;
                    if (range.key) {
                        entryCode += `.Key("${range.key}")`;
                    }
                    if (range.from !== undefined) {
                        entryCode += `.From(${this.formatValue(range.from)})`;
                    }
                    if (range.to !== undefined) {
                        entryCode += `.To(${this.formatValue(range.to)})`;
                    }
                    code += `.\n${this.getIndent()}Range(${entryCode})`;
                }
            }
            code = this.appendMeta(code, agg.date_range.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.filter && typeof agg.filter === 'object') {
            let code = `es.FilterAgg(${this.parseQueryBody(agg.filter)})`;
            code = this.appendMeta(code, agg.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.filters) {
            let code = `es.FiltersAgg()`;
            if (agg.filters.filters) {
                for (const [name, filter] of Object.entries(agg.filters.filters)) {
                    code += `.\n${this.getIndent()}Filter("${name}", ${this.parseQueryBody(filter)})`;
                }
            }
            if (agg.filters.other_bucket !== undefined) {
                code += `.\n${this.getIndent()}OtherBucket(${agg.filters.other_bucket})`;
            }
            if (agg.filters.other_bucket_key) {
                code += `.\n${this.getIndent()}OtherBucketKey("${agg.filters.other_bucket_key}")`;
            }
            code = this.appendMeta(code, agg.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.histogram) {
            let code = `es.HistogramAgg("${agg.histogram.field}", ${agg.histogram.interval})`;
            if (agg.histogram.min_doc_count !== undefined) {
                code += `.\n${this.getIndent()}MinDocCount(${agg.histogram.min_doc_count})`;
            }
            if (agg.histogram.offset !== undefined) {
                code += `.\n${this.getIndent()}Offset(${agg.histogram.offset})`;
            }
            if (agg.histogram.keyed !== undefined) {
                code += `.\n${this.getIndent()}Keyed(${agg.histogram.keyed})`;
            }
            if (agg.histogram.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.histogram.missing)})`;
            }
            if (agg.histogram.extended_bounds) {
                code += `.\n${this.getIndent()}ExtendedBounds(${agg.histogram.extended_bounds.min}, ${agg.histogram.extended_bounds.max})`;
            }
            if (agg.histogram.hard_bounds) {
                code += `.\n${this.getIndent()}HardBounds(${agg.histogram.hard_bounds.min}, ${agg.histogram.hard_bounds.max})`;
            }
            if (agg.histogram.order && typeof agg.histogram.order === 'object') {
                for (const [field, order] of Object.entries(agg.histogram.order)) {
                    code += `.\n${this.getIndent()}Order(es.AggOrder("${field}", ${this.formatEnum('Order', order)}))`;
                }
            }
            code = this.appendMeta(code, agg.histogram.meta);
            return this.appendSubAggs(code, agg);
        }

        if (agg.range) {
            let code = `es.RangeAgg("${agg.range.field}")`;
            if (agg.range.keyed !== undefined) {
                code += `.\n${this.getIndent()}Keyed(${agg.range.keyed})`;
            }
            if (agg.range.missing !== undefined) {
                code += `.\n${this.getIndent()}Missing(${this.formatValue(agg.range.missing)})`;
            }
            if (agg.range.ranges) {
                for (const range of agg.range.ranges) {
                    let entryCode = `es.RangeEntry()`;
                    if (range.key) {
                        entryCode += `.Key("${range.key}")`;
                    }
                    if (range.from !== undefined) {
                        entryCode += `.From(${range.from})`;
                    }
                    if (range.to !== undefined) {
                        entryCode += `.To(${range.to})`;
                    }
                    code += `.\n${this.getIndent()}Range(${entryCode})`;
                }
            }
            code = this.appendMeta(code, agg.range.meta);
            return this.appendSubAggs(code, agg);
        }

        throw new Error(`Unsupported aggregation type: ${Object.keys(agg)}`);
    }

    parseTopHitsAgg(topHits, agg) {
        let code = `es.TopHitsAgg()`;
        if (topHits.size !== undefined) {
            code += `.\n${this.getIndent()}Size(${topHits.size})`;
        }
        if (topHits.from !== undefined) {
            code += `.\n${this.getIndent()}From(${topHits.from})`;
        }
        if (topHits.sort) {
            const sortParts = topHits.sort.map(sortItem => this.parseSortItem(sortItem));
            code += `.\n${this.getIndent()}Sort(${sortParts.join(', ')})`;
        }
        if (topHits._source !== undefined) {
            if (topHits._source === false) {
                code += `.\n${this.getIndent()}SourceFalse()`;
            } else if (Array.isArray(topHits._source)) {
                code += `.\n${this.getIndent()}SourceIncludes(${topHits._source.map(f => `"${f}"`).join(', ')})`;
            } else if (typeof topHits._source === 'object') {
                if (topHits._source.includes) {
                    code += `.\n${this.getIndent()}SourceIncludes(${topHits._source.includes.map(f => `"${f}"`).join(', ')})`;
                }
                if (topHits._source.excludes) {
                    code += `.\n${this.getIndent()}SourceExcludes(${topHits._source.excludes.map(f => `"${f}"`).join(', ')})`;
                }
            }
        }
        if (topHits.highlight) {
            code += `.\n${this.getIndent()}Highlight(${this.parseHighlight(topHits.highlight)})`;
        }
        if (topHits.explain !== undefined) {
            code += `.\n${this.getIndent()}Explain(${topHits.explain})`;
        }
        if (topHits.version !== undefined) {
            code += `.\n${this.getIndent()}Version(${topHits.version})`;
        }
        if (topHits.seq_no_primary_term !== undefined) {
            code += `.\n${this.getIndent()}SeqNoPrimaryTerm(${topHits.seq_no_primary_term})`;
        }
        if (topHits.track_scores !== undefined) {
            code += `.\n${this.getIndent()}TrackScores(${topHits.track_scores})`;
        }
        code = this.appendMeta(code, topHits.meta || agg.meta);
        return code;
    }

    getIndent() {
        return '\t'.repeat(this.indent);
    }

    capitalize(str) {
        return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
    }

    parseScript(script) {
        if (typeof script === 'string') {
            return `es.ScriptSource(${this.formatValue(script)}, ScriptLanguage.Painless)`;
        }

        let scriptCode = '';
        const lang = script.lang
            ? this.formatEnum('ScriptLanguage', script.lang)
            : 'ScriptLanguage.Painless';

        if (script.source) {
            scriptCode = `es.ScriptSource(${this.formatValue(script.source)}, ${lang})`;
        } else if (script.id) {
            scriptCode = `es.ScriptID(${this.formatValue(script.id)}, ${lang})`;
        } else {
            throw new Error('Script requires source or id');
        }

        if (script.params) {
            for (const [key, value] of Object.entries(script.params)) {
                scriptCode += `.Parameter("${key}", ${this.formatValue(value)})`;
            }
        }

        if (script.options) {
            for (const [key, value] of Object.entries(script.options)) {
                scriptCode += `.Option("${key}", ${this.formatValue(String(value))})`;
            }
        }

        return scriptCode;
    }
}

if (typeof module !== 'undefined' && module.exports) {
    module.exports = ESQueryParser;
}
