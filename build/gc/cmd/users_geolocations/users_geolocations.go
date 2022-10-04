package users_geolocations

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
	Description = utils.FormatUsageDescription("users_geolocations", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/geolocations", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/geolocations", )
	users_geolocationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_geolocations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_geolocationsCmd)
}

func Cmdusers_geolocations() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/{userId}/geolocations/{clientId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Geolocation", "/api/v2/users/{userId}/geolocations/{clientId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Geolocation"
      }
    }
  }
}`)
	users_geolocationsCmd.AddCommand(getCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/geolocations/{clientId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Geolocation", "/api/v2/users/{userId}/geolocations/{clientId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Geolocation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Geolocation"
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
        "$ref" : "#/components/schemas/Geolocation"
      }
    }
  }
}`)
	users_geolocationsCmd.AddCommand(updateCmd)
	return users_geolocationsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [userId] [clientId]",
	Short: "Get a user`s Geolocation",
	Long:  "Get a user`s Geolocation",
	Args:  utils.DetermineArgs([]string{ "userId", "clientId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/geolocations/{clientId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
		clientId, args := args[0], args[1:]
		path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [userId] [clientId]",
	Short: "Patch a user`s Geolocation",
	Long:  "Patch a user`s Geolocation",
	Args:  utils.DetermineArgs([]string{ "userId", "clientId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Geolocation{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/geolocations/{clientId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
		clientId, args := args[0], args[1:]
		path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientId), -1)

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
