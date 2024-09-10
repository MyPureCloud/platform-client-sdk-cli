## gc architect prompts list

Get a pageable list of user prompts

### Synopsis

Get a pageable list of user prompts

```
gc architect prompts list [flags]
```

### Options

```
  -a, --autopaginate               Automatically paginate through the results stripping page information
      --description string         Description
      --filtercondition string     Filter list command output based on a given condition or regular expression
  -h, --help                       help for list
      --includeMediaUris string    Include the media URIs for each resource Valid values: true, false
      --includeResources string    Include the resources for each system prompt Valid values: true, false
      --language strings           Filter the resources down to the provided languages
      --name strings               Name
      --nameOrDescription string   Name or description
      --pageNumber int             Page number (default 1)
      --pageSize int               Page size (default 25)
      --sortBy string              Sort by
      --sortOrder string           Sort order
  -s, --stream                     Paginate and stream data as it is being processed leaving page information intact
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

* [gc architect prompts](gc_architect_prompts.html)	 - /api/v2/architect/prompts


