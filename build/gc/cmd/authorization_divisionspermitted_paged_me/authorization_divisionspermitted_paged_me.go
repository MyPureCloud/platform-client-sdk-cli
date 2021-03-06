package authorization_divisionspermitted_paged_me

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
	Description = utils.FormatUsageDescription("authorization_divisionspermitted_paged_me", "SWAGGER_OVERRIDE_/api/v2/authorization/divisionspermitted/paged/me", )
	authorization_divisionspermitted_paged_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_divisionspermitted_paged_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_divisionspermitted_paged_meCmd)
}

func Cmdauthorization_divisionspermitted_paged_me() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "permission", "", "The permission string, including the object to access, e.g. routing:queue:view - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/authorization/divisionspermitted/paged/me", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/divisionspermitted/paged/me")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("permission")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	authorization_divisionspermitted_paged_meCmd.AddCommand(listCmd)
	
	return authorization_divisionspermitted_paged_meCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns which divisions the current user has the given permission in.",
	Long:  "Returns which divisions the current user has the given permission in.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisionspermitted/paged/me"

		permission := utils.GetFlag(cmd.Flags(), "string", "permission")
		if permission != "" {
			queryParams["permission"] = permission
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
