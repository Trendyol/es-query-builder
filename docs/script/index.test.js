const assert = require('assert');
const ESQueryParser = require('./index.js');

const parser = new ESQueryParser();
let passed = 0;

function assertIncludes(code, snippet, label) {
    const normalized = code.replace(/\s+/g, '');
    const needle = snippet.replace(/\s+/g, '');
    assert.ok(normalized.includes(needle), `${label}\nExpected to include: ${snippet}\nGot:\n${code}`);
    passed++;
}

function assertNotIncludes(code, snippet, label) {
    const normalized = code.replace(/\s+/g, '');
    const needle = snippet.replace(/\s+/g, '');
    assert.ok(!normalized.includes(needle), `${label}\nExpected NOT to include: ${snippet}\nGot:\n${code}`);
    passed++;
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
    passed++;
}

function q(body) {
    return parser.parse({ query: body });
}

// =============================================================================
// Legacy queries
// =============================================================================

{
    const code = q({
        bool: {
            must: [{ term: { status: 'active' } }],
            should: [{ match: { title: 'hello' } }],
            filter: [{ exists: { field: 'user' } }],
            must_not: [{ term: { deleted: true } }],
            minimum_should_match: 1,
            adjust_pure_negative: false,
            boost: 1.5,
        },
    });
    assertIncludes(code, 'es.Bool()', 'bool ctor');
    assertIncludes(code, '.Must(', 'bool must');
    assertIncludes(code, '.Should(', 'bool should');
    assertIncludes(code, '.Filter(', 'bool filter');
    assertIncludes(code, '.MustNot(', 'bool must_not');
    assertIncludes(code, '.MinimumShouldMatch(1)', 'bool msm');
    assertIncludes(code, '.AdjustPureNegative(false)', 'bool adjust_pure_negative');
    assertIncludes(code, '.Boost(1.5)', 'bool boost');
}

{
    const code = q({ term: { color: { value: 'red', boost: 2, case_insensitive: true } } });
    assertIncludes(code, 'es.Term("color", "red")', 'term object');
    assertIncludes(code, '.Boost(2)', 'term boost');
    assertIncludes(code, '.CaseInsensitive(true)', 'term case_insensitive');
}

{
    const code = q({ term: { color: 'blue' } });
    assertIncludes(code, 'es.Term("color", "blue")', 'term shorthand');
}

{
    const code = q({ terms: { tags: ['a', 'b', 'c'] } });
    assertIncludes(code, 'es.Terms("tags", "a", "b", "c")', 'terms array');
}

{
    const code = q({ terms: { status: { value: ['open', 'closed'], boost: 2 } } });
    assertIncludes(code, 'es.Terms("status", "open", "closed")', 'terms object values');
    assertIncludes(code, '.Boost(2)', 'terms boost');
}

{
    const code = q({
        match: {
            title: {
                query: 'es builder',
                operator: 'and',
                boost: 1.2,
                cutoff_frequency: 0.001,
                fuzziness: 'AUTO',
                fuzzy_rewrite: 'constant_score',
                fuzzy_transpositions: true,
                lenient: true,
                max_expansions: 50,
                prefix_length: 1,
                zero_terms_query: 'all',
                auto_generate_synonyms_phrase_query: false,
            },
        },
    });
    assertIncludes(code, 'es.Match("title", "es builder")', 'match ctor');
    assertIncludes(code, '.Operator(Operator.And)', 'match operator');
    assertIncludes(code, '.CutoffFrequency(0.001)', 'match cutoff');
    assertIncludes(code, '.Fuzziness("AUTO")', 'match fuzziness');
    assertIncludes(code, '.FuzzyRewrite("constant_score")', 'match fuzzy_rewrite');
    assertIncludes(code, '.FuzzyTranspositions(true)', 'match fuzzy_transpositions');
    assertIncludes(code, '.Lenient(true)', 'match lenient');
    assertIncludes(code, '.MaxExpansions(50)', 'match max_expansions');
    assertIncludes(code, '.PrefixLength(1)', 'match prefix_length');
    assertIncludes(code, '.ZeroTermsQuery(ZeroTermsQuery.All)', 'match zero_terms');
    assertIncludes(code, '.AutoGenerateSynonymsPhraseQuery(false)', 'match auto_synonyms');
}

