## gc flows list

Get a pageable list of flows, filtered by query parameters

### Synopsis

Get a pageable list of flows, filtered by query parameters

```
gc flows list [flags]
```

### Options

```
  -a, --autopaginate                 Automatically paginate through the results stripping page information
      --deleted string               Include deleted Valid values: true, false
      --description string           Description
      --divisionId strings           division ID(s)
      --editableBy string            Editable by
      --filtercondition string       Filter list command output based on a given condition or regular expression
  -h, --help                         help for list
      --id strings                   ID
      --includeSchemas string        Include variable schemas Valid values: true, false
      --lockedBy string              Locked by
      --lockedByClientId string      Locked by client ID
      --name string                  Name
      --nameOrDescription string     Name or description
      --pageNumber string            Page number (default "1")
      --pageSize string              Page size (default "25")
      --publishVersionId string      Publish version ID
      --publishedAfter string        Published after
      --publishedBefore string       Published before
      --secure string                Secure Valid values: any, checkedin, published
      --sortBy string                Sort by
      --sortOrder string             Sort order
  -s, --stream                       Paginate and stream data as it is being processed leaving page information intact
      --varType strings              Type Valid values: bot, commonmodule, digitalbot, inboundcall, inboundchat, inboundemail, inboundshortmessage, outboundcall, inqueuecall, inqueueemail, inqueueshortmessage, speech, securecall, surveyinvite, voice, voicemail, voicesurvey, workflow, workitem
      --virtualAgentEnabled string   Include/exclude virtual agent flows Valid values: true, false
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

* [gc flows](gc_flows.html)	 - /api/v2/flows


