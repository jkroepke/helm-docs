helm-docs
=========
The helm-docs tool generates automatic documentation from helm charts into a markdown file. The resulting
file contains metadata about the chart and a table with all of your charts' values, their defaults, and an
optional description parsed from comments.

The markdown generation is entirely [gotemplate](https://golang.org/pkg/text/template) driven. The tool parses metadata
from charts and generates a number of sub-templates that can be referenced in a template file (by default `README.md.gotmpl`).
If no template file is provided, the tool has a default internal template that will generate a reasonably formatted README.


## Usage
To build from source:
```bash
cd cmd/helm-docs
go build
```

You can also download the latest release for a variety of platforms from [here](https://github.com/norwoodj/helm-docs/releases/latest).

To run:
```bash
helm-docs
# OR
helm-docs --dry-run # prints generated documentation to stdout rather than modifying READMEs
```

The tool searches recursively through subdirectories of the current directory for `Chart.yaml` files and generates documentation
for every chart that it finds.


## Available Templates
The templates generated by the tool are shown below, and can be included in your `README.md.gotmpl` file like so:
```
{{ template "template-name" . }}
```

| Name | Description |
|------|-------------|
| chart.header              | The main heading of the generated markdown file |
| chart.description         | A description line containing the _description_ field from the chart's `Chart.yaml` file, or "" if that field is not set |
| chart.version             | The _version_ field from the chart's `Chart.yaml` file |
| chart.versionLine         | A text line stating the current version of the chart |
| chart.sourceLink          | The _home_ link from the chart's `Chart.yaml` file, or "" if that field is not set |
| chart.sourceLinkLine      | A text line with the _home_ link from the chart's `Chart.yaml` file, or "" if that field is not set |
| chart.requirementsHeader  | The heading for the chart requirements section |
| chart.requirementsTable   | A table of the chart's required sub-charts |
| chart.requirementsSection | A section headed by the requirementsHeader from above containing the requirementsTable from above or "" if there are no requirements |
| chart.valuesHeader        | The heading for the chart values section |
| chart.valuesTable         | A table of the chart's values parsed from the `values.yaml` file (see below) |
| chart.valuesSection       | A section headed by the valuesHeader from above containing the valuesTable from above or "" if there are no values |

For an example of how these various templates can be used in a `README.md.gotmpl` file to generate a reasonable markdown file,
look at the charts in [example-charts](./example-charts).

If there is no `README.md.gotmpl` (or other specified gotmpl file) present, the default template is used to generate the README.
That template looks like so:
```
{{ template "chart.header" . }}
{{ template "chart.description" . }}

{{ template "chart.versionLine" . }}

{{ template "chart.sourceLinkLine" . }}

{{ template "chart.requirementsSection" . }}

{{ template "chart.valuesSection" . }}
```

The tool includes the [sprig templating library](https://github.com/Masterminds/sprig), so those functions can be used
in the templates you supply.


## values.yaml metadata
This tool can parse descriptions and defaults of values from `values.yaml` files. The defaults are pulled directly from
the yaml in the file. Descriptions can be added for parameters by specifying the full path of the value and
a particular comment format. I invite you to check out the [example-charts](./example-charts) to see how this is done in
practice. In order to add a description for a parameter you need only put a comment somewhere in the file of the format:

```yaml
controller:
  publishService:
    # controller.publishService.enabled -- Whether to expose the ingress controller to the public world
    enabled: false

  # controller.replicas -- Number of nginx-ingress pods to load balance between
  replicas: 2
```

The descriptions will be picked up and put in the table in the README. The comment need not be near the parameter it
explains, although this is probably preferable.

_Note:_ if the value in question contains any `.` characters, that section of the path must be quoted e.g.
```yaml
service:
  annotations:
    # ingress.annotations."external-dns.alpha.kubernetes.io/hostname" -- Hostname to be assigned to the ELB for the service
    external-dns.alpha.kubernetes.io/hostname: stupidchess.jmn23.com
```

### nil values
If you would like to define a key for a value, but leave the default empty, you can still specify a description for it
as well as a type. Like so:
```yaml
controller:
  # controller.replicas -- (int) Number of nginx-ingress pods to load balance between
  replicas:
```
This could be useful when wanting to enforce user-defined values for the chart, where there are no sensible defaults.
