package externalcontacts_graphs_clusterscans_clusters

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
	Description = utils.FormatUsageDescription("externalcontacts_graphs_clusterscans_clusters", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters", )
	externalcontacts_graphs_clusterscans_clustersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_graphs_clusterscans_clusters"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_graphs_clusterscans_clustersCmd)
}

func Cmdexternalcontacts_graphs_clusterscans_clusters() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}", utils.FormatPermissions([]string{ "externalContacts:graphCluster:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Cluster"
      }
    }
  }
}`)
	externalcontacts_graphs_clusterscans_clustersCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "limit", "20", "Max number of records to return (must be between 1 and 100)")
	utils.AddFlag(listCmd.Flags(), "string", "cursor", "", "Cursor to continue scanning")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionIds", "", "which divisions to filter results to, up to 50 (defaults to all divisions use has access to)")
	utils.AddFlag(listCmd.Flags(), "string", "mergeInfoStatus", "", "which merge statuses to filter results to Valid values: AutoQueued, AutoSucceeded, AutoFailed, ManualQueued, ManualSucceeded, ManualFailed, NotMerged")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters", utils.FormatPermissions([]string{ "externalContacts:graphCluster:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	externalcontacts_graphs_clusterscans_clustersCmd.AddCommand(listCmd)
	return externalcontacts_graphs_clusterscans_clustersCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [scanId] [clusterId]",
	Short: "Returns a single cluster found by a scan",
	Long:  "Returns a single cluster found by a scan",
	Args:  utils.DetermineArgs([]string{ "scanId", "clusterId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}"
		scanId, args := args[0], args[1:]
		path = strings.Replace(path, "{scanId}", fmt.Sprintf("%v", scanId), -1)
		clusterId, args := args[0], args[1:]
		path = strings.Replace(path, "{clusterId}", fmt.Sprintf("%v", clusterId), -1)

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

		const opId = "get"
		const httpMethod = "GET"
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
var listCmd = &cobra.Command{
	Use:   "list [scanId]",
	Short: "Returns a list of clusters found by a scan",
	Long:  "Returns a list of clusters found by a scan",
	Args:  utils.DetermineArgs([]string{ "scanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters"
		scanId, args := args[0], args[1:]
		path = strings.Replace(path, "{scanId}", fmt.Sprintf("%v", scanId), -1)

		limit := utils.GetFlag(cmd.Flags(), "int", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
		}
		divisionIds := utils.GetFlag(cmd.Flags(), "[]string", "divisionIds")
		if divisionIds != "" {
			queryParams["divisionIds"] = divisionIds
		}
		mergeInfoStatus := utils.GetFlag(cmd.Flags(), "string", "mergeInfoStatus")
		if mergeInfoStatus != "" {
			queryParams["mergeInfo.status"] = mergeInfoStatus
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

		const opId = "list"
		const httpMethod = "GET"
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
