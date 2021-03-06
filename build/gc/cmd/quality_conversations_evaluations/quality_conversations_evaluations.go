package quality_conversations_evaluations

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("quality_conversations_evaluations", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/{conversationId}/evaluations", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/{conversationId}/evaluations", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/{conversationId}/evaluations", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/{conversationId}/evaluations", )
	quality_conversations_evaluationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_conversations_evaluations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_conversations_evaluationsCmd)
}

func Cmdquality_conversations_evaluations() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "expand", "", "evaluatorId")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/quality/conversations/{conversationId}/evaluations", utils.FormatPermissions([]string{ "quality:evaluation:add",  }), utils.GenerateDevCentreLink("POST", "Quality", "/api/v2/quality/conversations/{conversationId}/evaluations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;evaluation&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	quality_conversations_evaluationsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "string", "expand", "", "evaluatorId")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}", utils.FormatPermissions([]string{ "quality:evaluation:delete",  }), utils.GenerateDevCentreLink("DELETE", "Quality", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	quality_conversations_evaluationsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "agent, evaluator, evaluationForm")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}", utils.FormatPermissions([]string{ "quality:evaluation:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	quality_conversations_evaluationsCmd.AddCommand(getCmd)
	
	utils.AddFlag(updateCmd.Flags(), "string", "expand", "", "evaluatorId")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}", utils.FormatPermissions([]string{ "quality:evaluation:edit", "quality:evaluation:editScore", "quality:evaluation:editAgentSignoff",  }), utils.GenerateDevCentreLink("PUT", "Quality", "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;evaluation&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Evaluation&quot;
  }
}`)
	quality_conversations_evaluationsCmd.AddCommand(updateCmd)
	
	return quality_conversations_evaluationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId]",
	Short: "Create an evaluation",
	Long:  "Create an evaluation",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/{conversationId}/evaluations"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

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
	Use:   "delete [conversationId] [evaluationId]",
	Short: "Delete an evaluation",
	Long:  "Delete an evaluation",
	Args:  utils.DetermineArgs([]string{ "conversationId", "evaluationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		evaluationId, args := args[0], args[1:]
		path = strings.Replace(path, "{evaluationId}", fmt.Sprintf("%v", evaluationId), -1)

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
	Use:   "get [conversationId] [evaluationId]",
	Short: "Get an evaluation",
	Long:  "Get an evaluation",
	Args:  utils.DetermineArgs([]string{ "conversationId", "evaluationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		evaluationId, args := args[0], args[1:]
		path = strings.Replace(path, "{evaluationId}", fmt.Sprintf("%v", evaluationId), -1)

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
	Use:   "update [conversationId] [evaluationId]",
	Short: "Update an evaluation",
	Long:  "Update an evaluation",
	Args:  utils.DetermineArgs([]string{ "conversationId", "evaluationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		evaluationId, args := args[0], args[1:]
		path = strings.Replace(path, "{evaluationId}", fmt.Sprintf("%v", evaluationId), -1)

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
