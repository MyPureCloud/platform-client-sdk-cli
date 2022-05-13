package authorization_divisionspermitted_me

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
	Description = utils.FormatUsageDescription("authorization_divisionspermitted_me", "SWAGGER_OVERRIDE_/api/v2/authorization/divisionspermitted/me", )
	authorization_divisionspermitted_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_divisionspermitted_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_divisionspermitted_meCmd)
}

func Cmdauthorization_divisionspermitted_me() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "name", "", "Search term to filter by division name")
	utils.AddFlag(getCmd.Flags(), "string", "permission", "", "The permission string, including the object to access, e.g. routing:queue:view - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/divisionspermitted/me", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/divisionspermitted/me")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("permission")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/AuthzDivision"
        }
      }
    }
  }
}`)
	authorization_divisionspermitted_meCmd.AddCommand(getCmd)
	return authorization_divisionspermitted_meCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns which divisions the current user has the given permission in.",
	Long:  "Returns which divisions the current user has the given permission in.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisionspermitted/me"

		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		permission := utils.GetFlag(cmd.Flags(), "string", "permission")
		if permission != "" {
			queryParams["permission"] = permission
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
