## gc knowledge knowledgebases documents get

Get document.

### Synopsis

Get document.

```
gc knowledge knowledgebases documents get [knowledgeBaseId] [documentId] [flags]
```

### Options

```
      --expand strings   The specified entity attributes will be filled. Comma separated values expected. Max No. of variations that can be returned on expand is 20. Valid values: category, labels, variations
  -h, --help             help for get
      --state string     when state is Draft, draft version of the document is returned,otherwise by default published version is returned in the response. Valid values: Draft, Published
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


