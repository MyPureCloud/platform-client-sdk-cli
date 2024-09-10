## gc webchat guest conversations members list

Get the members of a chat conversation.

### Synopsis

Get the members of a chat conversation.

```
gc webchat guest conversations members list [conversationId] [flags]
```

### Options

```
  -a, --autopaginate                        Automatically paginate through the results stripping page information
      --excludeDisconnectedMembers string   If true, the results will not contain members who have a DISCONNECTED state. Valid values: true, false
      --filtercondition string              Filter list command output based on a given condition or regular expression
  -h, --help                                help for list
      --pageNumber int                      The page number to return, or omitted for the first page. (default 1)
      --pageSize int                        The number of entries to return per page, or omitted for the default. (default 25)
  -s, --stream                              Paginate and stream data as it is being processed leaving page information intact
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

* [gc webchat guest conversations members](gc_webchat_guest_conversations_members.html)	 - /api/v2/webchat/guest/conversations/{conversationId}/members


