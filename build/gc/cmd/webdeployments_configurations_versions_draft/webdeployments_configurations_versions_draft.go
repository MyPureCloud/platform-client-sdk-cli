package webdeployments_configurations_versions_draft

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
	Description = utils.FormatUsageDescription("webdeployments_configurations_versions_draft", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations/{configurationId}/versions/draft", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations/{configurationId}/versions/draft", )
	webdeployments_configurations_versions_draftCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_configurations_versions_draft"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_configurations_versions_draftCmd)
}

func Cmdwebdeployments_configurations_versions_draft() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft", utils.FormatPermissions([]string{ "webDeployments:configuration:view",  }), utils.GenerateDevCentreLink("GET", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebDeploymentConfigurationVersion"
      }
    }
  }
}`)
	webdeployments_configurations_versions_draftCmd.AddCommand(getCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft", utils.FormatPermissions([]string{ "webDeployments:configuration:edit",  }), utils.GenerateDevCentreLink("PUT", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebDeploymentConfigurationVersion"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebDeploymentConfigurationVersion"
      }
    }
  }
}`)
	webdeployments_configurations_versions_draftCmd.AddCommand(updateCmd)
	return webdeployments_configurations_versions_draftCmd
}

var getCmd = &cobra.Command{
	Use:   "get [configurationId]",
	Short: "Get the configuration draft",
	Long:  "Get the configuration draft",
	Args:  utils.DetermineArgs([]string{ "configurationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}/versions/draft"
		configurationId, args := args[0], args[1:]
		path = strings.Replace(path, "{configurationId}", fmt.Sprintf("%v", configurationId), -1)

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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [configurationId]",
	Short: "Update the configuration draft",
	Long:  "Update the configuration draft",
	Args:  utils.DetermineArgs([]string{ "configurationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Webdeploymentconfigurationversion{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}/versions/draft"
		configurationId, args := args[0], args[1:]
		path = strings.Replace(path, "{configurationId}", fmt.Sprintf("%v", configurationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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

		utils.Render(results)
	},
}
