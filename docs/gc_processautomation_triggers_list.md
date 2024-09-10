## gc processautomation triggers list

Retrieves all triggers, optionally filtered by query parameters.

### Synopsis

Retrieves all triggers, optionally filtered by query parameters.

```
gc processautomation triggers list [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --enabled string           Boolean indicating desired enabled state of triggers Valid values: true, false
      --filtercondition string   Filter list command output based on a given condition or regular expression
      --hasDelayBy string        Boolean to filter based on delayBySeconds being set in triggers. Default returns all, true returns only those with delayBySeconds set, false returns those without delayBySeconds set. Valid values: true, false
  -h, --help                     help for list
      --pageSize string          Number of entities to return. Maximum of 200.
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --topicName string         Topic name(s). Separated by commas
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

* [gc processautomation triggers](gc_processautomation_triggers.html)	 - /api/v2/processautomation/triggers


