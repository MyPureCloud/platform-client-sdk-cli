## gc socialmedia topics list

Retrieve all social topics.

### Synopsis

Retrieve all social topics.

```
gc socialmedia topics list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionIds strings      One or more division IDs. If nothing is provided, the social topics associated withthe list of divisions that the user has access to will be returned.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --ids strings              One or more topic IDs to search through the topics.
      --includeDeleted string    Determines whether to include soft-deleted items in the result. Valid values: true, false
      --name string              Search for topic by name that contains the given search string, search is case insensitive
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
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

* [gc socialmedia topics](gc_socialmedia_topics.html)	 - /api/v2/socialmedia/topics