{
    const code = q({ match_all: { boost: 1.1 } });
    assertIncludes(code, 'es.MatchAll()', 'match_all');
    assertIncludes(code, '.Boost(1.1)', 'match_all boost');
}

{
    const code = q({ match_none: { title: { query: 'x' } } });
    assertIncludes(code, 'es.MatchNone("title", "x")', 'match_none with field');
}

{
    assertThrows(() => q({ match_none: {} }), 'Empty match_none', 'empty match_none');
}

{
    const code = q({
        match_bool_prefix: {
            title: {
                query: 'quick brown',
                analyzer: 'standard',
                minimum_should_match: '2',
                operator: 'or',
                boost: 1.5,
                fuzziness: 1,
                max_expansions: 10,
                prefix_length: 0,
            },
        },
    });
    assertIncludes(code, 'es.MatchBoolPrefix("title", "quick brown")', 'match_bool_prefix');
    assertIncludes(code, '.Analyzer("standard")', 'mbp analyzer');
    assertIncludes(code, '.Operator(Operator.Or)', 'mbp operator');
    assertIncludes(code, '.Fuzziness(1)', 'mbp fuzziness');
}

{
    const code = q({
        match_phrase: {
            title: { query: 'quick fox', analyzer: 'standard', boost: 2, zero_terms_query: 'none', slop: 2 },
        },
    });
    assertIncludes(code, 'es.MatchPhrase("title", "quick fox")', 'match_phrase');
    assertIncludes(code, '.Slop(2)', 'match_phrase slop');
    assertIncludes(code, '.ZeroTermsQuery(ZeroTermsQuery.None)', 'match_phrase ztq');
}

{
    const code = q({
        match_phrase_prefix: {
            title: { query: 'quick bro', analyzer: 'standard', max_expansions: 10, boost: 1, slop: 1, zero_terms_query: 'all' },
        },
    });
    assertIncludes(code, 'es.MatchPhrasePrefix("title", "quick bro")', 'match_phrase_prefix');
    assertIncludes(code, '.MaxExpansions(10)', 'mpp max_expansions');
    assertIncludes(code, '.Slop(1)', 'mpp slop');
}

{
    const code = q({
        multi_match: {
            query: 'foo bar',
            fields: ['title^2', 'body'],
            type: 'cross_fields',
            operator: 'and',
            minimum_should_match: '75%',
            tie_breaker: 0.3,
            analyzer: 'standard',
            boost: 1.5,
            cutoff_frequency: 0.01,
            fuzziness: 'AUTO',
            fuzzy_rewrite: 'top_terms_7',
            fuzzy_transpositions: false,
            lenient: true,
            prefix_length: 2,
            max_expansions: 20,
            slop: 1,
            zero_terms_query: 'none',
            auto_generate_synonyms_phrase_query: true,
        },
    });
    assertIncludes(code, 'es.MultiMatch("foo bar")', 'multi_match');
    assertIncludes(code, '.Fields("title^2", "body")', 'mm fields');
    assertIncludes(code, 'TextQueryType.Crossfields', 'mm type');
    assertIncludes(code, '.Operator(Operator.And)', 'mm operator');
    assertIncludes(code, '.TieBreaker(0.3)', 'mm tie_breaker');
    assertIncludes(code, '.Slop(1)', 'mm slop');
    assertIncludes(code, '.CutoffFrequency(0.01)', 'mm cutoff');
}

{
    const code = q({
        range: {
            age: { gt: 10, gte: 11, lt: 50, lte: 49, from: 11, to: 49, boost: 1, relation: 'intersects', format: 'yyyy' },
        },
    });
    assertIncludes(code, 'es.Range("age")', 'range');
    assertIncludes(code, '.GreaterThan(10)', 'range gt');
    assertIncludes(code, '.GreaterThanOrEqual(11)', 'range gte');
    assertIncludes(code, '.LessThan(50)', 'range lt');
    assertIncludes(code, '.LessThanOrEqual(49)', 'range lte');
    assertIncludes(code, '.From(11)', 'range from');
    assertIncludes(code, '.To(49)', 'range to');
    assertIncludes(code, 'RangeRelation.Intersects', 'range relation');
    assertIncludes(code, '.Format("yyyy")', 'range format');
    assertNotIncludes(code, 'TimeZone(', 'range no timezone');
}

{
    const code = q({ exists: { field: 'user', boost: 1.2 } });
    assertIncludes(code, 'es.Exists("user")', 'exists');
    assertIncludes(code, '.Boost(1.2)', 'exists boost');
}

