## gc integrations update

Update an integration.

### Synopsis

Update an integration.

```
gc integrations update [integrationId] [flags]
```

### Options

```
  -d, --directory string      Directory path with files containing request bodies
      --expand strings        variable name requested by expand list
  -f, --file string           File name containing the JSON body
  -h, --help                  help for update
      --nextPage string       next page token
      --pageNumber int        The page number requested (default 1)
      --pageSize int          The total page size requested (default 25)
      --previousPage string   Previous page token
  -b, --printrequestbody      Print the request body format of the API.
      --sortBy string         variable name requested to sort by
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

* [gc integrations](gc_integrations.html)	 - /api/v2/integrations


