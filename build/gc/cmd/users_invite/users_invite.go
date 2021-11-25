package users_invite

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
	Description = utils.FormatUsageDescription("users_invite", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/invite", )
	users_inviteCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_invite"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_inviteCmd)
}

func Cmdusers_invite() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "force", "false", "Resend the invitation even if one is already outstanding")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/users/{userId}/invite", utils.FormatPermissions([]string{ "directory:user:add",  }), utils.GenerateDevCentreLink("POST", "Users", "/api/v2/users/{userId}/invite")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Invitation Sent"
}`)
	users_inviteCmd.AddCommand(createCmd)
	
	return users_inviteCmd
}

var createCmd = &cobra.Command{
	Use:   "create [userId]",
	Short: "Send an activation email to the user",
	Long:  "Send an activation email to the user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/invite"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		force := utils.GetFlag(cmd.Flags(), "bool", "force")
		if force != "" {
			queryParams["force"] = force
		}
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