{
    const code = q({
        nested: {
            path: 'offers',
            query: { term: { 'offers.color': 'red' } },
            score_mode: 'avg',
            boost: 2,
            ignore_unmapped: true,
            inner_hits: {
                name: 'ih',
                from: 0,
                size: 5,
                sort: [{ 'offers.price': { order: 'desc', mode: 'min' } }],
                explain: true,
                version: true,
                track_scores: true,
                ignore_unmapped: true,
                fields: ['offers.color'],
                _source: { includes: ['offers.color'], excludes: ['offers.desc'] },
            },
        },
    });
    assertIncludes(code, 'es.Nested("offers"', 'nested');
    assertIncludes(code, 'ScoreMode.Avg', 'nested score_mode');
    assertIncludes(code, '.Boost(2)', 'nested boost');
    assertIncludes(code, '.IgnoreUnmapped(true)', 'nested ignore_unmapped');
    assertIncludes(code, 'es.InnerHits()', 'inner_hits');
    assertIncludes(code, '.Name("ih")', 'inner_hits name');
    assertIncludes(code, '.Explain(true)', 'inner_hits explain');
    assertIncludes(code, '.Version(true)', 'inner_hits version');
    assertIncludes(code, '.TrackScores(true)', 'inner_hits track_scores');
    assertIncludes(code, '.Fields("offers.color")', 'inner_hits fields');
    assertIncludes(code, '.SourceIncludes("offers.color")', 'inner_hits source includes');
    assertIncludes(code, '.SourceExcludes("offers.desc")', 'inner_hits source excludes');
    assertIncludes(code, 'Order.Desc', 'inner_hits sort order');
    assertIncludes(code, 'Mode.Min', 'inner_hits sort mode');
}

{
    const code = q({
        query_string: {
            query: 'foo AND bar',
            default_field: 'title',
            fields: ['title', 'body'],
            type: 'best_fields',
            default_operator: 'and',
            analyzer: 'standard',
            quote_analyzer: 'keyword',
            allow_leading_wildcard: false,
            enable_position_increments: true,
            fuzzy_max_expansions: 40,
            fuzzy_prefix_length: 1,
            fuzziness: 'AUTO',
            fuzzy_transpositions: true,
            lenient: true,
            max_determinized_states: 10000,
            minimum_should_match: 2,
            quote_field_suffix: '.exact',
            phrase_slop: 2,
            boost: 1.5,
            analyze_wildcard: true,
            time_zone: '+03:00',
            escape: true,
            rewrite: 'constant_score',
            fuzzy_rewrite: 'constant_score',
            auto_generate_synonyms_phrase_query: false,
            tie_breaker: 0.2,
        },
    });
    assertIncludes(code, 'es.QueryString("foo AND bar")', 'query_string');
    assertIncludes(code, '.DefaultField("title")', 'qs default_field');
    assertIncludes(code, '.Fields([]string{"title", "body"})', 'qs fields');
    assertIncludes(code, 'TextQueryType.Bestfields', 'qs type');
    assertIncludes(code, 'Operator.And', 'qs default_operator');
    assertIncludes(code, '.Analyzer("standard")', 'qs analyzer');
    assertIncludes(code, '.QuoteAnalyzer("keyword")', 'qs quote_analyzer');
    assertIncludes(code, '.EnablePositionIncrements(true)', 'qs enable_position_increments');
    assertIncludes(code, '.FuzzyPrefixLength(1)', 'qs fuzzy_prefix_length');
    assertIncludes(code, '.MinimumShouldMatch("2")', 'qs msm string coerce');
    assertIncludes(code, '.TimeZone("+03:00")', 'qs time_zone');
    assertIncludes(code, '.Escape(true)', 'qs escape');
}

{
    const code = q({
        simple_query_string: {
            query: 'foo bar',
            fields: ['title'],
            default_operator: 'or',
            analyzer: 'standard',
            flags: 'ALL',
            fuzzy_max_expansions: 50,
            fuzzy_prefix_length: 0,
            fuzzy_transpositions: true,
            lenient: true,
            minimum_should_match: 1,
            quote_field_suffix: '.exact',
            analyze_wildcard: true,
            auto_generate_synonyms_phrase_query: true,
            boost: 2,
        },
    });
    assertIncludes(code, 'es.SimpleQueryString("foo bar")', 'simple_query_string');
    assertIncludes(code, 'Operator.Or', 'sqs default_operator');
    assertIncludes(code, '.MinimumShouldMatch("1")', 'sqs msm string coerce');
    assertIncludes(code, '.Boost(2)', 'sqs boost');
}

