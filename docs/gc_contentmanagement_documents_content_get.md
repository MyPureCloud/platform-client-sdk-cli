## gc contentmanagement documents content get

Download a document.

### Synopsis

Download a document.

```
gc contentmanagement documents content get [documentId] [flags]
```

### Options

```
      --contentType string   The requested format for the specified document. If supported, the document will be returned in that format. Example contentType=audio/wav
      --disposition string   Request how the content will be downloaded: a file attachment or inline. Default is attachment. Valid values: attachment, inline
  -h, --help                 help for get
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

* [gc contentmanagement documents content](gc_contentmanagement_documents_content.html)	 - /api/v2/contentmanagement/documents/{documentId}/content


