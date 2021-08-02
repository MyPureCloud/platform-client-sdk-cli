package architect_dependencytracking_consumingresources

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
	Description = utils.FormatUsageDescription("architect_dependencytracking_consumingresources", "SWAGGER_OVERRIDE_/api/v2/architect/dependencytracking/consumingresources", )
	architect_dependencytracking_consumingresourcesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_dependencytracking_consumingresources"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_dependencytracking_consumingresourcesCmd)
}

func Cmdarchitect_dependencytracking_consumingresources() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "id", "", "Consumed object ID - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "objectType", "", "Consumed object type - REQUIRED Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, DATAACTION, DATATABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GROUP, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, VOICEMAILFLOW, WIDGET, WORKFLOW, WORKITEMFLOW")
	utils.AddFlag(listCmd.Flags(), "[]string", "resourceType", "", "Types of consuming resources to show.  Only versioned types are allowed here. Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, DATAACTION, DATATABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GROUP, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, VOICEMAILFLOW, WIDGET, WORKFLOW, WORKITEMFLOW")
	utils.AddFlag(listCmd.Flags(), "string", "version", "", "Object version")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "flowFilter", "", "Show only checkedIn or published flows Valid values: checkedIn, published")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/architect/dependencytracking/consumingresources", utils.FormatPermissions([]string{ "architect:dependencyTracking:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/dependencytracking/consumingresources")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("id")
	listCmd.MarkFlagRequired("objectType")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	architect_dependencytracking_consumingresourcesCmd.AddCommand(listCmd)
	
	return architect_dependencytracking_consumingresourcesCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get resources that consume a given Dependency Tracking object",
	Long:  "Get resources that consume a given Dependency Tracking object",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/architect/dependencytracking/consumingresources"

		id := utils.GetFlag(cmd.Flags(), "string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		objectType := utils.GetFlag(cmd.Flags(), "string", "objectType")
		if objectType != "" {
			queryParams["objectType"] = objectType
		}
		resourceType := utils.GetFlag(cmd.Flags(), "[]string", "resourceType")
		if resourceType != "" {
			queryParams["resourceType"] = resourceType
		}
		version := utils.GetFlag(cmd.Flags(), "string", "version")
		if version != "" {
			queryParams["version"] = version
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		flowFilter := utils.GetFlag(cmd.Flags(), "string", "flowFilter")
		if flowFilter != "" {
			queryParams["flowFilter"] = flowFilter
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
