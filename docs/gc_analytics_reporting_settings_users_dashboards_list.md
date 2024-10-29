## gc analytics reporting settings users dashboards list

Get list of dashboards for an user

### Synopsis

Get list of dashboards for an user

```
gc analytics reporting settings users dashboards list [userId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --deletedOnly string       If true, retrieve only deleted dashboards that are still recoverable Valid values: true, false
      --favoriteOnly string      If true, retrieve only favorite dashboards Valid values: true, false
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              retrieve dashboards that match with given name
      --pageNumber int            (default 1)
      --pageSize int              (default 50)
      --publicOnly string        If true, retrieve only public dashboards Valid values: true, false
      --sortBy string            
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

* [gc analytics reporting settings users dashboards](gc_analytics_reporting_settings_users_dashboards.html)	 - /api/v2/analytics/reporting/settings/users/{userId}/dashboards


