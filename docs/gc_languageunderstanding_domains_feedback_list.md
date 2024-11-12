## gc languageunderstanding domains feedback list

Get all feedback in the given NLU Domain Version.

### Synopsis

Get all feedback in the given NLU Domain Version.

```
gc languageunderstanding domains feedback list [domainId] [flags]
```

### Options

```
      --after string                       The cursor that points to the end of the set of entities that has been returned. This is considered only when enableCursorPagination=true
      --assessment string                  The top assessment to retrieve feedback for. Valid values: Incorrect, Correct, Unknown, Disabled
  -a, --autopaginate                       Automatically paginate through the results stripping page information
      --dateEnd string                     End of time window as ISO-8601 date.
      --dateStart string                   Begin of time window as ISO-8601 date.
      --enableCursorPagination string      Enable Cursor Pagination Valid values: true, false
      --fields strings                     Fields and properties to get, comma-separated Valid values: version, dateCreated, text, intents
      --filtercondition string             Filter list command output based on a given condition or regular expression
  -h, --help                               help for list
      --includeDeleted string              Whether to include soft-deleted items in the result. Valid values: true, false
      --includeTrainingUtterances string   Include Training Utterances. By default they`re included. Valid values: true, false
      --intentName string                  The top intent name to retrieve feedback for.
      --language string                    Whether to filter response based on the language, e.g. en-us, pt-br.
      --pageNumber string                  Page number (default "1")
      --pageSize string                    Page size (default "25")
  -s, --stream                             Paginate and stream data as it is being processed leaving page information intact
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

* [gc languageunderstanding domains feedback](gc_languageunderstanding_domains_feedback.html)	 - /api/v2/languageunderstanding/domains/{domainId}/feedback


