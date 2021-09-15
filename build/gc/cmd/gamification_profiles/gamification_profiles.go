package gamification_profiles

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
	Description = utils.FormatUsageDescription("gamification_profiles", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles", )
	gamification_profilesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profilesCmd)
}

func Cmdgamification_profiles() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/profiles", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/profiles")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;performanceProfile&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CreatePerformanceProfile&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GetProfilesResponse&quot;
  }
}`)
	gamification_profilesCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/profiles/{performanceProfileId}", utils.FormatPermissions([]string{ "gamification:profile:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/profiles/{performanceProfileId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/PerformanceProfile&quot;
  }
}`)
	gamification_profilesCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/profiles", utils.FormatPermissions([]string{ "gamification:profile:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/profiles")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GetProfilesResponse&quot;
  }
}`)
	gamification_profilesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/gamification/profiles/{performanceProfileId}", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("PUT", "Gamification", "/api/v2/gamification/profiles/{performanceProfileId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;performanceProfile&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/PerformanceProfile&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/PerformanceProfile&quot;
  }
}`)
	gamification_profilesCmd.AddCommand(updateCmd)
	
	return gamification_profilesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new custom performance profile",
	Long:  "Create a new custom performance profile",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles"

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
	Use:   "get [performanceProfileId]",
	Short: "Performance profile by id",
	Long:  "Performance profile by id",
	Args:  utils.DetermineArgs([]string{ "performanceProfileId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{performanceProfileId}"
		performanceProfileId, args := args[0], args[1:]
		path = strings.Replace(path, "{performanceProfileId}", fmt.Sprintf("%v", performanceProfileId), -1)

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
	Use:   "list",
	Short: "All performance profiles",
	Long:  "All performance profiles",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles"

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
	Use:   "update [performanceProfileId]",
	Short: "Updates a performance profile",
	Long:  "Updates a performance profile",
	Args:  utils.DetermineArgs([]string{ "performanceProfileId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{performanceProfileId}"
		performanceProfileId, args := args[0], args[1:]
		path = strings.Replace(path, "{performanceProfileId}", fmt.Sprintf("%v", performanceProfileId), -1)

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
