package authorization_divisions_objects

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
	Description = utils.FormatUsageDescription("authorization_divisions_objects", "SWAGGER_OVERRIDE_/api/v2/authorization/divisions/{divisionId}/objects", )
	authorization_divisions_objectsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_divisions_objects"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_divisions_objectsCmd)
}

func Cmdauthorization_divisions_objects() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Authorization", "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Object Id List",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "type" : "string"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "The divisions were updated successfully",
  "content" : { }
}`)
	authorization_divisions_objectsCmd.AddCommand(createCmd)
	return authorization_divisions_objectsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [divisionId] [objectType]",
	Short: "Assign a list of objects to a division",
	Long:  "Assign a list of objects to a division",
	Args:  utils.DetermineArgs([]string{ "divisionId", "objectType", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.String{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
		objectType, args := args[0], args[1:]
		path = strings.Replace(path, "{objectType}", fmt.Sprintf("%v", objectType), -1)

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
