## gc taskmanagement workitems get

Get a workitem

### Synopsis

Get a workitem

```
gc taskmanagement workitems get [workitemId] [flags]
```

### Options

```
      --expands string   Which fields to expand. Comma separated if more than one. Valid values: type, workbin, status, queue, assignee
  -h, --help             help for get
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

* [gc taskmanagement workitems](gc_taskmanagement_workitems.html)	 - /api/v2/taskmanagement/workitems /api/v2/taskmanagement/workitems/{workitemId}/users


