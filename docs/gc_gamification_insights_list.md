## gc gamification insights list

Get insights summary

### Synopsis

Get insights summary

```
gc gamification insights list [flags]
```

### Options

```
  -a, --autopaginate                           Automatically paginate through the results stripping page information
      --comparativePeriodStartWorkday string   The start work day of comparative period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --filterId string                        ID for the filter type. - REQUIRED
      --filterType string                      Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division
      --filtercondition string                 Filter list command output based on a given condition or regular expression
      --granularity string                     Granularity - REQUIRED Valid values: Weekly, Monthly
  -h, --help                                   help for list
      --pageNumber string                      Page number (default "1")
      --pageSize string                        Page size (default "25")
      --primaryPeriodStartWorkday string       The start work day of primary period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --reportsTo string                       The reportsTo used by ABAC policies.
      --sortKey string                         Sort key Valid values: percentOfGoal, percentOfGoalChange, overallPercentOfGoal, overallPercentOfGoalChange, value, valueChange
      --sortMetricId string                    Sort Metric Id
      --sortOrder string                       Sort order Valid values: asc, desc
  -s, --stream                                 Paginate and stream data as it is being processed leaving page information intact
      --userIds string                         A list of up to 100 comma-separated user Ids
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

* [gc gamification insights](gc_gamification_insights.html)	 - /api/v2/gamification/insights


