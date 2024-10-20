## gc contentmanagement shared get

Get shared documents. Securely download a shared document.

### Synopsis

Get shared documents. Securely download a shared document.

```
gc contentmanagement shared get [sharedId] [flags]
```

### Options

```
      --contentType string   The requested format for the specified document. If supported, the document will be returned in that format. Example contentType=audio/wav
      --disposition string   Request how the share content will be downloaded: attached as a file or inline. Default is attachment. Valid values: attachment, inline, none
      --expand string        Expand some document fields Valid values: document.acl
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

* [gc contentmanagement shared](gc_contentmanagement_shared.html)	 - /api/v2/contentmanagement/shared


