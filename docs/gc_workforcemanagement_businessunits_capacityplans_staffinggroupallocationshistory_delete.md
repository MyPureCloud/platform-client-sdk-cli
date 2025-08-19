## gc workforcemanagement businessunits capacityplans staffinggroupallocationshistory delete

Delete staffing group allocations history created for a capacity plan before the given date

### Synopsis

Delete staffing group allocations history created for a capacity plan before the given date

```
gc workforcemanagement businessunits capacityplans staffinggroupallocationshistory delete [businessUnitId] [capacityPlanId] [flags]
```

### Options

```
      --beforeDateId string   The date to delete records that are created on or before this date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
  -h, --help                  help for delete
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

* [gc workforcemanagement businessunits capacityplans staffinggroupallocationshistory](gc_workforcemanagement_businessunits_capacityplans_staffinggroupallocationshistory.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplans/{capacityPlanId}/staffinggroupallocationshistory


