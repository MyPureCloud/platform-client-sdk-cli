package languageunderstanding_miners_uploads

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
	Description = utils.FormatUsageDescription("languageunderstanding_miners_uploads", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/miners/{minerId}/uploads", )
	languageunderstanding_miners_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_miners_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_miners_uploadsCmd)
}

func Cmdlanguageunderstanding_miners_uploads() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/miners/{minerId}/uploads", utils.FormatPermissions([]string{ "languageUnderstanding:miner:upload",  }), utils.GenerateDevCentreLink("POST", "Uploads", "/api/v2/languageunderstanding/miners/{minerId}/uploads")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;query&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Presigned URL successfully created.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UploadUrlResponse&quot;
  }
}`)
	languageunderstanding_miners_uploadsCmd.AddCommand(createCmd)
	
	return languageunderstanding_miners_uploadsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [minerId]",
	Short: "Creates a presigned URL for uploading a chat corpus which will be used for mining by intent miner",
	Long:  "Creates a presigned URL for uploading a chat corpus which will be used for mining by intent miner",
	Args:  utils.DetermineArgs([]string{ "minerId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/miners/{minerId}/uploads"
		minerId, args := args[0], args[1:]
		path = strings.Replace(path, "{minerId}", fmt.Sprintf("%v", minerId), -1)

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
