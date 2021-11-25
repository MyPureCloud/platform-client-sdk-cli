package users_bulk

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
	Description = utils.FormatUsageDescription("users_bulk", "SWAGGER_OVERRIDE_/api/v2/users/bulk", )
	users_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_bulkCmd)
}

func Cmdusers_bulk() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/users/bulk", utils.FormatPermissions([]string{ "directory:user:add", "directory:user:edit",  }), utils.GenerateDevCentreLink("PATCH", "Users", "/api/v2/users/bulk")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "Users",
  "required" : true,
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/PatchUser"
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/UserEntityListing"
  }
}`)
	users_bulkCmd.AddCommand(updateCmd)
	
	return users_bulkCmd
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update bulk acd autoanswer on users",
	Long:  "Update bulk acd autoanswer on users",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Patchuser{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/bulk"

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
