package languages_translations_builtin

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
	Description = utils.FormatUsageDescription("languages_translations_builtin", "SWAGGER_OVERRIDE_/api/v2/languages/translations/builtin", )
	languages_translations_builtinCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languages_translations_builtin"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languages_translations_builtinCmd)
}

func Cmdlanguages_translations_builtin() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "language", "", "The language of the builtin translation to retrieve - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languages/translations/builtin", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Languages", "/api/v2/languages/translations/builtin")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("language")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "object",
        "additionalProperties" : {
          "type" : "object",
          "properties" : { }
        }
      }
    }
  }
}`)
	languages_translations_builtinCmd.AddCommand(getCmd)
	return languages_translations_builtinCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the builtin translation for a language",
	Long:  "Get the builtin translation for a language",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languages/translations/builtin"

		language := utils.GetFlag(cmd.Flags(), "string", "language")
		if language != "" {
			queryParams["language"] = language
		}
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
