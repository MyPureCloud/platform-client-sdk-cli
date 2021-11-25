package outbound_schedules_sequences

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
	Description = utils.FormatUsageDescription("outbound_schedules_sequences", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/sequences", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/sequences", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/sequences", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/sequences", )
	outbound_schedules_sequencesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_schedules_sequences"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_schedules_sequencesCmd)
}

func Cmdoutbound_schedules_sequences() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/schedules/sequences/{sequenceId}", utils.FormatPermissions([]string{ "outbound:schedule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/schedules/sequences/{sequenceId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful."
}`)
	outbound_schedules_sequencesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/schedules/sequences/{sequenceId}", utils.FormatPermissions([]string{ "outbound:schedule:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/schedules/sequences/{sequenceId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SequenceSchedule"
  }
}`)
	outbound_schedules_sequencesCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/schedules/sequences", utils.FormatPermissions([]string{ "outbound:schedule:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/schedules/sequences")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/SequenceSchedule"
    }
  }
}`)
	outbound_schedules_sequencesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/schedules/sequences/{sequenceId}", utils.FormatPermissions([]string{ "outbound:schedule:edit",  }), utils.GenerateDevCentreLink("PUT", "Outbound", "/api/v2/outbound/schedules/sequences/{sequenceId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "SequenceSchedule",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/SequenceSchedule"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SequenceSchedule"
  }
}`)
	outbound_schedules_sequencesCmd.AddCommand(updateCmd)
	
	return outbound_schedules_sequencesCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [sequenceId]",
	Short: "Delete a dialer sequence schedule.",
	Long:  "Delete a dialer sequence schedule.",
	Args:  utils.DetermineArgs([]string{ "sequenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/sequences/{sequenceId}"
		sequenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{sequenceId}", fmt.Sprintf("%v", sequenceId), -1)

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
	Use:   "get [sequenceId]",
	Short: "Get a dialer sequence schedule.",
	Long:  "Get a dialer sequence schedule.",
	Args:  utils.DetermineArgs([]string{ "sequenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/sequences/{sequenceId}"
		sequenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{sequenceId}", fmt.Sprintf("%v", sequenceId), -1)

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
	Short: "Query for a list of dialer sequence schedules.",
	Long:  "Query for a list of dialer sequence schedules.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/sequences"

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
	Use:   "update [sequenceId]",
	Short: "Update a new sequence schedule.",
	Long:  "Update a new sequence schedule.",
	Args:  utils.DetermineArgs([]string{ "sequenceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Sequenceschedule{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/sequences/{sequenceId}"
		sequenceId, args := args[0], args[1:]
		path = strings.Replace(path, "{sequenceId}", fmt.Sprintf("%v", sequenceId), -1)

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
