package contentmanagement_documents_audits

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
	Description = utils.FormatUsageDescription("contentmanagement_documents_audits", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents/{documentId}/audits", )
	contentmanagement_documents_auditsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_documents_audits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_documents_auditsCmd)
}

func Cmdcontentmanagement_documents_audits() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "transactionFilter", "", "Transaction filter")
	utils.AddFlag(listCmd.Flags(), "string", "level", "USER", "level")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "Sort order")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/documents/{documentId}/audits", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	contentmanagement_documents_auditsCmd.AddCommand(listCmd)
	
	return contentmanagement_documents_auditsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [documentId]",
	Short: "Get a list of audits for a document.",
	Long:  "Get a list of audits for a document.",
	Args:  utils.DetermineArgs([]string{ "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents/{documentId}/audits"
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		transactionFilter := utils.GetFlag(cmd.Flags(), "string", "transactionFilter")
		if transactionFilter != "" {
			queryParams["transactionFilter"] = transactionFilter
		}
		level := utils.GetFlag(cmd.Flags(), "string", "level")
		if level != "" {
			queryParams["level"] = level
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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
