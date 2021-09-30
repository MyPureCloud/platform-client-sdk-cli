package authorization_subjects_divisions_roles

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
	Description = utils.FormatUsageDescription("authorization_subjects_divisions_roles", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles", )
	authorization_subjects_divisions_rolesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_subjects_divisions_roles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_subjects_divisions_rolesCmd)
}

func Cmdauthorization_subjects_divisions_roles() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "subjectType", "PC_USER", "what the type of the subject is: PC_GROUP, PC_USER or PC_OAUTH_CLIENT (note: for cross-org authorization, please use the Organization Authorization endpoints)")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:grant:add",  }), utils.GenerateDevCentreLink("POST", "Authorization", "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "schema" : {
    "$ref" : "#/definitions/ErrorBody"
  },
  "x-inin-error-codes" : {
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable",
    "invalid.date" : "Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ",
    "invalid.query.param.value" : "Value [%s] is not valid for parameter [%s]. Allowable values are: %s",
    "invalid.property" : "Value [%s] is not a valid property for object [%s]",
    "constraint.validation" : "%s",
    "invalid.value" : "Value [%s] is not valid for field type [%s]. Allowable values are: %s"
  }
}`)
	authorization_subjects_divisions_rolesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:grant:delete",  }), utils.GenerateDevCentreLink("DELETE", "Authorization", "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "schema" : {
    "$ref" : "#/definitions/ErrorBody"
  },
  "x-inin-error-codes" : {
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	authorization_subjects_divisions_rolesCmd.AddCommand(deleteCmd)
	
	return authorization_subjects_divisions_rolesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [subjectId] [divisionId] [roleId]",
	Short: "Make a grant of a role in a division",
	Long:  "Make a grant of a role in a division",
	Args:  utils.DetermineArgs([]string{ "subjectId", "divisionId", "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}"
		subjectId, args := args[0], args[1:]
		path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)


		subjectType := utils.GetFlag(cmd.Flags(), "string", "subjectType")
		if subjectType != "" {
			queryParams["subjectType"] = subjectType
		}
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
var deleteCmd = &cobra.Command{
	Use:   "delete [subjectId] [divisionId] [roleId]",
	Short: "Delete a grant of a role in a division",
	Long:  "Delete a grant of a role in a division",
	Args:  utils.DetermineArgs([]string{ "subjectId", "divisionId", "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}"
		subjectId, args := args[0], args[1:]
		path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
