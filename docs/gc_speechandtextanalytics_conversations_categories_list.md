## gc speechandtextanalytics conversations categories list

Get the list of detected Speech and Text Analytics categories of conversation

### Synopsis

Get the list of detected Speech and Text Analytics categories of conversation

```
gc speechandtextanalytics conversations categories list [conversationId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           The page number for the listing (default 1)
      --pageSize int             The page size for the listing. The max that will be returned is 50. (default 25)
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

* [gc speechandtextanalytics conversations categories](gc_speechandtextanalytics_conversations_categories.html)	 - /api/v2/speechandtextanalytics/conversations/{conversationId}/categories


