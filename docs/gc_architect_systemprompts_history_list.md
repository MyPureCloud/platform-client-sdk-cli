## gc architect systemprompts history list

Get generated prompt history

### Synopsis

Get generated prompt history

```
gc architect systemprompts history list [promptId] [historyId] [flags]
```

### Options

```
      --action strings           Flow actions to include (omit to include all) Valid values: checkin, checkout, create, deactivate, debug, delete, publish, revert, save
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --sortBy string            Sort by Valid values: action, timestamp, user
      --sortOrder string         Sort order
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

* [gc architect systemprompts history](gc_architect_systemprompts_history.html)	 - /api/v2/architect/systemprompts/{promptId}/history