{
    const code = q({
        constant_score: {
            filter: { term: { status: 'ok' } },
            boost: 1.2,
            _name: 'cs',
        },
    });
    assertIncludes(code, 'es.ConstantScore(', 'constant_score');
    assertIncludes(code, '.Boost(1.2)', 'cs boost');
    assertIncludes(code, '.Name("cs")', 'cs name');
}

{
    const code = q({
        function_score: {
            query: { match_all: {} },
            boost: 1,
            max_boost: 10,
            score_mode: 'sum',
            boost_mode: 'multiply',
            min_score: 0.5,
            weight: 2,
            random_score: { seed: 10, field: '_seq_no' },
            script_score: { script: { source: '_score * 2', lang: 'painless' } },
            field_value_factor: { field: 'likes', factor: 1.2, modifier: 'log1p', missing: 1 },
            functions: [
                {
                    filter: { term: { tag: 'vip' } },
                    weight: 3,
                    gauss: {
                        location: { origin: '0,0', scale: '5km', offset: '1km', decay: 0.5, multi_value_mode: 'min' },
                    },
                },
                { script_score: { script: { source: '_score' } } },
                { random_score: { seed: 1, field: 'id' } },
                { field_value_factor: { field: 'price', factor: 2, modifier: 'sqrt' } },
                { weight: 5 },
            ],
        },
    });
    assertIncludes(code, 'es.FunctionScore(', 'function_score');
    assertIncludes(code, 'ScoreMode.Sum', 'fs score_mode');
    assertIncludes(code, 'BoostMode.Multiply', 'fs boost_mode');
    assertIncludes(code, '.MinScore(0.5)', 'fs min_score');
    assertIncludes(code, '.RandomScore(10, "_seq_no")', 'fs random_score');
    assertIncludes(code, '.ScriptScore(', 'fs script_score');
    assertIncludes(code, 'es.FieldValueFactor("likes")', 'fs fvf');
    assertIncludes(code, 'Modifier.Log1p', 'fs fvf modifier');
    assertIncludes(code, 'es.DecayFunction("gauss"', 'fs decay');
    assertIncludes(code, 'MultiValuesMode.Min', 'fs decay multi_value_mode');
    assertIncludes(code, 'es.WeightFunction(5)', 'fs weight fn');
    assertIncludes(code, 'es.ScriptScoreFunction(', 'fs script fn');
    assertIncludes(code, 'es.RandomScoreFunction()', 'fs random fn');
    assertIncludes(code, 'es.FieldValueFactorFunction(', 'fs fvf fn');
}

{
    const code = q({ ids: { values: ['1', '2'], boost: 1.5, _name: 'ids_q' } });
    assertIncludes(code, 'es.IDs("1", "2")', 'ids');
    assertIncludes(code, '.Boost(1.5)', 'ids boost');
    assertIncludes(code, '.Name("ids_q")', 'ids name');
}

{
    const code = q({
        regexp: {
            name: {
                value: 'j.*',
                flags: 'ALL',
                case_insensitive: true,
                max_determinized_states: 1000,
                rewrite: 'constant_score',
                boost: 1.1,
            },
        },
    });
    assertIncludes(code, 'es.Regexp("name", "j.*")', 'regexp');
    assertIncludes(code, '.Flags("ALL")', 'regexp flags');
    assertIncludes(code, '.CaseInsensitive(true)', 'regexp case_insensitive');
    assertIncludes(code, '.MaxDeterminizedStates(1000)', 'regexp max_states');
}

{
    const code = q({
        script: {
            script: {
                id: 'my_script',
                lang: 'painless',
                params: { threshold: 100 },
                options: { cache: 'true' },
            },
            boost: 1.2,
            _name: 'sq',
        },
    });
    assertIncludes(code, 'es.ScriptID("my_script", ScriptLanguage.Painless)', 'script_query id');
    assertIncludes(code, '.Parameter("threshold", 100)', 'script params');
    assertIncludes(code, '.Option("cache", "true")', 'script options');
    assertIncludes(code, '.Boost(1.2)', 'script_query boost');
    assertIncludes(code, '.Name("sq")', 'script_query name');
}

