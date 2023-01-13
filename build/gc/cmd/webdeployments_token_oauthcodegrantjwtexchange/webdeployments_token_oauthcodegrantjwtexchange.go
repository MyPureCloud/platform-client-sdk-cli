package webdeployments_token_oauthcodegrantjwtexchange

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
	Description = utils.FormatUsageDescription("webdeployments_token_oauthcodegrantjwtexchange", "SWAGGER_OVERRIDE_/api/v2/webdeployments/token/oauthcodegrantjwtexchange", )
	webdeployments_token_oauthcodegrantjwtexchangeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_token_oauthcodegrantjwtexchange"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_token_oauthcodegrantjwtexchangeCmd)
}

func Cmdwebdeployments_token_oauthcodegrantjwtexchange() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webdeployments/token/oauthcodegrantjwtexchange", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Web Deployments", "/api/v2/webdeployments/token/oauthcodegrantjwtexchange")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "webDeploymentsOAuthExchangeRequest",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebDeploymentsOAuthExchangeRequest"
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
        "$ref" : "#/components/schemas/WebDeploymentsAuthorizationResponse"
      }
    }
  }
}`)
	webdeployments_token_oauthcodegrantjwtexchangeCmd.AddCommand(createCmd)
	return webdeployments_token_oauthcodegrantjwtexchangeCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Exchange an oAuth code (obtained using the Authorization Code Flow) for a JWT that can be used by webdeployments.",
	Long:  "Exchange an oAuth code (obtained using the Authorization Code Flow) for a JWT that can be used by webdeployments.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Webdeploymentsoauthexchangerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/token/oauthcodegrantjwtexchange"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
