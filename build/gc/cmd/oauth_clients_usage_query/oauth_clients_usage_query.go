package oauth_clients_usage_query

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
	Description = utils.FormatUsageDescription("oauth_clients_usage_query", "SWAGGER_OVERRIDE_/api/v2/oauth/clients/{clientId}/usage/query", )
	oauth_clients_usage_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("oauth_clients_usage_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(oauth_clients_usage_queryCmd)
}

func Cmdoauth_clients_usage_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/oauth/clients/{clientId}/usage/query", utils.FormatPermissions([]string{ "oauth:client:view",  }), utils.GenerateDevCentreLink("POST", "OAuth", "/api/v2/oauth/clients/{clientId}/usage/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Query",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ApiUsageQuery"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UsageExecutionResult"
      }
    }
  }
}`)
	oauth_clients_usage_queryCmd.AddCommand(createCmd)
	return oauth_clients_usage_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create [clientId]",
	Short: "Query for OAuth client API usage",
	Long:  "Query for OAuth client API usage",
	Args:  utils.DetermineArgs([]string{ "clientId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Apiusagequery{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/oauth/clients/{clientId}/usage/query"
		clientId, args := args[0], args[1:]
		path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
