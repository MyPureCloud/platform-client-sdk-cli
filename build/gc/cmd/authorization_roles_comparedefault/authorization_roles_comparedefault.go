package authorization_roles_comparedefault

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
	Description = utils.FormatUsageDescription("authorization_roles_comparedefault", "SWAGGER_OVERRIDE_/api/v2/authorization/roles/{leftRoleId}/comparedefault", "SWAGGER_OVERRIDE_/api/v2/authorization/roles/{leftRoleId}/comparedefault", )
	authorization_roles_comparedefaultCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_roles_comparedefault"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_roles_comparedefaultCmd)
}

func Cmdauthorization_roles_comparedefault() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}", utils.FormatPermissions([]string{ "authorization:role:view",  }), utils.GenerateDevCentreLink("POST", "Authorization", "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Organization role",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DomainOrganizationRole"
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
        "$ref" : "#/components/schemas/DomainOrgRoleDifference"
      }
    }
  }
}`)
	authorization_roles_comparedefaultCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}", utils.FormatPermissions([]string{ "authorization:role:view",  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DomainOrgRoleDifference"
      }
    }
  }
}`)
	authorization_roles_comparedefaultCmd.AddCommand(getCmd)
	return authorization_roles_comparedefaultCmd
}

var createCmd = &cobra.Command{
	Use:   "create [leftRoleId] [rightRoleId]",
	Short: "Get an unsaved org role to default role comparison",
	Long:  "Get an unsaved org role to default role comparison",
	Args:  utils.DetermineArgs([]string{ "leftRoleId", "rightRoleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Domainorganizationrole{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}"
		leftRoleId, args := args[0], args[1:]
		path = strings.Replace(path, "{leftRoleId}", fmt.Sprintf("%v", leftRoleId), -1)
		rightRoleId, args := args[0], args[1:]
		path = strings.Replace(path, "{rightRoleId}", fmt.Sprintf("%v", rightRoleId), -1)

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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [leftRoleId] [rightRoleId]",
	Short: "Get an org role to default role comparison",
	Long:  "Get an org role to default role comparison",
	Args:  utils.DetermineArgs([]string{ "leftRoleId", "rightRoleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}"
		leftRoleId, args := args[0], args[1:]
		path = strings.Replace(path, "{leftRoleId}", fmt.Sprintf("%v", leftRoleId), -1)
		rightRoleId, args := args[0], args[1:]
		path = strings.Replace(path, "{rightRoleId}", fmt.Sprintf("%v", rightRoleId), -1)

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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
