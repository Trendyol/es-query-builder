const assert = require('assert');
const ESQueryParser = require('./index.js');

const parser = new ESQueryParser();

function assertIncludes(code, snippet, label) {
    const normalized = code.replace(/\s+/g, '');
    const needle = snippet.replace(/\s+/g, '');
    assert.ok(normalized.includes(needle), `${label}\nExpected to include: ${snippet}\nGot:\n${code}`);
}

function assertNotIncludes(code, snippet, label) {
    const normalized = code.replace(/\s+/g, '');
    const needle = snippet.replace(/\s+/g, '');
    assert.ok(!normalized.includes(needle), `${label}\nExpected NOT to include: ${snippet}\nGot:\n${code}`);
}

function assertThrows(fn, messagePart, label) {
    let threw = false;
    try {
        fn();
    } catch (e) {
        threw = true;
        assert.ok(
            e.message.includes(messagePart),
            `${label}\nExpected error containing: ${messagePart}\nGot: ${e.message}`
        );
    }
    assert.ok(threw, `${label}: expected throw`);
}

// --- New query types ---

{
    const code = parser.parse({
        query: {
            fuzzy: {
                user: { value: 'ki', fuzziness: 'AUTO', max_expansions: 50, boost: 1.2 },
            },
        },
    });
    assertIncludes(code, 'es.Fuzzy("user", "ki")', 'fuzzy ctor');
    assertIncludes(code, '.Fuzziness("AUTO")', 'fuzzy fuzziness');
    assertIncludes(code, '.MaxExpansions(50)', 'fuzzy max_expansions');
    assertIncludes(code, '.Boost(1.2)', 'fuzzy boost');
}

{
    const code = parser.parse({
        query: { prefix: { 'user.id': { value: 'ki', case_insensitive: true } } },
    });
    assertIncludes(code, 'es.Prefix("user.id", "ki")', 'prefix ctor');
    assertIncludes(code, '.CaseInsensitive(true)', 'prefix case_insensitive');
}

{
    const code = parser.parse({
        query: { wildcard: { 'user.id': { value: 'ki*y', rewrite: 'constant_score' } } },
    });
    assertIncludes(code, 'es.Wildcard("user.id", "ki*y")', 'wildcard ctor');
    assertIncludes(code, '.Rewrite("constant_score")', 'wildcard rewrite');
}

{
    const code = parser.parse({
        query: {
            dis_max: {
                queries: [
                    { term: { title: 'quick' } },
                    { term: { body: 'brown' } },
                ],
                tie_breaker: 0.7,
                boost: 1.2,
            },
        },
    });
    assertIncludes(code, 'es.DisMax(', 'dis_max ctor');
    assertIncludes(code, 'es.Term("title", "quick")', 'dis_max query1');
    assertIncludes(code, '.TieBreaker(0.7)', 'dis_max tie_breaker');
}

{
    const code = parser.parse({
        query: {
            geo_distance: {
                distance: '12km',
                'pin.location': { lat: 40.0, lon: -70.0 },
                distance_type: 'arc',
                validation_method: 'STRICT',
            },
        },
    });
    assertIncludes(code, 'es.GeoDistance("pin.location", 40, -70, "12km")', 'geo_distance ctor');
    assertIncludes(code, '.DistanceType(DistanceType.Arc)', 'geo_distance distance_type');
    assertIncludes(code, '.ValidationMethod(ValidationMethod.Strict)', 'geo_distance validation');
}

{
    const code = parser.parse({
        query: {
            geo_bounding_box: {
                'pin.location': {
                    top_left: { lat: 40.73, lon: -74.1 },
                    bottom_right: { lat: 40.01, lon: -71.12 },
                },
            },
        },
    });
    assertIncludes(
        code,
        'es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)',
        'geo_bounding_box ctor'
    );
}

// --- New aggs ---

{
    const code = parser.parse({
        aggs: {
            prices: { value_count: { field: 'price', missing: 0 } },
        },
    });
    assertIncludes(code, 'es.ValueCountAgg("price")', 'value_count ctor');
    assertIncludes(code, '.Missing(0)', 'value_count missing');
}

{
    const code = parser.parse({
        aggs: {
            top: {
                top_hits: {
                    size: 3,
                    sort: [{ date: { order: 'desc' } }],
                    _source: { includes: ['title'] },
                    explain: true,
                },
            },
        },
    });
    assertIncludes(code, 'es.TopHitsAgg()', 'top_hits ctor');
    assertIncludes(code, '.Size(3)', 'top_hits size');
    assertIncludes(code, 'Order.Desc', 'top_hits sort order enum');
    assertIncludes(code, 'SourceIncludes("title")', 'top_hits source');
    assertIncludes(code, '.Explain(true)', 'top_hits explain');
}

