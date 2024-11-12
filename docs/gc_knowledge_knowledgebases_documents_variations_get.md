## gc knowledge knowledgebases documents variations get

Get a variation for a document.

### Synopsis

Get a variation for a document.

```
gc knowledge knowledgebases documents variations get [documentVariationId] [documentId] [knowledgeBaseId] [flags]
```

### Options

```
      --documentState string   The state of the document. Valid values: Draft, Published
      --expand strings         The specified entity attributes will be filled. Comma separated values expected. Valid values: contentUrl
  -h, --help                   help for get
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

* [gc knowledge knowledgebases documents variations](gc_knowledge_knowledgebases_documents_variations.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations


