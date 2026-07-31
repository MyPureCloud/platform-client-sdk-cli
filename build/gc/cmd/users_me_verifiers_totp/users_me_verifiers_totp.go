package users_me_verifiers_totp

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
	Description = utils.FormatUsageDescription("users_me_verifiers_totp", "SWAGGER_OVERRIDE_/api/v2/users/me/verifiers/totp", "SWAGGER_OVERRIDE_/api/v2/users/me/verifiers/totp", )
	users_me_verifiers_totpCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_me_verifiers_totp"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_me_verifiers_totpCmd)
}

func Cmdusers_me_verifiers_totp() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/users/me/verifiers/totp", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Users", "/api/v2/users/me/verifiers/totp")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Verifier",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateVerifierRequest"
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
        "$ref" : "#/components/schemas/CreateVerifierResponse"
      }
    }
  }
}`)
	users_me_verifiers_totpCmd.AddCommand(createCmd)

	validateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", validateCmd.UsageTemplate(), "POST", "/api/v2/users/me/verifiers/totp/{verifierId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Users", "/api/v2/users/me/verifiers/totp/{verifierId}")))
	utils.AddFileFlagIfUpsert(validateCmd.Flags(), "POST", `{
  "description" : "Verifier Validate",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ValidateVerifierRequest"
      }
    }
  },
  "required" : true
}`)
	
	
	utils.AddPaginateFlagsIfListingResponse(validateCmd.Flags(), "POST", `{
  "description" : "Verifier validated successfully",
  "content" : { }
}`)
	users_me_verifiers_totpCmd.AddCommand(validateCmd)
	return users_me_verifiers_totpCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new TOTP verifier",
	Long:  "Add a new TOTP verifier",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createverifierrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/me/verifiers/totp"

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

		const opId = "create"
		const httpMethod = "POST"
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
var validateCmd = &cobra.Command{
	Use:   "validate [verifierId]",
	Short: "Validate a TOTP verifier",
	Long:  "Validate a TOTP verifier",
	Args:  utils.DetermineArgs([]string{ "verifierId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Validateverifierrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/me/verifiers/totp/{verifierId}"
		verifierId, args := args[0], args[1:]
		path = strings.Replace(path, "{verifierId}", fmt.Sprintf("%v", verifierId), -1)

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

		const opId = "validate"
		const httpMethod = "POST"
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
