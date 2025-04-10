## gc gamification contests me list

Get a List of Contests (Agent/Supervisor)

### Synopsis

Get a List of Contests (Agent/Supervisor)

```
gc gamification contests me list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --dateEnd string           End date for the query. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
      --dateStart string         Start date for the query. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string         (default "1")
      --pageSize string           (default "25")
      --sortBy string             Valid values: title, dateStart, dateEnd, dateFinalized, status, profile, participantCount
      --sortOrder string          Valid values: asc, desc
      --status strings            Valid values: Upcoming, Ongoing, Pending, RecentlyCompleted, Completed, Cancelled
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --view string               Valid values: participant, creator
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

* [gc gamification contests me](gc_gamification_contests_me.html)	 - /api/v2/gamification/contests/me


