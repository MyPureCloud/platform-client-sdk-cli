## gc workforcemanagement managementunits weeks shifttrades list

Gets all the shift trades for a given week

### Synopsis

Gets all the shift trades for a given week

```
gc workforcemanagement managementunits weeks shifttrades list [managementUnitId] [weekDateId] [flags]
```

### Options

```
      --evaluateMatches string          Whether to evaluate the matches for violations Valid values: true, false
      --forceDownloadService string     Force the result of this operation to be sent via download service. For testing/app development purposes Valid values: true, false
  -h, --help                            help for list
      --includeCrossWeekShifts string   Whether to include all shift trades with either the initiating shift or the receiving shift in the week Valid values: true, false
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

* [gc workforcemanagement managementunits weeks shifttrades](gc_workforcemanagement_managementunits_weeks_shifttrades.html)	 - /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades


