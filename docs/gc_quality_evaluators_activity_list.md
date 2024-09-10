## gc quality evaluators activity list

Get an evaluator activity

### Synopsis

Get an evaluator activity

```
gc quality evaluators activity list [flags]
```

### Options

```
      --agentTeamId string       team id of agents to be considered
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --endTime string           The end time specified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --expand strings           variable name requested by expand list
      --filtercondition string   Filter list command output based on a given condition or regular expression
      --group string             group id
  -h, --help                     help for list
      --name string              Evaluator name
      --nextPage string          next page token
      --pageNumber int           The page number requested (default 1)
      --pageSize int             The total page size requested (default 25)
      --permission strings       permission strings
      --previousPage string      Previous page token
      --sortBy string            variable name requested to sort by
      --startTime string         The start time specified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

* [gc quality evaluators activity](gc_quality_evaluators_activity.html)	 - /api/v2/quality/evaluators/activity


