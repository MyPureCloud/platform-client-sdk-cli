## gc knowledge sources synchronizations uploads create

Create presigned URL for uploading a file in the synchronization.

### Synopsis

Create presigned URL for uploading a file in the synchronization.

```
gc knowledge sources synchronizations uploads create [sourceId] [synchronizationId] [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
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

* [gc knowledge sources synchronizations uploads](gc_knowledge_sources_synchronizations_uploads.html)	 - /api/v2/knowledge/sources/{sourceId}/synchronizations/{synchronizationId}/uploads


