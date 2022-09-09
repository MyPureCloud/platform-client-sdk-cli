package webdeployments_deployments_configurations

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
	Description = utils.FormatUsageDescription("webdeployments_deployments_configurations", "SWAGGER_OVERRIDE_/api/v2/webdeployments/deployments/{deploymentId}/configurations", )
	webdeployments_deployments_configurationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_deployments_configurations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_deployments_configurationsCmd)
}

func Cmdwebdeployments_deployments_configurations() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "varType", "", "Get active configuration on a deployment")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/webdeployments/deployments/{deploymentId}/configurations", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Web Deployments", "/api/v2/webdeployments/deployments/{deploymentId}/configurations")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebDeploymentActiveConfigurationOnDeployment"
      }
    }
  }
}`)
	webdeployments_deployments_configurationsCmd.AddCommand(getCmd)
	return webdeployments_deployments_configurationsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [deploymentId]",
	Short: "Get active configuration for a given deployment",
	Long:  "Get active configuration for a given deployment",
	Args:  utils.DetermineArgs([]string{ "deploymentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/deployments/{deploymentId}/configurations"
		deploymentId, args := args[0], args[1:]
		path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)

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