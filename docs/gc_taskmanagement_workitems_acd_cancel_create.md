## gc taskmanagement workitems acd cancel create

Cancel the assignment process for a workitem that is currently queued for assignment through ACD.

### Synopsis

Cancel the assignment process for a workitem that is currently queued for assignment through ACD.

```
gc taskmanagement workitems acd cancel create [workitemId] [flags]
```

### Options

```
  -h, --help   help for create
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

* [gc taskmanagement workitems acd cancel](gc_taskmanagement_workitems_acd_cancel.html)	 - /api/v2/taskmanagement/workitems/{workitemId}/acd/cancel

