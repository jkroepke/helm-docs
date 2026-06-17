package document

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderValuesTableHTMLFromYAML(t *testing.T) {
	viper.Set("sort-values-order", AlphaNumSortOrder)
	viper.Set("sort-sections-order", FileSortOrder)

	valuesYAML := `# image.repository -- Container image repository
image:
  repository: nginx
  # image.tag -- Container image tag
  tag: "1.27"
`

	html, err := RenderValuesTableHTMLFromYAML(valuesYAML)
	require.NoError(t, err)

	assert.Contains(t, html, "<table>")
	assert.Contains(t, html, "<td>image.repository</td>")
	assert.Contains(t, html, "<td>Container image repository</td>")
	assert.Contains(t, html, "<td>image.tag</td>")
	assert.Contains(t, html, "<td>Container image tag</td>")
	assert.False(t, strings.Contains(html, "Chart.yaml"))
}
