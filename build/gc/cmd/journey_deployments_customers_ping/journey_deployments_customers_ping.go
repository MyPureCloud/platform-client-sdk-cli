package journey_deployments_customers_ping

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
	Description = utils.FormatUsageDescription("journey_deployments_customers_ping", "SWAGGER_OVERRIDE_/api/v2/journey/deployments/{deploymentId}/customers/{customerCookieId}/ping", )
	journey_deployments_customers_pingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_deployments_customers_ping"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_deployments_customers_pingCmd)
}

func Cmdjourney_deployments_customers_ping() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "dl", "", "Document Location: 1) Web Page URL if overridden or URL fragment identifier (window.location.hash). OR  2) Application screen name that the ping request was sent from in the app. e.g. `home` or `help. Pings without this parameter will not return actions.")
	utils.AddFlag(getCmd.Flags(), "string", "dt", "", "Document Title.  A human readable name for the page or screen")
	utils.AddFlag(getCmd.Flags(), "string", "appNamespace", "", "Namespace of the application (e.g. com.genesys.bancodinero). Used for domain filtering in application sessions")
	utils.AddFlag(getCmd.Flags(), "string", "sessionId", "", "UUID of the customer session. Use the same Session Id for all pings, AppEvents and ActionEvents in the session")
	utils.AddFlag(getCmd.Flags(), "int", "sinceLastBeaconMilliseconds", "", "How long (milliseconds) since the last app event or beacon was sent. The response may return a pollInternvalMilliseconds to reduce the frequency of pings.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/deployments/{deploymentId}/customers/{customerCookieId}/ping", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/deployments/{deploymentId}/customers/{customerCookieId}/ping")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DeploymentPing"
      }
    }
  }
}`)
	journey_deployments_customers_pingCmd.AddCommand(getCmd)
	return journey_deployments_customers_pingCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [deploymentId] [customerCookieId]",
	Short: "Send a ping.",
	Long:  "Send a ping.",
	Args:  utils.DetermineArgs([]string{ "deploymentId", "customerCookieId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/deployments/{deploymentId}/customers/{customerCookieId}/ping"
		deploymentId, args := args[0], args[1:]
		path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
		customerCookieId, args := args[0], args[1:]
		path = strings.Replace(path, "{customerCookieId}", fmt.Sprintf("%v", customerCookieId), -1)

		dl := utils.GetFlag(cmd.Flags(), "string", "dl")
		if dl != "" {
			queryParams["dl"] = dl
		}
		dt := utils.GetFlag(cmd.Flags(), "string", "dt")
		if dt != "" {
			queryParams["dt"] = dt
		}
		appNamespace := utils.GetFlag(cmd.Flags(), "string", "appNamespace")
		if appNamespace != "" {
			queryParams["appNamespace"] = appNamespace
		}
		sessionId := utils.GetFlag(cmd.Flags(), "string", "sessionId")
		if sessionId != "" {
			queryParams["sessionId"] = sessionId
		}
		sinceLastBeaconMilliseconds := utils.GetFlag(cmd.Flags(), "int", "sinceLastBeaconMilliseconds")
		if sinceLastBeaconMilliseconds != "" {
			queryParams["sinceLastBeaconMilliseconds"] = sinceLastBeaconMilliseconds
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
