package conversations_participants_secureivrsessions

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
	Description = utils.FormatUsageDescription("conversations_participants_secureivrsessions", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions", )
	conversations_participants_secureivrsessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_secureivrsessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_secureivrsessionsCmd)
}

func Cmdconversations_participants_secureivrsessions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CreateSecureSession&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SecureSession&quot;
  }
}`)
	conversations_participants_secureivrsessionsCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SecureSession&quot;
  }
}`)
	conversations_participants_secureivrsessionsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SecureSessionEntityListing&quot;
  }
}`)
	conversations_participants_secureivrsessionsCmd.AddCommand(listCmd)
	
	return conversations_participants_secureivrsessionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [participantId]",
	Short: "Create secure IVR session. Only a participant in the conversation can invoke a secure IVR.",
	Long:  "Create secure IVR session. Only a participant in the conversation can invoke a secure IVR.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)

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
var getCmd = &cobra.Command{
	Use:   "get [conversationId] [participantId] [secureSessionId]",
	Short: "Fetch info on a secure session",
	Long:  "Fetch info on a secure session",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", "secureSessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)
		secureSessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{secureSessionId}", fmt.Sprintf("%v", secureSessionId), -1)

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
	Use:   "list [conversationId] [participantId]",
	Short: "Get a list of secure sessions for this participant.",
	Long:  "Get a list of secure sessions for this participant.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)

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
