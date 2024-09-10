## gc contentmanagement documents update

Update a document.

### Synopsis

Update a document.

```
gc contentmanagement documents update [documentId] [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
      --expand string      Expand some document fields Valid values: acl
  -f, --file string        File name containing the JSON body
  -h, --help               help for update
      --override string    Override any lock on the document Valid values: true, false
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

* [gc contentmanagement documents](gc_contentmanagement_documents.html)	 - /api/v2/contentmanagement/documents