{
    const code = q({
        terms_set: {
            tags: {
                terms: ['a', 'b'],
                minimum_should_match_field: 'required_matches',
                boost: 1,
            },
        },
    });
    assertIncludes(code, 'es.TermsSet("tags", "a", "b")', 'terms_set');
    assertIncludes(code, '.MinimumShouldMatchField("required_matches")', 'terms_set msm field');
}

{
    const code = q({
        terms_set: {
            tags: {
                terms: ['x'],
                minimum_should_match_script: { source: '1', lang: 'painless' },
            },
        },
    });
    assertIncludes(code, '.MinimumShouldMatchScript(', 'terms_set msm script');
}

// =============================================================================
// Legacy top-level
// =============================================================================

{
    const code = parser.parse({
        query: { match_all: {} },
        size: 10,
        from: 5,
        track_total_hits: true,
        _source: { includes: ['a'], excludes: ['b'] },
        sort: [
            { date: { order: 'desc', mode: 'avg', nested: { path: 'offers', filter: { term: { 'offers.active': true } } } } },
            'score',
        ],
        highlight: {
            pre_tags: ['<em>'],
            post_tags: ['</em>'],
            type: 'unified',
            fragmenter: 'span',
            boundary_scanner: 'word',
            fields: {
                title: {
                    fragment_size: 150,
                    number_of_fragments: 3,
                    type: 'plain',
                    fragmenter: 'simple',
                    matched_fields: ['title', 'title.stemmed'],
                },
            },
        },
        post_filter: { term: { color: 'red' } },
        search_after: ['2024-01-01', 99],
        aggs: { by_color: { terms: { field: 'color', size: 5 } } },
    });
    assertIncludes(code, 'es.NewQuery(', 'top new query');
    assertIncludes(code, '.Size(10)', 'top size');
    assertIncludes(code, '.From(5)', 'top from');
    assertIncludes(code, '.TrackTotalHits(true)', 'top track_total_hits');
    assertIncludes(code, 'SourceIncludes("a")', 'top source includes');
    assertIncludes(code, 'SourceExcludes("b")', 'top source excludes');
    assertIncludes(code, 'Order.Desc', 'top sort order');
    assertIncludes(code, 'Mode.Avg', 'top sort mode');
    assertIncludes(code, 'es.NestedSort("offers")', 'top nested sort');
    assertIncludes(code, 'es.Sort("score")', 'top sort string');
    assertIncludes(code, 'es.Highlight()', 'top highlight');
    assertIncludes(code, 'HighlighterType.Unified', 'highlight type');
    assertIncludes(code, 'Fragmenter.Span', 'highlight fragmenter');
    assertIncludes(code, 'BoundaryScanner.Word', 'highlight boundary');
    assertIncludes(code, 'MatchedFields("title", "title.stemmed")', 'highlight matched_fields');
    assertIncludes(code, 'PostFilter(', 'top post_filter');
    assertIncludes(code, 'SearchAfter("2024-01-01", 99)', 'top search_after');
    assertIncludes(code, 'es.Agg("by_color"', 'top aggs');
    assertNotIncludes(code, 'MinScore(', 'no root min_score');
}

{
    const code = parser.parse({
        query: { match_all: {} },
        _source: false,
    });
    assertIncludes(code, 'SourceFalse()', 'source false');
}

{
    const code = parser.parse({
        query: { match_all: {} },
        _source: ['title', 'price'],
    });
    assertIncludes(code, 'SourceIncludes("title", "price")', 'source array');
}

// =============================================================================
// Legacy aggregations
// =============================================================================

{
    const code = parser.parse({
        aggs: {
            cats: {
                terms: {
                    field: 'category',
                    size: 10,
                    shard_size: 20,
                    min_doc_count: 1,
                    shard_min_doc_count: 0,
                    show_term_doc_count_error: true,
                    missing: 'N/A',
                    format: '##',
                    include: ['a', 'b'],
                    exclude: 'c*',
                    execution_hint: 'map',
                    collect_mode: 'depth_first',
                    order: { _count: 'desc' },
                    meta: { source: 'ui' },
                },
                aggs: {
                    avg_price: { avg: { field: 'price', missing: 0 } },
                },
            },
        },
    });
    assertIncludes(code, 'es.NewAggs(', 'aggs-only NewAggs');
    assertIncludes(code, 'es.TermsAgg("category")', 'terms agg');
    assertIncludes(code, 'ExecutionHint.Map', 'terms execution_hint');
    assertIncludes(code, 'CollectMode.DepthFirst', 'terms collect_mode');
    assertIncludes(code, 'Order.Desc', 'terms order');
    assertIncludes(code, 'Meta("source", "ui")', 'terms meta');
    assertIncludes(code, 'es.AvgAgg("price")', 'sub avg agg');
}

