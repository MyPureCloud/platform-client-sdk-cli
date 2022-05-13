package scripts_uploads_status

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
	Description = utils.FormatUsageDescription("scripts_uploads_status", "SWAGGER_OVERRIDE_/api/v2/scripts/uploads/{uploadId}/status", )
	scripts_uploads_statusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_uploads_status"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_uploads_statusCmd)
}

func Cmdscripts_uploads_status() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "longPoll", "false", "Enable longPolling endpoint")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/scripts/uploads/{uploadId}/status", utils.FormatPermissions([]string{ "scripter:script:search",  }), utils.GenerateDevCentreLink("GET", "Scripts", "/api/v2/scripts/uploads/{uploadId}/status")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ImportScriptStatusResponse"
      }
    }
  }
}`)
	scripts_uploads_statusCmd.AddCommand(getCmd)
	return scripts_uploads_statusCmd
}

var getCmd = &cobra.Command{
	Use:   "get [uploadId]",
	Short: "Get the upload status of an imported script",
	Long:  "Get the upload status of an imported script",
	Args:  utils.DetermineArgs([]string{ "uploadId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scripts/uploads/{uploadId}/status"
		uploadId, args := args[0], args[1:]
		path = strings.Replace(path, "{uploadId}", fmt.Sprintf("%v", uploadId), -1)

		longPoll := utils.GetFlag(cmd.Flags(), "bool", "longPoll")
		if longPoll != "" {
			queryParams["longPoll"] = longPoll
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
