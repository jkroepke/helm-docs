//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/norwoodj/helm-docs/pkg/document"
	"github.com/spf13/viper"
)

func renderValuesTableHTML(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return map[string]interface{}{
			"ok":    false,
			"error": "values.yaml input is required",
		}
	}

	html, err := document.RenderValuesTableHTMLFromYAML(args[0].String())
	if err != nil {
		return map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"ok":   true,
		"html": html,
	}
}

func main() {
	viper.Set("sort-values-order", document.AlphaNumSortOrder)
	viper.Set("sort-sections-order", document.FileSortOrder)

	js.Global().Set("renderHelmDocsValuesTable", js.FuncOf(renderValuesTableHTML))
	select {}
}
