## gc workforcemanagement businessunits shifttrading trades query jobs create

Queries shift trades in a management unit per user

### Synopsis

Queries shift trades in a management unit per user

```
gc workforcemanagement businessunits shifttrading trades query jobs create [businessUnitId] [flags]
```

### Options

```
  -d, --directory string              Directory path with files containing request bodies
  -f, --file string                   File name containing the JSON body
      --forceAsync string             Force the result of this operation to be sent asynchronously via notification. For testing/app development purposes Valid values: true, false
      --forceDownloadService string   Force the result of this operation to be sent via download service. For testing/app development purposes Valid values: true, false
  -h, --help                          help for create
  -b, --printrequestbody              Print the request body format of the API.
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

* [gc workforcemanagement businessunits shifttrading trades query jobs](gc_workforcemanagement_businessunits_shifttrading_trades_query_jobs.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/trades/query/jobs


