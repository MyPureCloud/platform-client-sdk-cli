## gc journey actionmaps list

Retrieve all action maps.

### Synopsis

Retrieve all action maps.

```
gc journey actionmaps list [flags]
```

### Options

```
      --actionMapIds strings      IDs of action maps to return. Use of this parameter is not compatible with pagination, filtering, sorting or querying. A maximum of 100 action maps are allowed per request.
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --filterField filterField   Field to filter by (e.g. filterField=weight or filterField=action.actionTemplate.id). Requires filterField to also be set.
      --filterValue filterValue   Value to filter by. Requires filterValue to also be set.
      --filtercondition string    Filter list command output based on a given condition or regular expression
  -h, --help                      help for list
      --pageNumber string         Page number (default "1")
      --pageSize string           Page size (default "25")
      --queryFields queryValue    Action Map field(s) to query on. Requires queryValue to also be set.
      --queryValue queryFields    Value to query on using fuzzy matching. Requires queryFields to also be set.
      --sortBy -                  Field(s) to sort by. Prefix with - for descending (e.g. sortBy=displayName,-createdDate).
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
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

* [gc journey actionmaps](gc_journey_actionmaps.html)	 - /api/v2/journey/actionmaps


