## gc conversations chats messages list

Get the messages of a chat conversation.

### Synopsis

Get the messages of a chat conversation.

```
gc conversations chats messages list [conversationId] [flags]
```

### Options

```
      --after string        If specified, get the messages chronologically after the id of this message
      --before string       If specified, get the messages chronologically before the id of this message
  -h, --help                help for list
      --maxResults string   Limit the returned number of messages, up to a maximum of 100 (default "100")
      --sortOrder string    Sort order Valid values: ascending, descending
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

* [gc conversations chats messages](gc_conversations_chats_messages.html)	 - /api/v2/conversations/chats/{conversationId}/messages


