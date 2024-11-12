## gc infrastructureascode accelerators list

Get a list of available accelerators

### Synopsis

Get a list of available accelerators

```
gc infrastructureascode accelerators list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --classification string    Filter by classification
      --description string       Filter by description
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Filter by name
      --origin string            Filter by origin Valid values: community, partner, genesys
      --pageNumber string        The page number requested (default "1")
      --pageSize string          The total page size requested (default "25")
      --sortBy string            variable name requested to sort by
      --sortOrder string         Sort order Valid values: asc, desc
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --tags string              Filter by tags
      --varType string           Filter by type Valid values: module, accelerator, blueprint
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

* [gc infrastructureascode accelerators](gc_infrastructureascode_accelerators.html)	 - /api/v2/infrastructureascode/accelerators


