## gc knowledge knowledgebases documents query create

Query for knowledge documents.

### Synopsis

Query for knowledge documents.

```
gc knowledge knowledgebases documents query create [knowledgeBaseId] [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
      --expand strings     Fields, if any, to expand for each document in the search result matching the query. Valid values: documentVariations, documentAlternatives, knowledgeBaseLanguageCode
  -f, --file string        File name containing the JSON body
  -h, --help               help for create
  -b, --printrequestbody   Print the request body format of the API.
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

* [gc knowledge knowledgebases documents query](gc_knowledge_knowledgebases_documents_query.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/query

