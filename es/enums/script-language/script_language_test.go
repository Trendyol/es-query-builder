package scriptlanguage_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ScriptLanguageString(t *testing.T) {
	tests := []struct {
		scriptLanguage ScriptLanguage.ScriptLanguage
		result         string
	}{
		{ScriptLanguage.Painless, "painless"},
		{ScriptLanguage.Expression, "expression"},
		{ScriptLanguage.Mustache, "mustache"},
		{ScriptLanguage.Java, "java"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.scriptLanguage.String())
		})
	}
}
