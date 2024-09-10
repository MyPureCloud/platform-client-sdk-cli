## gc analytics reporting settings dashboards query list

Get list of dashboard configurations

### Synopsis

Get list of dashboard configurations

```
gc analytics reporting settings dashboards query list [flags]
```

### Options

```
  -a, --autopaginate                   Automatically paginate through the results stripping page information
      --dashboardAccessFilter string   Filter dashboard based on the owner of dashboard - REQUIRED Valid values: OwnedByMe, OwnedByAnyone, NotOwnedByMe
      --dashboardType string           List dashboard of given type - REQUIRED Valid values: All, Public, Private, Shared, Favorites
      --filtercondition string         Filter list command output based on a given condition or regular expression
  -h, --help                           help for list
      --name string                    name of the dashboard
      --pageNumber int                  (default 1)
      --pageSize int                    (default 9)
      --sortBy string                  
  -s, --stream                         Paginate and stream data as it is being processed leaving page information intact
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

* [gc analytics reporting settings dashboards query](gc_analytics_reporting_settings_dashboards_query.html)	 - /api/v2/analytics/reporting/settings/dashboards/query


