## gc gamification leaderboard get

Leaderboard of the requesting user`s division or performance profile

### Synopsis

Leaderboard of the requesting user`s division or performance profile

```
gc gamification leaderboard get [flags]
```

### Options

```
      --endWorkday string     End workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
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

* [gc gamification leaderboard](gc_gamification_leaderboard.html)	 - /api/v2/gamification/leaderboard


