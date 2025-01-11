package profiles

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"strings"

	"github.com/spf13/cobra"
)

var listProfilesCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns all configured profiles",
	Long:  "Returns all configured profiles",
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		configs, err := config.ListConfigs()
		if err != nil {
			logger.Fatal(err)
		}

		configStrings := make([]string, 0)
		for _, config := range configs {
			configStrings = append(configStrings, config.String())
		}

		var finalJSONString string
		switch len(configStrings) {
		case 0:
			finalJSONString = "[]"
		case 1:
			finalJSONString = fmt.Sprintf("[%s]", configStrings)
		default:
			finalJSONString = fmt.Sprintf("[%s]", strings.Join(configStrings, ","))
		}

		utils.Render(finalJSONString)
	},
}
