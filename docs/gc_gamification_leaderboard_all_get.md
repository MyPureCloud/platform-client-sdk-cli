## gc gamification leaderboard all get

Leaderboard by filter type

### Synopsis

Leaderboard by filter type

```
gc gamification leaderboard all get [flags]
```

### Options

```
      --endWorkday string     End workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --filterId string       ID for the filter type. For example, division or performance profile Id - REQUIRED
      --filterType string     Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division
  -h, --help                  help for get
      --metricId string       Metric Id for which the leaderboard is to be generated. The total points is used if nothing is given.
      --startWorkday string   Start workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
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

* [gc gamification leaderboard all](gc_gamification_leaderboard_all.html)	 - /api/v2/gamification/leaderboard/all


