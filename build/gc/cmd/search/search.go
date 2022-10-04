package search

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
	Description = utils.FormatUsageDescription("search", "SWAGGER_OVERRIDE_/api/v2/search", "SWAGGER_OVERRIDE_/api/v2/search", )
	searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(searchCmd)
}

func Cmdsearch() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "profile", "true", "profile")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/search", utils.FormatPermissions([]string{ "directory:user:view",  }), utils.GenerateDevCentreLink("POST", "Search", "/api/v2/search")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Search request options",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SearchRequest"
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
        "$ref" : "#/components/schemas/JsonNodeSearchResponse"
      }
    }
  }
}`)
	searchCmd.AddCommand(createCmd)

	utils.AddFlag(listCmd.Flags(), "string", "q64", "", "q64 - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, dateLastLogin, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography, callerUser.routingStatus, callerUser.primaryPresence, callerUser.conversationSummary, callerUser.outOfOffice, callerUser.geolocation, images, addressVerificationDetails")
	utils.AddFlag(listCmd.Flags(), "bool", "profile", "true", "profile")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/search", utils.FormatPermissions([]string{ "directory:user:view",  }), utils.GenerateDevCentreLink("GET", "Search", "/api/v2/search")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("q64")
	
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
	searchCmd.AddCommand(listCmd)
	return searchCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Search resources.",
	Long:  "Search resources.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Searchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/search"

		profile := utils.GetFlag(cmd.Flags(), "bool", "profile")
		if profile != "" {
			queryParams["profile"] = profile
		}
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
	Short: "Search using the q64 value returned from a previous search.",
	Long:  "Search using the q64 value returned from a previous search.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/search"

		q64 := utils.GetFlag(cmd.Flags(), "string", "q64")
		if q64 != "" {
			queryParams["q64"] = q64
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		profile := utils.GetFlag(cmd.Flags(), "bool", "profile")
		if profile != "" {
			queryParams["profile"] = profile
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
