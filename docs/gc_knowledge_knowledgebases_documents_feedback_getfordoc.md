## gc knowledge knowledgebases documents feedback getfordoc

Get a list of feedback records given on a document

### Synopsis

Get a list of feedback records given on a document

```
gc knowledge knowledgebases documents feedback getfordoc [knowledgeBaseId] [documentId] [flags]
```

### Options

```
      --after string                 The cursor that points to the end of the set of entities that has been returned.
      --appType string               Application type to filter by. Supported only if onlyCommented=true is set. Valid values: Assistant, BotFlow, MessengerKnowledgeApp, SmartAdvisor, SupportCenter
  -a, --autopaginate                 Automatically paginate through the results stripping page information
      --before string                The cursor that points to the start of the set of entities that has been returned.
      --documentVariationId string   Document variation ID to filter by. Supported only if onlyCommented=true is set.
      --documentVersionId string     Document version ID to filter by. Supported only if onlyCommented=true is set.
      --filtercondition string       Filter list command output based on a given condition or regular expression
  -h, --help                         help for getfordoc
      --onlyCommented string         If true, only feedback records that have comment are returned. If false, feedback records with and without comment are returned. Default: false. Valid values: true, false
      --pageSize string              Number of entities to return. Maximum of 200.
      --queryType string             Query type to filter by. Supported only if onlyCommented=true is set. Valid values: Unknown, Article, AutoSearch, Category, ManualSearch, Recommendation, Suggestion, ExpandedArticle
      --queueId string               Queue ID to filter by. Supported only if onlyCommented=true is set.
      --state string                 State to filter by. Supported only if onlyCommented=true is set. Default: Final Valid values: All, Draft, Final
  -s, --stream                       Paginate and stream data as it is being processed leaving page information intact
      --userId string                The ID of the user, who created the feedback, to filter by. Supported only if onlyCommented=true is set.
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

* [gc knowledge knowledgebases documents feedback](gc_knowledge_knowledgebases_documents_feedback.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback


