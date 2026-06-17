package document

import (
	"bytes"
	"text/template"

	"github.com/norwoodj/helm-docs/pkg/helm"
	"github.com/norwoodj/helm-docs/pkg/util"
)

func RenderValuesTableHTMLFromYAML(valuesYAML string) (string, error) {
	values, err := helm.ParseChartValuesFromYAML(valuesYAML)
	if err != nil {
		return "", err
	}

	descriptions, err := helm.ParseChartValuesCommentsFromYAML(
		&values,
		valuesYAML,
		helm.ChartValuesDocumentationParsingConfig{},
	)
	if err != nil {
		return "", err
	}

	valuesTableRows, err := getUnsortedValueRows(&values, descriptions)
	if err != nil {
		return "", err
	}

	sortValueRows(valuesTableRows)
	sectionedRows := getSectionedValueRows(valuesTableRows)
	sortSectionedValueRows(sectionedRows)

	tmpl := template.New("values-table-html")
	tmpl.Funcs(util.FuncMap())
	if _, err := tmpl.Parse(getValuesTableTemplates()); err != nil {
		return "", err
	}
	if _, err := tmpl.Parse(`{{ template "chart.valuesTableHtml" . }}`); err != nil {
		return "", err
	}

	data := chartTemplateData{
		Values:   valuesTableRows,
		Sections: sectionedRows,
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		return "", err
	}

	return output.String(), nil
}
