## gc taskmanagement workitems bulk jobs users me list

Get bulk jobs created by the currently logged in user.

### Synopsis

Get bulk jobs created by the currently logged in user.

```
gc taskmanagement workitems bulk jobs users me list [flags]
```

### Options

```
      --action string            The bulk job action. Valid values: TerminateWorkitems, AddWorkitems
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize after           Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an after key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200. (default "25")
      --sortOrder string         Ascending or descending sort order Valid values: ascending, descending
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

* [gc taskmanagement workitems bulk jobs users me](gc_taskmanagement_workitems_bulk_jobs_users_me.html)	 - /api/v2/taskmanagement/workitems/bulk/jobs/users/me


