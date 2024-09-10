## gc conversations emails inboundmessages

/api/v2/conversations/emails/{conversationId}/inboundmessages

### Synopsis

/api/v2/conversations/emails/{conversationId}/inboundmessages

### Options

```
  -h, --help   help for inboundmessages
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

* [gc conversations emails](gc_conversations_emails.html)	 - /api/v2/conversations/emails
* [gc conversations emails inboundmessages create](gc_conversations_emails_inboundmessages_create.html)	 - Send an email to an external conversation. An external conversation is one where the provider is not PureCloud based. This endpoint allows the sender of the external email to reply or send a new message to the existing conversation. The new message will be treated as part of the existing conversation and chained to it.


