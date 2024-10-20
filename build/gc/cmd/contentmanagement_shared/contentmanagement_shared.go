package contentmanagement_shared

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
	Description = utils.FormatUsageDescription("contentmanagement_shared", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/shared", )
	contentmanagement_sharedCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_shared"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_sharedCmd)
}

func Cmdcontentmanagement_shared() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "disposition", "attachment", "Request how the share content will be downloaded: attached as a file or inline. Default is attachment. Valid values: attachment, inline, none")
	utils.AddFlag(getCmd.Flags(), "string", "contentType", "", "The requested format for the specified document. If supported, the document will be returned in that format. Example contentType=audio/wav")
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Expand some document fields Valid values: document.acl")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/shared/{sharedId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/shared/{sharedId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Download location is returned in header, if redirect is set to false and disposition is not set to none. If disposition is none, location header will not be populated, DownloadUri and ViewUri will be populated.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SharedResponse"
      }
    }
  }
}`)
	contentmanagement_sharedCmd.AddCommand(getCmd)
	return contentmanagement_sharedCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [sharedId]",
	Short: "Get shared documents. Securely download a shared document.",
	Long:  "Get shared documents. Securely download a shared document.",
	Args:  utils.DetermineArgs([]string{ "sharedId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/shared/{sharedId}"
		sharedId, args := args[0], args[1:]
		path = strings.Replace(path, "{sharedId}", fmt.Sprintf("%v", sharedId), -1)

		disposition := utils.GetFlag(cmd.Flags(), "string", "disposition")
		if disposition != "" {
			queryParams["disposition"] = disposition
		}
		contentType := utils.GetFlag(cmd.Flags(), "string", "contentType")
		if contentType != "" {
			queryParams["contentType"] = contentType
		}
		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
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
