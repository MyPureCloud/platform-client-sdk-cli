## gc knowledge knowledgebases sources list

Get Knowledge integration sources

### Synopsis

Get Knowledge integration sources

```
gc knowledge knowledgebases sources list [knowledgeBaseId] [flags]
```

### Options

```
      --expand strings   The specified entity attributes will be filled. Comma separated values expected. Valid values: lastsync
  -h, --help             help for list
      --ids strings      If specified, retrieves integration sources with specified IDs.
      --varType string   If specified, retrieves integration sources with specified integration type. Valid values: Salesforce, ServiceNow
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

* [gc knowledge knowledgebases sources](gc_knowledge_knowledgebases_sources.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources


