## gc outbound messagingcampaigns list

Query a list of Messaging Campaigns

### Synopsis

Query a list of Messaging Campaigns

```
gc outbound messagingcampaigns list [flags]
```

### Options

```
  -a, --autopaginate                  Automatically paginate through the results stripping page information
      --campaignStatus string         Campaign Status Valid values: on, stopping, off, complete, invalid, forced_off, forced_stopping
      --contactListId string          Contact List ID
      --contentTemplateId string      Content template ID
      --divisionId strings            Division ID(s)
      --filtercondition string        Filter list command output based on a given condition or regular expression
  -h, --help                          help for list
      --id strings                    A list of messaging campaign ids to bulk fetch
      --name string                   Name
      --pageNumber string             Page number (default "1")
      --pageSize string               Page size. The max that will be returned is 100. (default "25")
      --ruleSetIds strings            Ruleset ID(s)
      --senderSmsPhoneNumber string   Sender SMS Phone Number
      --sortBy string                 The field to sort by Valid values: campaignStatus, name, type
      --sortOrder string              The direction to sort Valid values: ascending, descending
  -s, --stream                        Paginate and stream data as it is being processed leaving page information intact
      --varType string                Campaign Type Valid values: EMAIL, SMS, WHATSAPP
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

* [gc outbound messagingcampaigns](gc_outbound_messagingcampaigns.html)	 - /api/v2/outbound/messagingcampaigns


