## gc outbound campaigns list

Query a list of dialer campaigns.

### Synopsis

Query a list of dialer campaigns.

```
gc outbound campaigns list [flags]
```

### Options

```
  -a, --autopaginate                       Automatically paginate through the results stripping page information
      --callAnalysisResponseSetId string   Call analysis response set ID
      --contactListId string               Contact List ID
      --distributionQueueId string         Distribution queue ID
      --divisionId strings                 Division ID(s)
      --dncListIds string                  DNC list ID
      --edgeGroupId string                 Edge group ID
      --filterType string                  Filter type Valid values: Equals, RegEx, Contains, Prefix, LessThan, LessThanEqualTo, GreaterThan, GreaterThanEqualTo, BeginsWith, EndsWith
      --filtercondition string             Filter list command output based on a given condition or regular expression
  -h, --help                               help for list
      --id strings                         id
      --name string                        Name
      --pageNumber string                  Page number (default "1")
      --pageSize string                    Page size. The max that will be returned is 100. (default "25")
      --sortBy string                      Sort by
      --sortOrder string                   Sort order Valid values: ascending, descending
  -s, --stream                             Paginate and stream data as it is being processed leaving page information intact
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

* [gc outbound campaigns](gc_outbound_campaigns.html)	 - /api/v2/outbound/campaigns

