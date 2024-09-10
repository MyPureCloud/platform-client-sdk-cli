## gc knowledge knowledgebases languages trainings list

Get all trainings information for a knowledgebase

### Synopsis

Get all trainings information for a knowledgebase

```
gc knowledge knowledgebases languages trainings list [knowledgeBaseId] [languageCode] [flags]
```

### Options

```
      --after string                     The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate                     Automatically paginate through the results stripping page information
      --before string                    The cursor that points to the start of the set of entities that has been returned.
      --filtercondition string           Filter list command output based on a given condition or regular expression
  -h, --help                             help for list
      --knowledgeDocumentsState string   Return the training with the specified state of the trained documents. Valid values: Draft, Active, Discarded, Archived
      --limit string                     Number of entities to return. Maximum of 200. Deprecated in favour of pageSize
      --pageSize string                  Number of entities to return. Maximum of 200.
  -s, --stream                           Paginate and stream data as it is being processed leaving page information intact
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

* [gc knowledge knowledgebases languages trainings](gc_knowledge_knowledgebases_languages_trainings.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings


