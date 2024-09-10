## gc workforcemanagement managementunits workplans get

Get a work plan

### Synopsis

Get a work plan

```
gc workforcemanagement managementunits workplans get [managementUnitId] [workPlanId] [flags]
```

### Options

```
  -h, --help                  help for get
      --includeOnly strings   limit response to the specified fields Valid values: agentCount, agents, optionalDays, shifts, shiftStartVariances
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

* [gc workforcemanagement managementunits workplans](gc_workforcemanagement_managementunits_workplans.html)	 - /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans


