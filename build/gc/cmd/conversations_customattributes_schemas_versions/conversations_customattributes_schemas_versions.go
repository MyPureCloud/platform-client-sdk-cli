package conversations_customattributes_schemas_versions

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
	Description = utils.FormatUsageDescription("conversations_customattributes_schemas_versions", "SWAGGER_OVERRIDE_/api/v2/conversations/customattributes/schemas/{schemaId}/versions", "SWAGGER_OVERRIDE_/api/v2/conversations/customattributes/schemas/{schemaId}/versions", )
	conversations_customattributes_schemas_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_customattributes_schemas_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_customattributes_schemas_versionsCmd)
}

func Cmdconversations_customattributes_schemas_versions() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/customattributes/schemas/{schemaId}/versions/{versionId}", utils.FormatPermissions([]string{ "conversation:customAttributes:schemaView",  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/customattributes/schemas/{schemaId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ConversationDataSchema"
      }
    }
  }
}`)
	conversations_customattributes_schemas_versionsCmd.AddCommand(getCmd)

	getAllConversationsCustomattributesSchemaVersionsCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getAllConversationsCustomattributesSchemaVersionsCmd.UsageTemplate(), "GET", "/api/v2/conversations/customattributes/schemas/{schemaId}/versions", utils.FormatPermissions([]string{ "conversation:customAttributes:schemaView",  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/customattributes/schemas/{schemaId}/versions")))
	utils.AddFileFlagIfUpsert(getAllConversationsCustomattributesSchemaVersionsCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getAllConversationsCustomattributesSchemaVersionsCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ConversationDataSchemaListing"
      }
    }
  }
}`)
	conversations_customattributes_schemas_versionsCmd.AddCommand(getAllConversationsCustomattributesSchemaVersionsCmd)
	return conversations_customattributes_schemas_versionsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [schemaId] [versionId]",
	Short: "Get a specific version of a schema",
	Long:  "Get a specific version of a schema",
	Args:  utils.DetermineArgs([]string{ "schemaId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/customattributes/schemas/{schemaId}/versions/{versionId}"
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
var getAllConversationsCustomattributesSchemaVersionsCmd = &cobra.Command{
	Use:   "getAllConversationsCustomattributesSchemaVersions [schemaId]",
	Short: "Get all versions of a CCA schema",
	Long:  "Get all versions of a CCA schema",
	Args:  utils.DetermineArgs([]string{ "schemaId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/customattributes/schemas/{schemaId}/versions"
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "getAllConversationsCustomattributesSchemaVersions"
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
