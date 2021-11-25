package outbound_messagingcampaigns

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
	Description = utils.FormatUsageDescription("outbound_messagingcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns", )
	outbound_messagingcampaignsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_messagingcampaigns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_messagingcampaignsCmd)
}

func Cmdoutbound_messagingcampaigns() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/messagingcampaigns", utils.FormatPermissions([]string{ "outbound:messagingCampaign:add", "outbound:emailCampaign:add",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/messagingcampaigns")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Messaging Campaign",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	outbound_messagingcampaignsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}", utils.FormatPermissions([]string{ "outbound:messagingCampaign:delete", "outbound:emailCampaign:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	outbound_messagingcampaignsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}", utils.FormatPermissions([]string{ "outbound:messagingCampaign:view", "outbound:emailCampaign:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	outbound_messagingcampaignsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The max that will be returned is 100.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "The field to sort by Valid values: campaignStatus, name, type")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "The direction to sort Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "contactListId", "", "Contact List ID")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	utils.AddFlag(listCmd.Flags(), "string", "varType", "", "Campaign Type Valid values: EMAIL, SMS")
	utils.AddFlag(listCmd.Flags(), "string", "senderSmsPhoneNumber", "", "Sender SMS Phone Number")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "A list of messaging campaign ids to bulk fetch")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns", utils.FormatPermissions([]string{ "outbound:messagingCampaign:view", "outbound:emailCampaign:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	outbound_messagingcampaignsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}", utils.FormatPermissions([]string{ "outbound:messagingCampaign:edit", "outbound:emailCampaign:edit",  }), utils.GenerateDevCentreLink("PUT", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "MessagingCampaign",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/MessagingCampaign"
  }
}`)
	outbound_messagingcampaignsCmd.AddCommand(updateCmd)
	
	return outbound_messagingcampaignsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Messaging Campaign",
	Long:  "Create a Messaging Campaign",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Messagingcampaign{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns"

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
var deleteCmd = &cobra.Command{
	Use:   "delete [messagingCampaignId]",
	Short: "Delete an Outbound Messaging Campaign",
	Long:  "Delete an Outbound Messaging Campaign",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [messagingCampaignId]",
	Short: "Get an Outbound Messaging Campaign",
	Long:  "Get an Outbound Messaging Campaign",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Query a list of Messaging Campaigns",
	Long:  "Query a list of Messaging Campaigns",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		contactListId := utils.GetFlag(cmd.Flags(), "string", "contactListId")
		if contactListId != "" {
			queryParams["contactListId"] = contactListId
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
		senderSmsPhoneNumber := utils.GetFlag(cmd.Flags(), "string", "senderSmsPhoneNumber")
		if senderSmsPhoneNumber != "" {
			queryParams["senderSmsPhoneNumber"] = senderSmsPhoneNumber
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [messagingCampaignId]",
	Short: "Update an Outbound Messaging Campaign",
	Long:  "Update an Outbound Messaging Campaign",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Messagingcampaign{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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
