## gc architect dependencytracking deletedresourceconsumers list

Get Dependency Tracking objects that consume deleted resources

### Synopsis

Get Dependency Tracking objects that consume deleted resources

```
gc architect dependencytracking deletedresourceconsumers list [flags]
```

### Options

```
  -a, --autopaginate                   Automatically paginate through the results stripping page information
      --consumedResourceType strings   Resource type(s) to return Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, AGENTICVIRTUALAGENT, AGENTICVIRTUALAGENTVERSION, AUDIOCONNECTORBOT, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, CONVERSATIONCUSTOMATTRIBUTESCHEMA, DATAACTION, DATATABLE, DECISIONTABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, DIGITALBOTCONNECTOR, DIGITALBOTCONNECTORINTEGRATION, DIGITALBOTFLOW, DIVISION, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GRAMMAR, GROUP, GUIDE, GUIDEVERSION, IMAGE, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, KNOWLEDGEBASEDOCUMENT, KNOWLEDGESETTING, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OAUTHCLIENT, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SMSPHONENUMBER, STTENGINE, SURVEYFORM, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, UTILIZATIONLABEL, VOICEFLOW, VOICEMAILFLOW, VOICESURVEYFLOW, WIDGET, WORKFLOW, WORKITEMFLOW, WORKTYPE
      --consumedResources string       Return consumed resources? Valid values: true, false
      --filtercondition string         Filter list command output based on a given condition or regular expression
      --flowFilter string              Show only checkedIn or published flows Valid values: checkedIn, published
  -h, --help                           help for list
      --name string                    Name to search for
      --objectType strings             Object type(s) to search for Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, AGENTICVIRTUALAGENT, AGENTICVIRTUALAGENTVERSION, AUDIOCONNECTORBOT, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, CONVERSATIONCUSTOMATTRIBUTESCHEMA, DATAACTION, DATATABLE, DECISIONTABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, DIGITALBOTCONNECTOR, DIGITALBOTCONNECTORINTEGRATION, DIGITALBOTFLOW, DIVISION, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GRAMMAR, GROUP, GUIDE, GUIDEVERSION, IMAGE, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, KNOWLEDGEBASEDOCUMENT, KNOWLEDGESETTING, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OAUTHCLIENT, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SMSPHONENUMBER, STTENGINE, SURVEYFORM, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, UTILIZATIONLABEL, VOICEFLOW, VOICEMAILFLOW, VOICESURVEYFLOW, WIDGET, WORKFLOW, WORKITEMFLOW, WORKTYPE
      --pageNumber string              Page number (default "1")
      --pageSize string                Page size (default "25")
  -s, --stream                         Paginate and stream data as it is being processed leaving page information intact
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc architect dependencytracking deletedresourceconsumers](gc_architect_dependencytracking_deletedresourceconsumers.html)	 - /api/v2/architect/dependencytracking/deletedresourceconsumers


