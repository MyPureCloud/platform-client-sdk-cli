package recording_keyconfigurations_validate

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
	Description = utils.FormatUsageDescription("recording_keyconfigurations_validate", "SWAGGER_OVERRIDE_/api/v2/recording/keyconfigurations/validate", )
	recording_keyconfigurations_validateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_keyconfigurations_validate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_keyconfigurations_validateCmd)
}

func Cmdrecording_keyconfigurations_validate() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/recording/keyconfigurations/validate", utils.FormatPermissions([]string{ "recording:encryptionKey:edit",  }), utils.GenerateDevCentreLink("POST", "Recording", "/api/v2/recording/keyconfigurations/validate")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Encryption Configuration",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/RecordingEncryptionConfiguration"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/RecordingEncryptionConfiguration"
  }
}`)
	recording_keyconfigurations_validateCmd.AddCommand(createCmd)
	
	return recording_keyconfigurations_validateCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Validate encryption key configurations without saving it",
	Long:  "Validate encryption key configurations without saving it",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Recordingencryptionconfiguration{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/keyconfigurations/validate"

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
