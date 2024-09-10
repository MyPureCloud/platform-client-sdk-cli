## gc conversations messages communications messages create

Send message

### Synopsis

Send message

```
gc conversations messages communications messages create [conversationId] [communicationId] [flags]
```

### Options

```
  -d, --directory string              Directory path with files containing request bodies
  -f, --file string                   File name containing the JSON body
  -h, --help                          help for create
  -b, --printrequestbody              Print the request body format of the API.
      --useNormalizedMessage string   If true, response removes deprecated fields (textBody, media, stickers) Valid values: true, false
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

* [gc conversations messages communications messages](gc_conversations_messages_communications_messages.html)	 - /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages


