## gc taskmanagement workitems users wrapups get

Get all wrapup codes added for the given user for a workitem.

### Synopsis

Get all wrapup codes added for the given user for a workitem.

```
gc taskmanagement workitems users wrapups get [workitemId] [userId] [flags]
```

### Options

```
      --after string       The cursor that points to the end of the set of entities that has been returned.
      --expands string     Which fields, if any, to expand. Valid values: wrapupCode
  -h, --help               help for get
      --pageSize after     Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an after key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 50. (default 25)
      --sortOrder string   Ascending or descending sort order Valid values: ascending, descending
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

* [gc taskmanagement workitems users wrapups](gc_taskmanagement_workitems_users_wrapups.html)	 - /api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups


