## gc analytics reporting dashboards users list

Get dashboards summary for users in a org

### Synopsis

Get dashboards summary for users in a org

```
gc analytics reporting dashboards users list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --deletedOnly string       Only list deleted dashboards that are still recoverable Valid values: true, false
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               A list of user IDs to fetch by bulk
      --pageNumber string         (default "1")
      --pageSize string           (default "25")
      --sortBy string            
      --state string             Only list users of this state Valid values: active, inactive
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

* [gc analytics reporting dashboards users](gc_analytics_reporting_dashboards_users.html)	 - /api/v2/analytics/reporting/dashboards/users


