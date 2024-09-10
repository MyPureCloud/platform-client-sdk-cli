## gc conversations calls history list

Get call history

### Synopsis

Get call history

```
gc conversations calls history list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Which fields, if any, to expand. Valid values: externalorganization, externalcontact, user, queue, group
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --interval /               Interval string; format is ISO-8601. Separate start and end times with forward slash /
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size, maximum 50 (default 25)
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

* [gc conversations calls history](gc_conversations_calls_history.html)	 - /api/v2/conversations/calls/history


