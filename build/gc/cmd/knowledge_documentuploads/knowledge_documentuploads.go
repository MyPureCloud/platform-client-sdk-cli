package knowledge_documentuploads

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
	Description = utils.FormatUsageDescription("knowledge_documentuploads", "SWAGGER_OVERRIDE_/api/v2/knowledge/documentuploads", )
	knowledge_documentuploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_documentuploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_documentuploadsCmd)
}

func Cmdknowledge_documentuploads() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/documentuploads", utils.FormatPermissions([]string{ "knowledge:document:upload",  }), utils.GenerateDevCentreLink("POST", "Knowledge", "/api/v2/knowledge/documentuploads")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;query&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UploadUrlRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Presigned URL successfully created.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UploadUrlResponse&quot;
  }
}`)
	knowledge_documentuploadsCmd.AddCommand(createCmd)
	
	return knowledge_documentuploadsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a presigned URL for uploading a knowledge import file with a set of documents",
	Long:  "Creates a presigned URL for uploading a knowledge import file with a set of documents",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/documentuploads"

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
