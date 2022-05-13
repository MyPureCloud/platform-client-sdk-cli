package scripts_published_variables

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("scripts_published_variables", "SWAGGER_OVERRIDE_/api/v2/scripts/published/{scriptId}/variables", )
	scripts_published_variablesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_published_variables"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_published_variablesCmd)
}

func Cmdscripts_published_variables() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "input", "", "input")
	utils.AddFlag(getCmd.Flags(), "string", "output", "", "output")
	utils.AddFlag(getCmd.Flags(), "string", "varType", "", "type")
	utils.AddFlag(getCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/scripts/published/{scriptId}/variables", utils.FormatPermissions([]string{ "scripter:publishedScript:view",  }), utils.GenerateDevCentreLink("GET", "Scripts", "/api/v2/scripts/published/{scriptId}/variables")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "object"
      }
    }
  }
}`)
	scripts_published_variablesCmd.AddCommand(getCmd)
	return scripts_published_variablesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [scriptId]",
	Short: "Get the published variables",
	Long:  "Get the published variables",
	Args:  utils.DetermineArgs([]string{ "scriptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scripts/published/{scriptId}/variables"
		scriptId, args := args[0], args[1:]
		path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)

		input := utils.GetFlag(cmd.Flags(), "string", "input")
		if input != "" {
			queryParams["input"] = input
		}
		output := utils.GetFlag(cmd.Flags(), "string", "output")
		if output != "" {
			queryParams["output"] = output
		}
		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
