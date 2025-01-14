class ESQueryParser {
    constructor() {
        this.indent = 0;
    }

    parse(esQuery) {
        try {
            let code = 'es.NewQuery(\n';
            this.indent = 1;

            // Ana query gövdesi
            if (esQuery.query) {
                code += this.parseQueryBody(esQuery.query);
            }

            // NewQuery parantezini kapat
            code = code.trimEnd();
            code += ')';

            // Top level parametreler
            code += this.parseTopLevelParams(esQuery);

            return code;
        } catch (error) {
            console.error('Parsing error:', error);
            throw new Error(`Query parsing failed: ${error.message}`);
        }
    }

    parseTopLevelParams(query) {
        let code = '';
        
        // Sort
        if (query.sort) {
            code += `.\n${this.getIndent()}Sort(\n`;
            this.indent++;
            const sortParts = query.sort.map(sortItem => {
                const field = Object.keys(sortItem)[0];
                const options = sortItem[field];
                let sortCode = `${this.getIndent()}es.Sort("${field}")`;
                if (options.order) {
                    sortCode += `.Order(Order.${this.capitalize(options.order)})`;
                }
                if (options.mode) {
                    sortCode += `.Mode(Mode.${this.capitalize(options.mode)})`;
                }
                return sortCode;
            });
            code += sortParts.join(',\n');
            code += ',\n';
            this.indent--;
            code += `${this.getIndent()})`;
        }
        
        // Aggregations
        if (query.aggs || query.aggregations) {
            const aggs = query.aggs || query.aggregations;
            for (const [name, agg] of Object.entries(aggs)) {
                code += `.\n${this.getIndent()}Aggs("${name}", ${this.parseAggregation(agg)})`;
            }
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
        if (query.range) return this.parseRange(query.range);
        if (query.exists) return this.parseExists(query.exists);
        if (query.nested) return this.parseNested(query.nested);
        if (query.query_string) return this.parseQueryString(query.query_string);
        if (query.simple_query_string) return this.parseSimpleQueryString(query.simple_query_string);
        
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

        return code;
    }

    parseTerm(term) {
        const field = Object.keys(term)[0];
        const value = term[field];
        if (typeof value === 'object') {
            let code = `${this.getIndent()}es.Term("${field}", "${value.value}")`;
            if (value.boost) {
                code += `.Boost(${value.boost})`;
            }
            return code;
        }
        return `${this.getIndent()}es.Term("${field}", "${value}")`;
    }

    parseTerms(terms) {
        const field = Object.keys(terms)[0];
        const values = terms[field];
        
        // Array tipini belirle
        let arrayType = this.determineArrayType(values);
        let arrayPrefix = arrayType ? `${arrayType}` : 'es.Array';
        
        // JSON.stringify yerine özel array formatlama kullan
        return `${this.getIndent()}es.Terms("${field}", ${arrayPrefix}${this.formatArray(values)})`;
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
                code += `.Operator(Operator.${this.capitalize(value.operator)})`;
            }
            if (value.boost) {
                code += `.Boost(${value.boost})`;
            }
            if (value.fuzziness) {
                code += `.Fuzziness("${value.fuzziness}")`;
            }
            return code;
        }
        return `${this.getIndent()}es.Match("${field}", "${value}")`;
    }

    parseRange(range) {
        const field = Object.keys(range)[0];
        const conditions = range[field];
        let code = `${this.getIndent()}es.Range("${field}")`;

        if (conditions.gte !== undefined) {
            code += `.GreaterThanOrEqual(${JSON.stringify(conditions.gte)})`;
        }
        if (conditions.gt !== undefined) {
            code += `.GreaterThan(${JSON.stringify(conditions.gt)})`;
        }
        if (conditions.lte !== undefined) {
            code += `.LesserThanOrEqual(${JSON.stringify(conditions.lte)})`;
        }
        if (conditions.lt !== undefined) {
            code += `.LesserThan(${JSON.stringify(conditions.lt)})`;
        }
        if (conditions.boost) {
            code += `.Boost(${conditions.boost})`;
        }

        return code;
    }

    parseExists(exists) {
        return `${this.getIndent()}es.Exists("${exists.field}")`;
    }

    parseNested(nested) {
        let code = `${this.getIndent()}es.Nested("${nested.path}", `;
        this.indent++;
        code += this.parseQueryBody(nested.query);
        this.indent--;
        code += ')';
        
        if (nested.score_mode) {
            code += `.ScoreMode(ScoreMode.${this.capitalize(nested.score_mode)})`;
        }
        if (nested.ignore_unmapped !== undefined) {
            code += `.IgnoreUnmapped(${nested.ignore_unmapped})`;
        }
        
        return code;
    }

    parseQueryString(qs) {
        let code = `${this.getIndent()}es.QueryString("${qs.query}")`;
        
        if (qs.fields) {
            code += `.Fields(${JSON.stringify(qs.fields)})`;
        }
        if (qs.default_operator) {
            code += `.DefaultOperator(Operator.${this.capitalize(qs.default_operator)})`;
        }
        if (qs.analyzer) {
            code += `.Analyzer("${qs.analyzer}")`;
        }
        
        return code;
    }

    parseSimpleQueryString(sqs) {
        let code = `${this.getIndent()}es.SimpleQueryString("${sqs.query}")`;
        
        if (sqs.fields) {
            code += `.Fields(${JSON.stringify(sqs.fields)})`;
        }
        if (sqs.default_operator) {
            code += `.DefaultOperator(Operator.${this.capitalize(sqs.default_operator)})`;
        }
        
        return code;
    }

    parseMatchAll(matchAll) {
        let code = `${this.getIndent()}es.MatchAll()`;
        if (matchAll && matchAll.boost) {
            code += `.Boost(${matchAll.boost})`;
        }
        return code;
    }

    parseMatchNone() {
        return `${this.getIndent()}es.MatchNone()`;
    }

    parseAggregation(agg) {
        const type = Object.keys(agg)[0];
        const params = agg[type];

        switch (type) {
            case 'terms':
                let termsCode = `es.AggTerms()`;
                if (params.field) {
                    termsCode += `.Field("${params.field}")`;
                }
                if (params.size) {
                    termsCode += `.Size(${params.size})`;
                }
                if (params.order) {
                    for (const [field, order] of Object.entries(params.order)) {
                        termsCode += `.Order("${field}", Order.${this.capitalize(order)})`;
                    }
                }
                return termsCode;

            case 'avg':
                return `es.AggAvg().Field("${params.field}")`;

            case 'sum':
                return `es.AggSum().Field("${params.field}")`;

            case 'min':
                return `es.AggMin().Field("${params.field}")`;

            case 'max':
                return `es.AggMax().Field("${params.field}")`;

            case 'nested':
                return `es.AggNested().Path("${params.path}")`;

            default:
                throw new Error(`Unsupported aggregation type: ${type}`);
        }
    }

    getIndent() {
        return '\t'.repeat(this.indent);
    }

    capitalize(str) {
        return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
    }
}