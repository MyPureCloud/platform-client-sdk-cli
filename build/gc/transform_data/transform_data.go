package transform_data

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
)

var (
	TemplateFile string
	TemplateStr  string
)

func ConvertJsonToMap(data string) interface{} {
	// if the data is an array of json objects then we need to store that in an array of maps and not just a map
	if data[0] == '[' {
		var res []map[string]interface{}
		err := json.Unmarshal([]byte(data), &res)
		if err != nil {
			logger.Fatalf("Error unmarshalling %v\n", err)
		}
		return res
	}
	var res map[string]interface{}
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		logger.Fatalf("Error unmarshalling %v\n", err)
	}
	return res
}

func ProcessTemplateFile(mp interface{}) string {
	path := []string{TemplateFile}
	tmpl := handleParse(template.ParseFiles(path...))
	return process(tmpl, mp)
}

func ProcessTemplateStr(mp interface{}) string {
	tmpl := handleParse(template.New("tmpl").Parse(TemplateStr))
	return process(tmpl, mp)
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer
	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		logger.Fatalf("Error applying parsed template to data object: %v\n", err)
	}
	return tmplBytes.String()
}

// handleParse() provides similar functionality to Must() from the text/template package
// https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/text/template/helper.go;l=23;drc=4f1b0a44cb46f3df28f5ef82e5769ebeac1bc493
// but does not "panic" and print the stack trace if there is an error, instead it prints an error message and exits the application if the template syntax is not valid
// otherwise, it returns the template
func handleParse(t *template.Template, err error) *template.Template {
	if err != nil {
		logger.Fatalf("Error invalid syntax in %v\n", err)
	}
	return t
}
