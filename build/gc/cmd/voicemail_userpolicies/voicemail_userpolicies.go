package voicemail_userpolicies

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
	Description = utils.FormatUsageDescription("voicemail_userpolicies", "SWAGGER_OVERRIDE_/api/v2/voicemail/userpolicies", "SWAGGER_OVERRIDE_/api/v2/voicemail/userpolicies", "SWAGGER_OVERRIDE_/api/v2/voicemail/userpolicies", )
	voicemail_userpoliciesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_userpolicies"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_userpoliciesCmd)
}

func Cmdvoicemail_userpolicies() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/voicemail/userpolicies/{userId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Voicemail", "/api/v2/voicemail/userpolicies/{userId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/VoicemailUserPolicy"
      }
    }
  }
}`)
	voicemail_userpoliciesCmd.AddCommand(getCmd)

	putCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", putCmd.UsageTemplate(), "PUT", "/api/v2/voicemail/userpolicies/{userId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("PUT", "Voicemail", "/api/v2/voicemail/userpolicies/{userId}")))
	utils.AddFileFlagIfUpsert(putCmd.Flags(), "PUT", `{
  "description" : "The user's voicemail policy",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/VoicemailUserPolicy"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(putCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/VoicemailUserPolicy"
      }
    }
  }
}`)
	voicemail_userpoliciesCmd.AddCommand(putCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/voicemail/userpolicies/{userId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Voicemail", "/api/v2/voicemail/userpolicies/{userId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "The user's voicemail policy",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/VoicemailUserPolicy"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/VoicemailUserPolicy"
      }
    }
  }
}`)
	voicemail_userpoliciesCmd.AddCommand(updateCmd)
	return voicemail_userpoliciesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [userId]",
	Short: "Get a user`s voicemail policy",
	Long:  "Get a user`s voicemail policy",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/voicemail/userpolicies/{userId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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
var putCmd = &cobra.Command{
	Use:   "put [userId]",
	Short: "Update a user`s voicemail policy",
	Long:  "Update a user`s voicemail policy",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Voicemailuserpolicy{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/voicemail/userpolicies/{userId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "put"
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
var updateCmd = &cobra.Command{
	Use:   "update [userId]",
	Short: "Update a user`s voicemail policy",
	Long:  "Update a user`s voicemail policy",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Voicemailuserpolicy{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/voicemail/userpolicies/{userId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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
