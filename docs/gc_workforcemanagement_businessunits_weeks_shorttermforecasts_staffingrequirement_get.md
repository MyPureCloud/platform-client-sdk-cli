## gc workforcemanagement businessunits weeks shorttermforecasts staffingrequirement get

Get the staffing requirement by planning group for a forecast

### Synopsis

Get the staffing requirement by planning group for a forecast

```
gc workforcemanagement businessunits weeks shorttermforecasts staffingrequirement get [businessUnitId] [weekDateId] [forecastId] [flags]
```

### Options

```
      --expand strings        Expand to include minimum staffing values in (staffing requirement response or applied to base staffing requirement values) Valid values: results.planningGroupStaffingRequirements.minimumStaffPerInterval, results.planningGroupStaffingRequirements.effectiveStaffPerInterval
  -h, --help                  help for get
      --weekNumbers strings   The week numbers to fetch (for multi-week forecasts) staffing requirements. Returns all week data if the list is not specified
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

* [gc workforcemanagement businessunits weeks shorttermforecasts staffingrequirement](gc_workforcemanagement_businessunits_weeks_shorttermforecasts_staffingrequirement.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/staffingrequirement


