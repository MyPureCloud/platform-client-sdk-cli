package flows_datatables_rows

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
	Description = utils.FormatUsageDescription("flows_datatables_rows", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/rows", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/rows", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/rows", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/rows", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/rows", )
	flows_datatables_rowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_rows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_rowsCmd)
}

func Cmdflows_datatables_rows() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/datatables/{datatableId}/rows", utils.FormatPermissions([]string{ "architect:datatable:add",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/datatables/{datatableId}/rows")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "dataTableRow",
  "required" : true,
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	flows_datatables_rowsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}", utils.FormatPermissions([]string{ "architect:datatable:delete",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The row was deleted successfully"
}`)
	flows_datatables_rowsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "bool", "showbrief", "true", "if true returns just the key field for the row")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}", utils.FormatPermissions([]string{ "architect:datatable:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	flows_datatables_rowsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "showbrief", "true", "If true returns just the key value of the row")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/{datatableId}/rows", utils.FormatPermissions([]string{ "architect:datatable:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/{datatableId}/rows")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	flows_datatables_rowsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}", utils.FormatPermissions([]string{ "architect:datatable:edit",  }), utils.GenerateDevCentreLink("PUT", "Architect", "/api/v2/flows/datatables/{datatableId}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "datatable row",
  "required" : false,
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	flows_datatables_rowsCmd.AddCommand(updateCmd)
	
	return flows_datatables_rowsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [datatableId]",
	Short: "Create a new row entry for the datatable.",
	Long:  "Create a new row entry for the datatable.",
	Args:  utils.DetermineArgs([]string{ "datatableId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interface{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/rows"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [datatableId] [rowId]",
	Short: "Delete a row entry",
	Long:  "Delete a row entry",
	Args:  utils.DetermineArgs([]string{ "datatableId", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/rows/{rowId}"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [datatableId] [rowId]",
	Short: "Returns a specific row for the datatable",
	Long:  "Returns a specific row for the datatable",
	Args:  utils.DetermineArgs([]string{ "datatableId", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/rows/{rowId}"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)


		showbrief := utils.GetFlag(cmd.Flags(), "bool", "showbrief")
		if showbrief != "" {
			queryParams["showbrief"] = showbrief
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list [datatableId]",
	Short: "Returns the rows for the datatable with the given id",
	Long:  "Returns the rows for the datatable with the given id",
	Args:  utils.DetermineArgs([]string{ "datatableId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/rows"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)


		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		showbrief := utils.GetFlag(cmd.Flags(), "bool", "showbrief")
		if showbrief != "" {
			queryParams["showbrief"] = showbrief
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [datatableId] [rowId]",
	Short: "Update a row entry",
	Long:  "Update a row entry",
	Args:  utils.DetermineArgs([]string{ "datatableId", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interface{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/rows/{rowId}"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
