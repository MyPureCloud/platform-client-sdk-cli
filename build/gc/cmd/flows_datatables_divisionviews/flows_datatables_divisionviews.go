package flows_datatables_divisionviews

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
	Description = utils.FormatUsageDescription("flows_datatables_divisionviews", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/divisionviews", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/divisionviews", )
	flows_datatables_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_divisionviewsCmd)
}

func Cmdflows_datatables_divisionviews() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Expand instructions for the result Valid values: schema")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/divisionviews/{datatableId}", utils.FormatPermissions([]string{ "architect:datatable:search",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/divisionviews/{datatableId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DataTable"
      }
    }
  }
}`)
	flows_datatables_divisionviewsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "Expand instructions for the result Valid values: schema")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "id", "Sort by Valid values: id, name")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "Sort order")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "division ID(s)")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name to filter by")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/divisionviews", utils.FormatPermissions([]string{ "architect:datatable:search",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/divisionviews")))
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
	flows_datatables_divisionviewsCmd.AddCommand(listCmd)
	return flows_datatables_divisionviewsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [datatableId]",
	Short: "Returns a specific datatable by id",
	Long:  "Returns a specific datatable by id",
	Args:  utils.DetermineArgs([]string{ "datatableId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/divisionviews/{datatableId}"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
	Use:   "list",
	Short: "Retrieve a list of datatables for the org",
	Long:  "Retrieve a list of datatables for the org",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/divisionviews"

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
