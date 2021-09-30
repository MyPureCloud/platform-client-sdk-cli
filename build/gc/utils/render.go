package utils

import (
	"fmt"
	"strings"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"sigs.k8s.io/yaml"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/data_format"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/transform_data"
	"github.com/tidwall/pretty"
)

func Render(data string) {
	if strings.EqualFold("yaml", data_format.OutputFormat) && isJSON(data) {
		result, err := yaml.JSONToYAML([]byte(data))
		if err != nil {
			logger.Fatalf("Error converting JSON to YAML: %v\n", err)
		}
		fmt.Println("---")
		fmt.Printf("%s", result)
		return
	}
	if transform_data.TemplateFile != "" {
		mp := transform_data.ConvertJsonToMap(data)
		res := transform_data.ProcessTemplateFile(mp)
		fmt.Println(res)
		return
	}
	if transform_data.TemplateStr != "" {
		mp := transform_data.ConvertJsonToMap(data)
		res := transform_data.ProcessTemplateStr(mp)
		fmt.Println(res)
		return
	}
	result := pretty.Pretty([]byte(data))
	fmt.Printf("%s", result)
}
