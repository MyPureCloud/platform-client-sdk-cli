## gc gamification insights rankings get

Get insights rankings

### Synopsis

Get insights rankings

```
gc gamification insights rankings get [flags]
```

### Options

```
      --comparativePeriodStartWorkday string   The start work day of comparative period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --filterId string                        ID for the filter type. - REQUIRED
      --filterType string                      Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division
      --granularity string                     Granularity - REQUIRED Valid values: Weekly, Monthly
  -h, --help                                   help for get
      --primaryPeriodStartWorkday string       The start work day of primary period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --reportsTo string                       The reportsTo used by ABAC policies.
      --sectionSize string                     The number of top and bottom users to return before ties
      --sortKey string                         Sort key - REQUIRED Valid values: percentOfGoal, percentOfGoalChange, overallPercentOfGoal, overallPercentOfGoalChange, value, valueChange
      --sortMetricId string                    Sort Metric Id
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

* [gc gamification insights rankings](gc_gamification_insights_rankings.html)	 - /api/v2/gamification/insights/rankings


