package workforcemanagement_businessunits_adherence_explanations_query

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_adherence_explanations_query", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query", )
	workforcemanagement_businessunits_adherence_explanations_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_adherence_explanations_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_adherence_explanations_queryCmd)
}

func Cmdworkforcemanagement_businessunits_adherence_explanations_query() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "forceAsync", "", "Force the result of this operation to be sent asynchronously via notification. For testing/app development purposes")
	utils.AddFlag(createCmd.Flags(), "bool", "forceDownloadService", "", "Force the result of this operation to be sent via download service. For testing/app development purposes")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query", utils.FormatPermissions([]string{ "wfm:adherenceExplanation:view",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The request body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BuQueryAdherenceExplanationsRequest"
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
        "$ref" : "#/components/schemas/BuQueryAdherenceExplanationsResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunits_adherence_explanations_queryCmd.AddCommand(createCmd)
	return workforcemanagement_businessunits_adherence_explanations_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId]",
	Short: "Query adherence explanations across an entire business unit for the requested period",
	Long:  "Query adherence explanations across an entire business unit for the requested period",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Buqueryadherenceexplanationsrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

		forceAsync := utils.GetFlag(cmd.Flags(), "bool", "forceAsync")
		if forceAsync != "" {
			queryParams["forceAsync"] = forceAsync
		}
		forceDownloadService := utils.GetFlag(cmd.Flags(), "bool", "forceDownloadService")
		if forceDownloadService != "" {
			queryParams["forceDownloadService"] = forceDownloadService
		}
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
