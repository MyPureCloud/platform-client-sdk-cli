package externalcontacts_graphs_clusterscans_clusters_merge

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
	Description = utils.FormatUsageDescription("externalcontacts_graphs_clusterscans_clusters_merge", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}/merge", )
	externalcontacts_graphs_clusterscans_clusters_mergeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_graphs_clusterscans_clusters_merge"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_graphs_clusterscans_clusters_mergeCmd)
}

func Cmdexternalcontacts_graphs_clusterscans_clusters_merge() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}/merge", utils.FormatPermissions([]string{ "externalContacts:graphCluster:merge",  }), utils.GenerateDevCentreLink("PUT", "External Contacts", "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}/merge")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Cluster"
      }
    }
  }
}`)
	externalcontacts_graphs_clusterscans_clusters_mergeCmd.AddCommand(updateCmd)
	return externalcontacts_graphs_clusterscans_clusters_mergeCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var updateCmd = &cobra.Command{
	Use:   "update [scanId] [clusterId]",
	Short: "Merge a single cluster found by a scan",
	Long:  "Merge a single cluster found by a scan",
	Args:  utils.DetermineArgs([]string{ "scanId", "clusterId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters/{clusterId}/merge"
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

		const opId = "update"
		const httpMethod = "PUT"
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
