## gc conversations recordings update

Updates the retention records on a recording.

### Synopsis

Updates the retention records on a recording.

```
gc conversations recordings update [conversationId] [recordingId] [flags]
```

### Options

```
      --clearExport string   Whether to clear the pending export for the recording Valid values: true, false
  -d, --directory string     Directory path with files containing request bodies
  -f, --file string          File name containing the JSON body
  -h, --help                 help for update
  -b, --printrequestbody     Print the request body format of the API.
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

* [gc conversations recordings](gc_conversations_recordings.html)	 - /api/v2/conversations/{conversationId}/recordings


