package flows_divisionviews

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
	Description = utils.FormatUsageDescription("flows_divisionviews", "SWAGGER_OVERRIDE_/api/v2/flows/divisionviews", )
	flows_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_divisionviewsCmd)
}

func Cmdflows_divisionviews() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "varType", "", "Type Valid values: bot, commonmodule, inboundcall, inboundchat, inboundemail, inboundshortmessage, outboundcall, inqueuecall, inqueueemail, inqueueshortmessage, speech, securecall, surveyinvite, voicemail, workflow, workitem")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "id", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "asc", "Sort order")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "ID")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "publishVersionId", "", "Publish version ID")
	utils.AddFlag(listCmd.Flags(), "string", "publishedAfter", "", "Published after")
	utils.AddFlag(listCmd.Flags(), "string", "publishedBefore", "", "Published before")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "division ID(s)")
	utils.AddFlag(listCmd.Flags(), "bool", "includeSchemas", "false", "Include variable schemas")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/divisionviews", utils.FormatPermissions([]string{ "architect:flow:search",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/divisionviews")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	flows_divisionviewsCmd.AddCommand(listCmd)
	
	return flows_divisionviewsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a pageable list of basic flow information objects filterable by query parameters.",
	Long:  "Get a pageable list of basic flow information objects filterable by query parameters.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/divisionviews"


		varType := utils.GetFlag(cmd.Flags(), "[]string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		publishVersionId := utils.GetFlag(cmd.Flags(), "string", "publishVersionId")
		if publishVersionId != "" {
			queryParams["publishVersionId"] = publishVersionId
		}
		publishedAfter := utils.GetFlag(cmd.Flags(), "string", "publishedAfter")
		if publishedAfter != "" {
			queryParams["publishedAfter"] = publishedAfter
		}
		publishedBefore := utils.GetFlag(cmd.Flags(), "string", "publishedBefore")
		if publishedBefore != "" {
			queryParams["publishedBefore"] = publishedBefore
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		includeSchemas := utils.GetFlag(cmd.Flags(), "bool", "includeSchemas")
		if includeSchemas != "" {
			queryParams["includeSchemas"] = includeSchemas
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
