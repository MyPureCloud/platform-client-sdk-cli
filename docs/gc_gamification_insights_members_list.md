## gc gamification insights members list

Query users in a profile during a period of time

### Synopsis

Query users in a profile during a period of time

```
gc gamification insights members list [flags]
```

### Options

```
      --filterId string       ID for the filter type. - REQUIRED
      --filterType string     Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division
      --granularity string    Granularity - REQUIRED Valid values: Weekly, Monthly
  -h, --help                  help for list
      --reportsTo string      The reportsTo used by ABAC policies.
      --startWorkday string   The start work day. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
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

* [gc gamification insights members](gc_gamification_insights_members.html)	 - /api/v2/gamification/insights/members