{
    const code = parser.parse({
        aggs: {
            s: { stats: { field: 'price' } },
            e: { extended_stats: { field: 'price' } },
            mi: { min: { field: 'price' } },
            ma: { max: { field: 'price' } },
            a: { avg: { field: 'price' } },
            su: { sum: { field: 'price' } },
            c: { cardinality: { field: 'user_id', precision_threshold: 100 } },
        },
    });
    assertIncludes(code, 'es.StatsAgg("price")', 'stats agg');
    assertIncludes(code, 'es.ExtendedStatsAgg("price")', 'extended_stats agg');
    assertIncludes(code, 'es.MinAgg("price")', 'min agg');
    assertIncludes(code, 'es.MaxAgg("price")', 'max agg');
    assertIncludes(code, 'es.AvgAgg("price")', 'avg agg');
    assertIncludes(code, 'es.SumAgg("price")', 'sum agg');
    assertIncludes(code, 'es.CardinalityAgg("user_id")', 'cardinality agg');
    assertIncludes(code, 'PrecisionThreshold(100)', 'cardinality precision');
}

{
    const code = parser.parse({
        aggs: {
            mt: {
                multi_terms: {
                    terms: [{ field: 'a' }, { field: 'b' }],
                    size: 5,
                    execution_hint: 'global_ordinals',
                    order: { _key: 'asc' },
                },
            },
        },
    });
    assertIncludes(code, 'es.MultiTermsAgg(es.TermAgg("a"), es.TermAgg("b"))', 'multi_terms');
    assertIncludes(code, 'ExecutionHint.GlobalOrdinals', 'multi_terms hint');
    assertIncludes(code, 'Order.Asc', 'multi_terms order');
}

{
    const code = parser.parse({
        aggs: {
            n: {
                nested: { path: 'offers' },
                aggs: {
                    rn: {
                        reverse_nested: { path: '' },
                        aggs: { t: { terms: { field: 'brand' } } },
                    },
                },
            },
        },
    });
    assertIncludes(code, 'es.NestedAgg("offers")', 'nested agg');
    assertIncludes(code, 'es.ReverseNestedAgg()', 'reverse_nested agg');
    assertIncludes(code, 'es.TermsAgg("brand")', 'nested sub terms');
}

{
    const code = parser.parse({
        aggs: {
            by_day: {
                date_histogram: {
                    field: 'date',
                    calendar_interval: 'day',
                    format: 'yyyy-MM-dd',
                    time_zone: '+03:00',
                    min_doc_count: 0,
                    keyed: true,
                    extended_bounds: { min: '2020-01-01', max: '2020-12-31' },
                    order: { _key: 'asc' },
                },
            },
        },
    });
    assertIncludes(code, 'es.DateHistogramAgg("date")', 'date_histogram');
    assertIncludes(code, 'CalendarInterval("day")', 'dh calendar');
    assertIncludes(code, 'TimeZone("+03:00")', 'dh tz');
    assertIncludes(code, 'es.AggOrder("_key", Order.Asc)', 'dh order');
}

{
    const code = parser.parse({
        aggs: {
            ranges: {
                date_range: {
                    field: 'date',
                    format: 'yyyy',
                    ranges: [{ key: 'old', to: '2020' }, { key: 'new', from: '2020' }],
                },
            },
        },
    });
    assertIncludes(code, 'es.DateRangeAgg("date")', 'date_range');
    assertIncludes(code, 'es.DateRangeEntry()', 'date_range entry');
    assertIncludes(code, '.Key("old")', 'date_range key');
}

{
    const code = parser.parse({
        aggs: {
            filtered: {
                filter: { term: { status: 'active' } },
                aggs: { avg_price: { avg: { field: 'price' } } },
            },
        },
    });
    assertIncludes(code, 'es.FilterAgg(', 'filter agg');
    assertIncludes(code, 'es.Term("status", "active")', 'filter agg query');
}

