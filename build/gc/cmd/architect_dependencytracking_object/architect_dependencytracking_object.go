package architect_dependencytracking_object

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
	Description = utils.FormatUsageDescription("architect_dependencytracking_object", "SWAGGER_OVERRIDE_/api/v2/architect/dependencytracking/object", )
	architect_dependencytracking_objectCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_dependencytracking_object"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_dependencytracking_objectCmd)
}

func Cmdarchitect_dependencytracking_object() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "id", "", "Object ID - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "version", "", "Object version")
	utils.AddFlag(getCmd.Flags(), "string", "objectType", "", "Object type Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, DATAACTION, DATATABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GROUP, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, VOICEMAILFLOW, WIDGET, WORKFLOW, WORKITEMFLOW")
	utils.AddFlag(getCmd.Flags(), "bool", "consumedResources", "", "Include resources this item consumes")
	utils.AddFlag(getCmd.Flags(), "bool", "consumingResources", "", "Include resources that consume this item")
	utils.AddFlag(getCmd.Flags(), "[]string", "consumedResourceType", "", "Types of consumed resources to return, if consumed resources are requested Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, DATAACTION, DATATABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GROUP, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, VOICEMAILFLOW, WIDGET, WORKFLOW, WORKITEMFLOW")
	utils.AddFlag(getCmd.Flags(), "[]string", "consumingResourceType", "", "Types of consuming resources to return, if consuming resources are requested Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, DATAACTION, DATATABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GROUP, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, VOICEMAILFLOW, WIDGET, WORKFLOW, WORKITEMFLOW")
	utils.AddFlag(getCmd.Flags(), "bool", "consumedResourceRequest", "", "Indicate that this is going to look up a consumed resource object")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/architect/dependencytracking/object", utils.FormatPermissions([]string{ "architect:dependencyTracking:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/dependencytracking/object")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("id")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DependencyObject"
  }
}`)
	architect_dependencytracking_objectCmd.AddCommand(getCmd)
	
	return architect_dependencytracking_objectCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Dependency Tracking object",
	Long:  "Get a Dependency Tracking object",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/dependencytracking/object"


		id := utils.GetFlag(cmd.Flags(), "string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		version := utils.GetFlag(cmd.Flags(), "string", "version")
		if version != "" {
			queryParams["version"] = version
		}
		objectType := utils.GetFlag(cmd.Flags(), "string", "objectType")
		if objectType != "" {
			queryParams["objectType"] = objectType
		}
		consumedResources := utils.GetFlag(cmd.Flags(), "bool", "consumedResources")
		if consumedResources != "" {
			queryParams["consumedResources"] = consumedResources
		}
		consumingResources := utils.GetFlag(cmd.Flags(), "bool", "consumingResources")
		if consumingResources != "" {
			queryParams["consumingResources"] = consumingResources
		}
		consumedResourceType := utils.GetFlag(cmd.Flags(), "[]string", "consumedResourceType")
		if consumedResourceType != "" {
			queryParams["consumedResourceType"] = consumedResourceType
		}
		consumingResourceType := utils.GetFlag(cmd.Flags(), "[]string", "consumingResourceType")
		if consumingResourceType != "" {
			queryParams["consumingResourceType"] = consumingResourceType
		}
		consumedResourceRequest := utils.GetFlag(cmd.Flags(), "bool", "consumedResourceRequest")
		if consumedResourceRequest != "" {
			queryParams["consumedResourceRequest"] = consumedResourceRequest
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
