package journey_actionmaps

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
	Description = utils.FormatUsageDescription("journey_actionmaps", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps", )
	journey_actionmapsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_actionmaps"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_actionmapsCmd)
}

func Cmdjourney_actionmaps() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/journey/actionmaps", utils.FormatPermissions([]string{ "journey:actionmap:add",  }), utils.GenerateDevCentreLink("POST", "Journey", "/api/v2/journey/actionmaps")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ActionMap"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ActionMap"
      }
    }
  }
}`)
	journey_actionmapsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/journey/actionmaps/{actionMapId}", utils.FormatPermissions([]string{ "journey:actionmap:delete",  }), utils.GenerateDevCentreLink("DELETE", "Journey", "/api/v2/journey/actionmaps/{actionMapId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Action map deleted.",
  "content" : { }
}`)
	journey_actionmapsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/actionmaps/{actionMapId}", utils.FormatPermissions([]string{ "journey:actionmap:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/actionmaps/{actionMapId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ActionMap"
      }
    }
  }
}`)
	journey_actionmapsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Field(s) to sort by. Prefix with `-` for descending (e.g. sortBy=displayName,-createdDate).")
	utils.AddFlag(listCmd.Flags(), "string", "filterField", "", "Field to filter by (e.g. filterField=weight or filterField=action.actionTemplate.id). Requires `filterField` to also be set.")
	utils.AddFlag(listCmd.Flags(), "string", "filterValue", "", "Value to filter by. Requires `filterValue` to also be set.")
	utils.AddFlag(listCmd.Flags(), "[]string", "actionMapIds", "", "IDs of action maps to return. Use of this parameter is not compatible with pagination, filtering, sorting or querying. A maximum of 100 action maps are allowed per request.")
	utils.AddFlag(listCmd.Flags(), "[]string", "queryFields", "", "Action Map field(s) to query on. Requires `queryValue` to also be set.")
	utils.AddFlag(listCmd.Flags(), "string", "queryValue", "", "Value to query on. Requires `queryFields` to also be set.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/journey/actionmaps", utils.FormatPermissions([]string{ "journey:actionmap:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/actionmaps")))
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
	journey_actionmapsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/journey/actionmaps/{actionMapId}", utils.FormatPermissions([]string{ "journey:actionmap:edit",  }), utils.GenerateDevCentreLink("PATCH", "Journey", "/api/v2/journey/actionmaps/{actionMapId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PatchActionMap"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ActionMap"
      }
    }
  }
}`)
	journey_actionmapsCmd.AddCommand(updateCmd)
	return journey_actionmapsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an action map.",
	Long:  "Create an action map.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Actionmap{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/actionmaps"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [actionMapId]",
	Short: "Delete single action map.",
	Long:  "Delete single action map.",
	Args:  utils.DetermineArgs([]string{ "actionMapId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/actionmaps/{actionMapId}"
		actionMapId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionMapId}", fmt.Sprintf("%v", actionMapId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "delete"
		const httpMethod = "DELETE"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
	Use:   "get [actionMapId]",
	Short: "Retrieve a single action map.",
	Long:  "Retrieve a single action map.",
	Args:  utils.DetermineArgs([]string{ "actionMapId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/actionmaps/{actionMapId}"
		actionMapId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionMapId}", fmt.Sprintf("%v", actionMapId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve all action maps.",
	Long:  "Retrieve all action maps.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/actionmaps"

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
		filterField := utils.GetFlag(cmd.Flags(), "string", "filterField")
		if filterField != "" {
			queryParams["filterField"] = filterField
		}
		filterValue := utils.GetFlag(cmd.Flags(), "string", "filterValue")
		if filterValue != "" {
			queryParams["filterValue"] = filterValue
		}
		actionMapIds := utils.GetFlag(cmd.Flags(), "[]string", "actionMapIds")
		if actionMapIds != "" {
			queryParams["actionMapIds"] = actionMapIds
		}
		queryFields := utils.GetFlag(cmd.Flags(), "[]string", "queryFields")
		if queryFields != "" {
			queryParams["queryFields"] = queryFields
		}
		queryValue := utils.GetFlag(cmd.Flags(), "string", "queryValue")
		if queryValue != "" {
			queryParams["queryValue"] = queryValue
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [actionMapId]",
	Short: "Update single action map.",
	Long:  "Update single action map.",
	Args:  utils.DetermineArgs([]string{ "actionMapId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Patchactionmap{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/actionmaps/{actionMapId}"
		actionMapId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionMapId}", fmt.Sprintf("%v", actionMapId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PATCH"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
