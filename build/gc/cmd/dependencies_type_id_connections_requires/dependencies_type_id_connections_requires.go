package dependencies_type_id_connections_requires

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
	Description = utils.FormatUsageDescription("dependencies_type_id_connections_requires", "SWAGGER_OVERRIDE_/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires", )
	dependencies_type_id_connections_requiresCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dependencies_type_id_connections_requires"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dependencies_type_id_connections_requiresCmd)
}

func Cmddependencies_type_id_connections_requires() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "25", "Page size (max 100)")
	utils.AddFlag(listCmd.Flags(), "string", "beforeSourceType", "", "Cursor for previous page")
	utils.AddFlag(listCmd.Flags(), "string", "beforeSourceId", "", "Cursor for previous page")
	utils.AddFlag(listCmd.Flags(), "string", "afterSourceType", "", "Cursor for next page")
	utils.AddFlag(listCmd.Flags(), "string", "afterSourceId", "", "Cursor for next page")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires", utils.FormatPermissions([]string{ "dependencies:dependency:view",  }), utils.GenerateDevCentreLink("GET", "Dependencies", "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires")))
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
	dependencies_type_id_connections_requiresCmd.AddCommand(listCmd)
	return dependencies_type_id_connections_requiresCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [entityType] [entityId]",
	Short: "Get entities that the given entity requires",
	Long:  "Get entities that the given entity requires",
	Args:  utils.DetermineArgs([]string{ "entityType", "entityId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires"
		entityType, args := args[0], args[1:]
		path = strings.Replace(path, "{entityType}", fmt.Sprintf("%v", entityType), -1)
		entityId, args := args[0], args[1:]
		path = strings.Replace(path, "{entityId}", fmt.Sprintf("%v", entityId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		beforeSourceType := utils.GetFlag(cmd.Flags(), "string", "beforeSourceType")
		if beforeSourceType != "" {
			queryParams["beforeSourceType"] = beforeSourceType
		}
		beforeSourceId := utils.GetFlag(cmd.Flags(), "string", "beforeSourceId")
		if beforeSourceId != "" {
			queryParams["beforeSourceId"] = beforeSourceId
		}
		afterSourceType := utils.GetFlag(cmd.Flags(), "string", "afterSourceType")
		if afterSourceType != "" {
			queryParams["afterSourceType"] = afterSourceType
		}
		afterSourceId := utils.GetFlag(cmd.Flags(), "string", "afterSourceId")
		if afterSourceId != "" {
			queryParams["afterSourceId"] = afterSourceId
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
