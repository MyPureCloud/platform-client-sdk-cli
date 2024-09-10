## gc workforcemanagement managementunits workplans list

Get work plans

### Synopsis

Get work plans

```
gc workforcemanagement managementunits workplans list [managementUnitId] [flags]
```

### Options

```
      --exclude strings   Exclude specific data on the work plans from the response Valid values: shifts.activities
      --expand strings    Include to access additional data on the work plans Valid values: agentCount, agents, optionalDays, shifts, shiftStartVariances, details
  -h, --help              help for list
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