{
    const code = parser.parse({
        aggs: {
            messages: {
                filters: {
                    filters: {
                        errors: { term: { level: 'error' } },
                        warnings: { term: { level: 'warn' } },
                    },
                    other_bucket: true,
                    other_bucket_key: 'other',
                },
            },
        },
    });
    assertIncludes(code, 'es.FiltersAgg()', 'filters agg');
    assertIncludes(code, 'Filter("errors"', 'filters named');
    assertIncludes(code, 'OtherBucket(true)', 'filters other_bucket');
    assertIncludes(code, 'OtherBucketKey("other")', 'filters other_bucket_key');
}

{
    const code = parser.parse({
        aggs: {
            prices: {
                histogram: {
                    field: 'price',
                    interval: 50,
                    min_doc_count: 1,
                    offset: 0,
                    keyed: true,
                    extended_bounds: { min: 0, max: 500 },
                    order: { _key: 'desc' },
                },
            },
        },
    });
    assertIncludes(code, 'es.HistogramAgg("price", 50)', 'histogram');
    assertIncludes(code, 'ExtendedBounds(0, 500)', 'histogram bounds');
    assertIncludes(code, 'es.AggOrder("_key", Order.Desc)', 'histogram order');
}

{
    const code = parser.parse({
        aggs: {
            price_ranges: {
                range: {
                    field: 'price',
                    keyed: true,
                    ranges: [
                        { key: 'cheap', to: 50 },
                        { key: 'expensive', from: 50 },
                    ],
                },
            },
        },
    });
    assertIncludes(code, 'es.RangeAgg("price")', 'range agg');
    assertIncludes(code, 'es.RangeEntry()', 'range entry');
    assertIncludes(code, '.Key("cheap")', 'range entry key');
}

// =============================================================================
// New types (kept)
// =============================================================================

{
    const code = q({ fuzzy: { user: { value: 'ki', fuzziness: 'AUTO', max_expansions: 50, boost: 1.2 } } });
    assertIncludes(code, 'es.Fuzzy("user", "ki")', 'fuzzy');
}

{
    const code = q({ prefix: { 'user.id': { value: 'ki', case_insensitive: true } } });
    assertIncludes(code, 'es.Prefix("user.id", "ki")', 'prefix');
}

{
    const code = q({ wildcard: { 'user.id': { value: 'ki*y', rewrite: 'constant_score' } } });
    assertIncludes(code, 'es.Wildcard("user.id", "ki*y")', 'wildcard');
}

{
    const code = q({
        dis_max: {
            queries: [{ term: { title: 'quick' } }, { term: { body: 'brown' } }],
            tie_breaker: 0.7,
        },
    });
    assertIncludes(code, 'es.DisMax(', 'dis_max');
    assertIncludes(code, '.TieBreaker(0.7)', 'dis_max tie');
}

{
    const code = q({
        geo_distance: {
            distance: '12km',
            'pin.location': { lat: 40, lon: -70 },
            distance_type: 'plane',
            validation_method: 'COERCE',
        },
    });
    assertIncludes(code, 'es.GeoDistance("pin.location", 40, -70, "12km")', 'geo_distance');
    assertIncludes(code, 'DistanceType.Plane', 'geo distance_type');
    assertIncludes(code, 'ValidationMethod.Coerce', 'geo validation');
}

{
    const code = q({
        geo_bounding_box: {
            'pin.location': {
                top_left: '40.73, -74.1',
                bottom_right: [-71.12, 40.01],
            },
        },
    });
    assertIncludes(code, 'es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)', 'geo_bbox');
}

{
    const code = parser.parse({
        aggs: {
            vc: { value_count: { field: 'price', missing: 0 } },
            th: {
                top_hits: {
                    size: 3,
                    sort: [{ date: 'desc' }],
                    _source: false,
                    explain: true,
                    version: true,
                    seq_no_primary_term: true,
                    track_scores: true,
                },
            },
        },
    });
    assertIncludes(code, 'es.ValueCountAgg("price")', 'value_count');
    assertIncludes(code, 'es.TopHitsAgg()', 'top_hits');
    assertIncludes(code, 'SourceFalse()', 'top_hits source false');
    assertIncludes(code, 'SeqNoPrimaryTerm(true)', 'top_hits seq_no');
    assertIncludes(code, 'Order.Desc', 'top_hits sort string order');
}

console.log(`All tests passed (${passed} assertions).`);
