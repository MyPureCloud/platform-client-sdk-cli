## gc gamification scorecards users points trends get

Points trend for a user

### Synopsis

Points trend for a user

```
gc gamification scorecards users points trends get [userId] [flags]
```

### Options

```
      --dayOfWeek string      Optional filter to specify which day of weeks to be included in the response Valid values: Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday
      --endWorkday string     End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
  -h, --help                  help for get
      --startWorkday string   Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
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

* [gc gamification scorecards users points trends](gc_gamification_scorecards_users_points_trends.html)	 - /api/v2/gamification/scorecards/users/{userId}/points/trends


