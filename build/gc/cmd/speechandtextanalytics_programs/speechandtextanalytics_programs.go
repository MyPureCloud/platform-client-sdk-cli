package speechandtextanalytics_programs

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_programs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs", )
	speechandtextanalytics_programsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_programs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_programsCmd)
}

func Cmdspeechandtextanalytics_programs() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/speechandtextanalytics/programs", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:add",  }), utils.GenerateDevCentreLink("POST", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The program to create&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ProgramRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Program&quot;
  }
}`)
	speechandtextanalytics_programsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "forceDelete", "false", "Indicates whether the program is forced to be deleted or not. Required when the program to delete is the default program.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/speechandtextanalytics/programs/{programId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:delete",  }), utils.GenerateDevCentreLink("DELETE", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs/{programId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;The program was deleted successfully&quot;
}`)
	speechandtextanalytics_programsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/programs/{programId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs/{programId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Program&quot;
  }
}`)
	speechandtextanalytics_programsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "The key for listing the next page")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "20", "The page size for the listing")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/programs", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	speechandtextanalytics_programsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/speechandtextanalytics/programs/{programId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:program:edit",  }), utils.GenerateDevCentreLink("PUT", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/programs/{programId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The program to update&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ProgramRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Program&quot;
  }
}`)
	speechandtextanalytics_programsCmd.AddCommand(updateCmd)
	
	return speechandtextanalytics_programsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Speech and Text Analytics program",
	Long:  "Create new Speech and Text Analytics program",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs"

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
	Use:   "delete [programId]",
	Short: "Delete a Speech and Text Analytics program by id",
	Long:  "Delete a Speech and Text Analytics program by id",
	Args:  utils.DetermineArgs([]string{ "programId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs/{programId}"
		programId, args := args[0], args[1:]
		path = strings.Replace(path, "{programId}", fmt.Sprintf("%v", programId), -1)

		forceDelete := utils.GetFlag(cmd.Flags(), "bool", "forceDelete")
		if forceDelete != "" {
			queryParams["forceDelete"] = forceDelete
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
	Use:   "get [programId]",
	Short: "Get a Speech and Text Analytics program by id",
	Long:  "Get a Speech and Text Analytics program by id",
	Args:  utils.DetermineArgs([]string{ "programId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs/{programId}"
		programId, args := args[0], args[1:]
		path = strings.Replace(path, "{programId}", fmt.Sprintf("%v", programId), -1)

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
	Short: "Get the list of Speech and Text Analytics programs",
	Long:  "Get the list of Speech and Text Analytics programs",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs"

		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
	Use:   "update [programId]",
	Short: "Update existing Speech and Text Analytics program",
	Long:  "Update existing Speech and Text Analytics program",
	Args:  utils.DetermineArgs([]string{ "programId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/programs/{programId}"
		programId, args := args[0], args[1:]
		path = strings.Replace(path, "{programId}", fmt.Sprintf("%v", programId), -1)

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