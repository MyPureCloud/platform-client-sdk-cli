## gc outbound messagingcampaigns divisionviews list

Query a list of basic Messaging Campaign information objects

### Synopsis

Query a list of basic Messaging Campaign information objects

```
gc outbound messagingcampaigns divisionviews list [flags]
```

### Options

```
  -a, --autopaginate                  Automatically paginate through the results stripping page information
      --filtercondition string        Filter list command output based on a given condition or regular expression
  -h, --help                          help for list
      --id strings                    id
      --name string                   Name
      --pageNumber string             Page number (default "1")
      --pageSize string               Page size. The max that will be returned is 100. (default "25")
      --senderSmsPhoneNumber string   Sender SMS Phone Number
      --sortOrder string              The direction to sort Valid values: ascending, descending
  -s, --stream                        Paginate and stream data as it is being processed leaving page information intact
      --varType string                Campaign Type Valid values: EMAIL, SMS
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

* [gc outbound messagingcampaigns divisionviews](gc_outbound_messagingcampaigns_divisionviews.html)	 - /api/v2/outbound/messagingcampaigns/divisionviews

