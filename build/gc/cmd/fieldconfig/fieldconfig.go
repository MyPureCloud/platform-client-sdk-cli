package fieldconfig

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("fieldconfig", "SWAGGER_OVERRIDE_/api/v2/fieldconfig", )
	fieldconfigCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("fieldconfig"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(fieldconfigCmd)
}

func Cmdfieldconfig() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "varType", "", "Field type - REQUIRED Valid values: person, group, org, externalContact")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/fieldconfig", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Users", "/api/v2/fieldconfig")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("varType")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/FieldConfig&quot;
  }
}`)
	fieldconfigCmd.AddCommand(getCmd)
	
	return fieldconfigCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch field config for an entity type",
	Long:  "Fetch field config for an entity type",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/fieldconfig"

		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
