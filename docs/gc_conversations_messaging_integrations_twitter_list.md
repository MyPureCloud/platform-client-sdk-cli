## gc conversations messaging integrations twitter list

Get a list of Twitter Integrations

### Synopsis

Get a list of Twitter Integrations

```
gc conversations messaging integrations twitter list [flags]
```

### Options

```
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --expand string               Expand instructions for the return value. Valid values: supportedContent, messagingSetting, identityresolution
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --messagingSettingId string   Filter integrations returned based on the setting ID
      --pageNumber string           Page number (default "1")
      --pageSize string             Page size (default "25")
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
      --supportedContentId string   Filter integrations returned based on the supported content ID
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

* [gc conversations messaging integrations twitter](gc_conversations_messaging_integrations_twitter.html)	 - /api/v2/conversations/messaging/integrations/twitter