// --- Bug fixes ---

{
    const code = parser.parse({
        query: { match_all: {} },
        post_filter: { term: { color: 'red' } },
        search_after: ['2024-01-01', 12345],
    });
    assertIncludes(code, 'PostFilter(', 'post_filter');
    assertIncludes(code, 'es.Term("color", "red")', 'post_filter term');
    assertIncludes(code, 'SearchAfter("2024-01-01", 12345)', 'search_after');
    assertNotIncludes(code, 'MinScore(', 'root min_score removed');
}

{
    const code = parser.parse({
        query: {
            match: {
                title: {
                    query: 'es-query-builder',
                    operator: 'and',
                    zero_terms_query: 'none',
                },
            },
        },
    });
    assertIncludes(code, '.Operator(Operator.And)', 'match operator enum');
    assertIncludes(code, '.ZeroTermsQuery(ZeroTermsQuery.None)', 'match zero_terms enum');
    assertNotIncludes(code, '.Operator("and")', 'no string operator');
}

{
    const code = parser.parse({
        query: {
            multi_match: {
                query: 'foo',
                fields: ['title', 'body'],
                type: 'best_fields',
            },
        },
    });
    assertIncludes(code, 'TextQueryType.Bestfields', 'multi_match type enum');
}

{
    const code = parser.parse({
        query: {
            range: {
                age: { gte: 10, lte: 20, from: 10, to: 20, relation: 'within' },
            },
        },
    });
    assertIncludes(code, '.From(10)', 'range from');
    assertIncludes(code, '.To(20)', 'range to');
    assertIncludes(code, 'RangeRelation.Within', 'range relation enum');
    assertNotIncludes(code, 'TimeZone(', 'range no TimeZone');
}

{
    const code = parser.parse({
        query: {
            script: {
                script: {
                    source: "doc['price'].value > params.t",
                    lang: 'painless',
                    params: { t: 100 },
                },
            },
        },
    });
    assertIncludes(code, 'es.ScriptSource(', 'script source');
    assertIncludes(code, 'ScriptLanguage.Painless', 'script lang');
    assertIncludes(code, '.Parameter("t", 100)', 'script parameter');
    assertNotIncludes(code, 'ScriptId(', 'no ScriptId typo');
    assertNotIncludes(code, '.Params(', 'no Params method');
}

{
    const code = parser.parse({
        query: {
            terms: {
                status: { value: ['open', 'closed'], boost: 2 },
            },
        },
    });
    assertIncludes(code, 'es.Terms("status", "open", "closed")', 'terms object form');
    assertIncludes(code, '.Boost(2)', 'terms boost');
}

{
    assertThrows(
        () => parser.parse({ query: { match_none: {} } }),
        'Empty match_none',
        'empty match_none throws'
    );
}

{
    const code = parser.parse({
        query: { match_all: {} },
        sort: [{ price: { order: 'asc', mode: 'avg' } }],
    });
    assertIncludes(code, '.Order(Order.Asc)', 'sort order enum');
    assertIncludes(code, '.Mode(Mode.Avg)', 'sort mode enum');
}

{
    const code = parser.parse({
        query: {
            bool: {
                must: [{ term: { a: 1 } }],
                adjust_pure_negative: false,
                boost: 1.5,
            },
        },
    });
    assertIncludes(code, '.AdjustPureNegative(false)', 'bool adjust_pure_negative');
}

{
    const code = parser.parse({
        query: {
            term: { name: { value: 'foo', case_insensitive: true, boost: 2 } },
        },
    });
    assertIncludes(code, '.CaseInsensitive(true)', 'term case_insensitive');
}

{
    const code = parser.parse({
        query: {
            nested: {
                path: 'offers',
                query: { term: { 'offers.color': 'red' } },
                boost: 2,
                ignore_unmapped: true,
            },
        },
    });
    assertIncludes(code, '.Boost(2)', 'nested boost');
    assertIncludes(code, '.IgnoreUnmapped(true)', 'nested ignore_unmapped');
}

{
    const code = parser.parse({
        aggs: {
            cats: {
                terms: {
                    field: 'category',
                    execution_hint: 'global_ordinals',
                    collect_mode: 'breadth_first',
                },
            },
        },
    });
    assertIncludes(code, 'ExecutionHint.GlobalOrdinals', 'execution_hint enum');
    assertIncludes(code, 'CollectMode.BreadthFirst', 'collect_mode enum');
}

console.log('All smoke tests passed.');
