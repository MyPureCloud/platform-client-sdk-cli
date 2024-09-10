## gc conversations emails messages draft patch

Reset conversation draft to its initial state and/or auto-fill draft content

### Synopsis

Reset conversation draft to its initial state and/or auto-fill draft content

```
gc conversations emails messages draft patch [conversationId] [flags]
```

### Options

```
      --autoFill string    autoFill Valid values: true, false
  -d, --directory string   Directory path with files containing request bodies
      --discard string     discard Valid values: true, false
  -f, --file string        File name containing the JSON body
  -h, --help               help for patch
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

* [gc conversations emails messages draft](gc_conversations_emails_messages_draft.html)	 - /api/v2/conversations/emails/{conversationId}/messages/draft


