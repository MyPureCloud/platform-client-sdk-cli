package outbound_contactlistfilters_preview

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
	Description = utils.FormatUsageDescription("outbound_contactlistfilters_preview", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlistfilters/preview", )
	outbound_contactlistfilters_previewCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlistfilters_preview"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlistfilters_previewCmd)
}

func Cmdoutbound_contactlistfilters_preview() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/contactlistfilters/preview", utils.FormatPermissions([]string{ "outbound:contact:view",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/contactlistfilters/preview")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "ContactListFilter",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/ContactListFilter"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/FilterPreviewResponse"
  }
}`)
	outbound_contactlistfilters_previewCmd.AddCommand(createCmd)
	
	return outbound_contactlistfilters_previewCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get a preview of the output of a contact list filter",
	Long:  "Get a preview of the output of a contact list filter",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Contactlistfilter{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlistfilters/preview"

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
