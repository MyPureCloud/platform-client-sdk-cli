package flows_actions_publish

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
	Description = utils.FormatUsageDescription("flows_actions_publish", "SWAGGER_OVERRIDE_/api/v2/flows/actions/publish", )
	flows_actions_publishCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_actions_publish"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_actions_publishCmd)
}

func Cmdflows_actions_publish() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "flow", "", "Flow ID - REQUIRED")
	utils.AddFlag(createCmd.Flags(), "string", "version", "", "version")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/actions/publish", utils.FormatPermissions([]string{ "architect:flow:unlock", "architect:flow:publish",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/actions/publish")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	createCmd.MarkFlagRequired("flow")
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Operation"
      }
    }
  }
}`)
	flows_actions_publishCmd.AddCommand(createCmd)
	return flows_actions_publishCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Publish flow",
	Long:  "Publish flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/actions/publish"

		flow := utils.GetFlag(cmd.Flags(), "string", "flow")
		if flow != "" {
			queryParams["flow"] = flow
		}
		version := utils.GetFlag(cmd.Flags(), "string", "version")
		if version != "" {
			queryParams["version"] = version
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "create"
		const httpMethod = "POST"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
