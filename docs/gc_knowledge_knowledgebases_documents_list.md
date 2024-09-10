## gc knowledge knowledgebases documents list

Get documents.

### Synopsis

Get documents.

```
gc knowledge knowledgebases documents list [knowledgeBaseId] [flags]
```

### Options

```
      --after string                      The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate                      Automatically paginate through the results stripping page information
      --before string                     The cursor that points to the start of the set of entities that has been returned.
      --categoryId strings                If specified, retrieves documents associated with category ids, comma separated values expected.
      --documentId strings                Retrieves the specified documents, comma separated values expected.
      --expand strings                    The specified entity attributes will be filled. Comma separated values expected. Valid values: category, labels, variations
      --externalIds strings               If specified, retrieves documents associated with external ids, comma separated values expected.
      --filtercondition string            Filter list command output based on a given condition or regular expression
  -h, --help                              help for list
      --includeDrafts string              If includeDrafts is true, Documents in the draft state are also returned in the response. Valid values: true, false
      --includeSubcategories categoryId   Works along with categoryId query parameter. If specified, retrieves documents associated with category ids and its children categories. Valid values: true, false
      --interval string                   Retrieves the documents modified in specified date and time range. If the after and before cursor parameters are within this interval, it would return valid data, otherwise it throws an error.The dates in the interval are represented in ISO-8601 format: YYYY-MM-DDThh:mm:ssZ/YYYY-MM-DDThh:mm:ssZ
      --labelIds strings                  If specified, retrieves documents associated with label ids, comma separated values expected.
      --pageSize string                   Number of entities to return. Maximum of 200.
  -s, --stream                            Paginate and stream data as it is being processed leaving page information intact
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

* [gc knowledge knowledgebases documents](gc_knowledge_knowledgebases_documents.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents


