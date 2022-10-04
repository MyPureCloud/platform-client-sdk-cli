package learning_rules_query

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
	Description = utils.FormatUsageDescription("learning_rules_query", "SWAGGER_OVERRIDE_/api/v2/learning/rules/query", )
	learning_rules_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_rules_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_rules_queryCmd)
}

func Cmdlearning_rules_query() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "int", "pageSize", "50", "Page size - REQUIRED")
	utils.AddFlag(createCmd.Flags(), "int", "pageNumber", "1", "Page number - REQUIRED")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/rules/query", utils.FormatPermissions([]string{ "learning:rule:view",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/rules/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The learning module rule to fetch users",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/LearningAssignmentUserQuery"
      }
    }
  },
  "required" : true
}`)
	createCmd.MarkFlagRequired("pageSize")
	createCmd.MarkFlagRequired("pageNumber")
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/LearningAssignmentUserListing"
      }
    }
  }
}`)
	learning_rules_queryCmd.AddCommand(createCmd)
	return learning_rules_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get users for learning module rule",
	Long:  "Get users for learning module rule",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Learningassignmentuserquery{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/rules/query"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
