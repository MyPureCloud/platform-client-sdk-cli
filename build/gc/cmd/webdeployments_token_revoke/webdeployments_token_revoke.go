package webdeployments_token_revoke

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
	Description = utils.FormatUsageDescription("webdeployments_token_revoke", "SWAGGER_OVERRIDE_/api/v2/webdeployments/token/revoke", )
	webdeployments_token_revokeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_token_revoke"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_token_revokeCmd)
}

func Cmdwebdeployments_token_revoke() *cobra.Command { 
	utils.AddFlag(deleteCmd.Flags(), "string", "xJourneySessionId", "", "The Customer`s journey sessionId.")
	utils.AddFlag(deleteCmd.Flags(), "string", "xJourneySessionType", "", "The Customer`s journey session type.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/webdeployments/token/revoke", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Web Deployments", "/api/v2/webdeployments/token/revoke")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Revoke any tokens associate with the JWT. \nIf the JWT is expired the refresh endpoint should be called to obtain a valid JWT and this endpoint called again.",
  "content" : { }
}`)
	webdeployments_token_revokeCmd.AddCommand(deleteCmd)
	return webdeployments_token_revokeCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Invalidate JWT",
	Long:  "Invalidate JWT",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/token/revoke"

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

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		// header params "X-Journey-Session-Id"
		xJourneySessionId := utils.GetFlag(cmd.Flags(), "string", "xJourneySessionId")
		if xJourneySessionId != "" {
			headerParams["X-Journey-Session-Id"] = xJourneySessionId
		}
		// header params "X-Journey-Session-Type"
		xJourneySessionType := utils.GetFlag(cmd.Flags(), "string", "xJourneySessionType")
		if xJourneySessionType != "" {
			headerParams["X-Journey-Session-Type"] = xJourneySessionType
		}


		const opId = "delete"
		const httpMethod = "DELETE"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
