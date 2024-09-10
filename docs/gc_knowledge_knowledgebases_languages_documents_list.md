## gc knowledge knowledgebases languages documents list

Get documents

### Synopsis

Get documents

```
gc knowledge knowledgebases languages documents list [knowledgeBaseId] [languageCode] [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --categories string        Filter by categories ids, comma separated values expected.
      --documentIds strings      Comma-separated list of document identifiers to fetch by.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --limit string             Number of entities to return. Maximum of 200. Deprecated in favour of pageSize
      --pageSize string          Number of entities to return. Maximum of 200.
      --sortBy string            Sort by. Valid values: Title, Date
      --sortOrder string         Sort Order. Valid values: ASC, ascending, DESC, descending
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --title string             Filter by document title.
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

* [gc knowledge knowledgebases languages documents](gc_knowledge_knowledgebases_languages_documents.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents


