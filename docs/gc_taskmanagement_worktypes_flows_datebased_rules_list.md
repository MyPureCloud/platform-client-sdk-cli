## gc taskmanagement worktypes flows datebased rules list

Get all date based rules for a worktype

### Synopsis

Get all date based rules for a worktype

```
gc taskmanagement worktypes flows datebased rules list [worktypeId] [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize after           Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an after key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200. (default "25")
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

* [gc taskmanagement worktypes flows datebased rules](gc_taskmanagement_worktypes_flows_datebased_rules.html)	 - /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules


