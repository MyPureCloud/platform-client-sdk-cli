package presencedefinitions

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
	Description = utils.FormatUsageDescription("presencedefinitions", "SWAGGER_OVERRIDE_/api/v2/presencedefinitions", "SWAGGER_OVERRIDE_/api/v2/presencedefinitions", "SWAGGER_OVERRIDE_/api/v2/presencedefinitions", "SWAGGER_OVERRIDE_/api/v2/presencedefinitions", "SWAGGER_OVERRIDE_/api/v2/presencedefinitions", )
	presencedefinitionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("presencedefinitions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(presencedefinitionsCmd)
}

func Cmdpresencedefinitions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/presencedefinitions", utils.FormatPermissions([]string{ "presence:presenceDefinition:add",  }), utils.GenerateDevCentreLink("POST", "Presence", "/api/v2/presencedefinitions")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The Presence Definition to create",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrganizationPresence"
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
        "$ref" : "#/components/schemas/OrganizationPresence"
      }
    }
  }
}`)
	presencedefinitionsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/presencedefinitions/{presenceId}", utils.FormatPermissions([]string{ "presence:presenceDefinition:delete",  }), utils.GenerateDevCentreLink("DELETE", "Presence", "/api/v2/presencedefinitions/{presenceId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ErrorBody"
      }
    }
  },
  "x-inin-error-codes" : {
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	presencedefinitionsCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "string", "localeCode", "", "The locale code to fetch for the presence definition. Use ALL to fetch everything.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/presencedefinitions/{presenceId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Presence", "/api/v2/presencedefinitions/{presenceId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrganizationPresence"
      }
    }
  }
}`)
	presencedefinitionsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "deleted", "false", "Deleted query can be TRUE, FALSE or ALL")
	utils.AddFlag(listCmd.Flags(), "string", "localeCode", "", "The locale code to fetch for each presence definition. Use ALL to fetch everything.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/presencedefinitions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Presence", "/api/v2/presencedefinitions")))
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
	presencedefinitionsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/presencedefinitions/{presenceId}", utils.FormatPermissions([]string{ "presence:presenceDefinition:edit",  }), utils.GenerateDevCentreLink("PUT", "Presence", "/api/v2/presencedefinitions/{presenceId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "The OrganizationPresence to update",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrganizationPresence"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrganizationPresence"
      }
    }
  }
}`)
	presencedefinitionsCmd.AddCommand(updateCmd)
	return presencedefinitionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Presence Definition",
	Long:  "Create a Presence Definition",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Organizationpresence{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/presencedefinitions"

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
	Use:   "delete [presenceId]",
	Short: "Delete a Presence Definition",
	Long:  "Delete a Presence Definition",
	Args:  utils.DetermineArgs([]string{ "presenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/presencedefinitions/{presenceId}"
		presenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{presenceId}", fmt.Sprintf("%v", presenceId), -1)

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
	Use:   "get [presenceId]",
	Short: "Get a Presence Definition",
	Long:  "Get a Presence Definition",
	Args:  utils.DetermineArgs([]string{ "presenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/presencedefinitions/{presenceId}"
		presenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{presenceId}", fmt.Sprintf("%v", presenceId), -1)

		localeCode := utils.GetFlag(cmd.Flags(), "string", "localeCode")
		if localeCode != "" {
			queryParams["localeCode"] = localeCode
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get an Organization`s list of Presence Definitions",
	Long:  "Get an Organization`s list of Presence Definitions",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/presencedefinitions"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		deleted := utils.GetFlag(cmd.Flags(), "string", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
		}
		localeCode := utils.GetFlag(cmd.Flags(), "string", "localeCode")
		if localeCode != "" {
			queryParams["localeCode"] = localeCode
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
	Use:   "update [presenceId]",
	Short: "Update a Presence Definition",
	Long:  "Update a Presence Definition",
	Args:  utils.DetermineArgs([]string{ "presenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Organizationpresence{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/presencedefinitions/{presenceId}"
		presenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{presenceId}", fmt.Sprintf("%v", presenceId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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
