package knowledge_knowledgebases_languages_categories

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
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_languages_categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", )
	knowledge_knowledgebases_languages_categoriesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_languages_categories"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_languages_categoriesCmd)
}

func Cmdknowledge_knowledgebases_languages_categories() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", utils.FormatPermissions([]string{ "knowledge:category:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeCategoryRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeExtendedCategory&quot;
  }
}`)
	knowledge_knowledgebases_languages_categoriesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}", utils.FormatPermissions([]string{ "knowledge:category:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeCategory&quot;
  }
}`)
	knowledge_knowledgebases_languages_categoriesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}", utils.FormatPermissions([]string{ "knowledge:category:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeExtendedCategory&quot;
  }
}`)
	knowledge_knowledgebases_languages_categoriesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cursor that points to the start of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "limit", "", "Number of entities to return. Maximum of 200. Deprecated in favour of pageSize.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 200.")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter to return the categories that starts with the given category name.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories", utils.FormatPermissions([]string{ "knowledge:category:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	knowledge_knowledgebases_languages_categoriesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}", utils.FormatPermissions([]string{ "knowledge:category:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeCategoryRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeExtendedCategory&quot;
  }
}`)
	knowledge_knowledgebases_languages_categoriesCmd.AddCommand(updateCmd)
	
	return knowledge_knowledgebases_languages_categoriesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [knowledgeBaseId] [languageCode]",
	Short: "Create new category",
	Long:  "Create new category",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

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
	Use:   "delete [categoryId] [knowledgeBaseId] [languageCode]",
	Short: "Delete category",
	Long:  "Delete category",
	Args:  utils.DetermineArgs([]string{ "categoryId", "knowledgeBaseId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
		categoryId, args := args[0], args[1:]
		path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

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
var getCmd = &cobra.Command{
	Use:   "get [categoryId] [knowledgeBaseId] [languageCode]",
	Short: "Get category",
	Long:  "Get category",
	Args:  utils.DetermineArgs([]string{ "categoryId", "knowledgeBaseId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
		categoryId, args := args[0], args[1:]
		path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [knowledgeBaseId] [languageCode]",
	Short: "Get categories",
	Long:  "Get categories",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		limit := utils.GetFlag(cmd.Flags(), "string", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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
var updateCmd = &cobra.Command{
	Use:   "update [categoryId] [knowledgeBaseId] [languageCode]",
	Short: "Update category",
	Long:  "Update category",
	Args:  utils.DetermineArgs([]string{ "categoryId", "knowledgeBaseId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
		categoryId, args := args[0], args[1:]
		path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
