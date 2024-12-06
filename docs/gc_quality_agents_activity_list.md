## gc quality agents activity list

Gets a list of Agent Activities

### Synopsis

Gets a list of Agent Activities

```
gc quality agents activity list [flags]
```

### Options

```
      --agentTeamId string       team id of agents requested
      --agentUserId strings      user id of agent requested
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --endTime string           End time of agent activity based on assigned date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --evaluatorUserId string   user id of the evaluator
      --expand strings           variable name requested by expand list
      --filtercondition string   Filter list command output based on a given condition or regular expression
      --formContextId string     shared id between form versions
      --group string             group id
  -h, --help                     help for list
      --name string              name
      --nextPage string          next page token
      --pageNumber string        The page number requested (default "1")
      --pageSize string          The total page size requested (default "25")
      --previousPage string      Previous page token
      --sortBy string            variable name requested to sort by
      --startTime string         Start time of agent activity based on assigned date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --userState Legacy         Legacy fetches active and inactive users when evaluatorUserId or no user filters are supplied; otherwise fetches active users.  `Any` fetches users of `active`, `inactive` and `deleted` states. Valid values: Any, Legacy
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

* [gc quality agents activity](gc_quality_agents_activity.html)	 - /api/v2/quality/agents/activity


