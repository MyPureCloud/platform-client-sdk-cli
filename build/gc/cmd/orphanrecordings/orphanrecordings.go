package orphanrecordings

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
	Description = utils.FormatUsageDescription("orphanrecordings", "SWAGGER_OVERRIDE_/api/v2/orphanrecordings", "SWAGGER_OVERRIDE_/api/v2/orphanrecordings", "SWAGGER_OVERRIDE_/api/v2/orphanrecordings", "SWAGGER_OVERRIDE_/api/v2/orphanrecordings", )
	orphanrecordingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orphanrecordings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orphanrecordingsCmd)
}

func Cmdorphanrecordings() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/orphanrecordings/{orphanId}", utils.FormatPermissions([]string{ "recording:orphan:delete",  }), utils.GenerateDevCentreLink("DELETE", "Recording", "/api/v2/orphanrecordings/{orphanId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrphanRecording"
      }
    }
  }
}`)
	orphanrecordingsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orphanrecordings/{orphanId}", utils.FormatPermissions([]string{ "recording:orphan:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/orphanrecordings/{orphanId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrphanRecording"
      }
    }
  }
}`)
	orphanrecordingsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "bool", "hasConversation", "false", "Filter resulting orphans by whether the conversation is known. False returns all orphans for the organization.")
	utils.AddFlag(listCmd.Flags(), "string", "media", "", "Filter resulting orphans based on their media type Valid values: Call, Screen")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/orphanrecordings", utils.FormatPermissions([]string{ "recording:orphan:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/orphanrecordings")))
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
	orphanrecordingsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orphanrecordings/{orphanId}", utils.FormatPermissions([]string{ "recording:orphan:edit",  }), utils.GenerateDevCentreLink("PUT", "Recording", "/api/v2/orphanrecordings/{orphanId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrphanUpdateRequest"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Recording"
      }
    }
  }
}`)
	orphanrecordingsCmd.AddCommand(updateCmd)
	return orphanrecordingsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [orphanId]",
	Short: "Deletes a single orphan recording",
	Long:  "Deletes a single orphan recording",
	Args:  utils.DetermineArgs([]string{ "orphanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orphanrecordings/{orphanId}"
		orphanId, args := args[0], args[1:]
		path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)

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
	Use:   "get [orphanId]",
	Short: "Gets a single orphan recording",
	Long:  "Gets a single orphan recording",
	Args:  utils.DetermineArgs([]string{ "orphanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orphanrecordings/{orphanId}"
		orphanId, args := args[0], args[1:]
		path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)

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
	Short: "Gets all orphan recordings",
	Long:  "Gets all orphan recordings",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orphanrecordings"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		previousPage := utils.GetFlag(cmd.Flags(), "string", "previousPage")
		if previousPage != "" {
			queryParams["previousPage"] = previousPage
		}
		hasConversation := utils.GetFlag(cmd.Flags(), "bool", "hasConversation")
		if hasConversation != "" {
			queryParams["hasConversation"] = hasConversation
		}
		media := utils.GetFlag(cmd.Flags(), "string", "media")
		if media != "" {
			queryParams["media"] = media
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
	Use:   "update [orphanId]",
	Short: "Updates an orphan recording to a regular recording with retention values",
	Long:  "Updates an orphan recording to a regular recording with retention values",
	Args:  utils.DetermineArgs([]string{ "orphanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Orphanupdaterequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orphanrecordings/{orphanId}"
		orphanId, args := args[0], args[1:]
		path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)

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
