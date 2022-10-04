package knowledge_guest_sessions_categories

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
	Description = utils.FormatUsageDescription("knowledge_guest_sessions_categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/guest/sessions/{sessionId}/categories", )
	knowledge_guest_sessions_categoriesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_guest_sessions_categories"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_guest_sessions_categoriesCmd)
}

func Cmdknowledge_guest_sessions_categories() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cursor that points to the start of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 200.")
	utils.AddFlag(listCmd.Flags(), "string", "parentId", "", "If specified, retrieves the children categories by parent category ID.")
	utils.AddFlag(listCmd.Flags(), "bool", "isRoot", "", "If specified, retrieves only the root categories.")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter to return the categories that starts with the given category name.")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "Name", "Name: sort by category names alphabetically; Hierarchy: sort by the full path of hierarchical category names alphabetically Valid values: Name, Hierarchy")
	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "The specified entity attribute will be filled. Supported value:Ancestors: every ancestors will be filled via the parent attribute recursively,but only the id, name, parentId will be present for the ancestors.")
	utils.AddFlag(listCmd.Flags(), "bool", "includeDocumentCount", "", "If specified, retrieves the number of documents related to category.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/knowledge/guest/sessions/{sessionId}/categories", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Knowledge", "/api/v2/knowledge/guest/sessions/{sessionId}/categories")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	knowledge_guest_sessions_categoriesCmd.AddCommand(listCmd)
	return knowledge_guest_sessions_categoriesCmd
}

var listCmd = &cobra.Command{
	Use:   "list [sessionId]",
	Short: "Get categories",
	Long:  "Get categories",
	Args:  utils.DetermineArgs([]string{ "sessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/guest/sessions/{sessionId}/categories"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)

		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		parentId := utils.GetFlag(cmd.Flags(), "string", "parentId")
		if parentId != "" {
			queryParams["parentId"] = parentId
		}
		isRoot := utils.GetFlag(cmd.Flags(), "bool", "isRoot")
		if isRoot != "" {
			queryParams["isRoot"] = isRoot
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includeDocumentCount := utils.GetFlag(cmd.Flags(), "bool", "includeDocumentCount")
		if includeDocumentCount != "" {
			queryParams["includeDocumentCount"] = includeDocumentCount
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
