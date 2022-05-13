package responsemanagement_responseassets_uploads

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
	Description = utils.FormatUsageDescription("responsemanagement_responseassets_uploads", "SWAGGER_OVERRIDE_/api/v2/responsemanagement/responseassets/uploads", )
	responsemanagement_responseassets_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("responsemanagement_responseassets_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(responsemanagement_responseassets_uploadsCmd)
}

func Cmdresponsemanagement_responseassets_uploads() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/responsemanagement/responseassets/uploads", utils.FormatPermissions([]string{ "responseAssets:asset:add",  }), utils.GenerateDevCentreLink("POST", "Response Management", "/api/v2/responsemanagement/responseassets/uploads")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateResponseAssetRequest"
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
        "$ref" : "#/components/schemas/CreateResponseAssetResponse"
      }
    }
  }
}`)
	responsemanagement_responseassets_uploadsCmd.AddCommand(createCmd)
	return responsemanagement_responseassets_uploadsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates pre-signed url for uploading response asset",
	Long:  "Creates pre-signed url for uploading response asset",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createresponseassetrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/responsemanagement/responseassets/uploads"

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
