package utils

import (
	"fmt"
	"strings"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"sigs.k8s.io/yaml"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/data_format"
	"github.com/tidwall/pretty"
)

func Render(data string) {
	if strings.EqualFold("yaml", data_format.OutputFormat) && isJSON(data) {
		result, err := yaml.JSONToYAML([]byte(data))
		if err != nil {
			logger.Fatalf("Error converting JSON to YAML: %v\n", err)
		}
		fmt.Printf("%s", result)
		return
	}
	result := pretty.Pretty([]byte(data))
	fmt.Printf("%s", result)
}
