## gc outbound digitalrulesets list

Query a list of Outbound Digital Rule Sets

### Synopsis

Query a list of Outbound Digital Rule Sets

```
gc outbound digitalrulesets list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               A list of digital rule set ids to bulk fetch
      --name string              Name
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size. The max that will be returned is 100. (default 25)
      --sortBy string            The field to sort by Valid values: name
      --sortOrder string         The direction to sort Valid values: ascending, descending
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc outbound digitalrulesets](gc_outbound_digitalrulesets.html)	 - /api/v2/outbound/digitalrulesets


