## gc users chats me list

Get chats for a user

### Synopsis

Get chats for a user

```
gc users chats me list [flags]
```

### Options

```
      --after string             The key to start after
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --excludeClosed string     Whether or not to exclude closed chats Valid values: true, false
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --includePresence string   Whether or not to include user presence Valid values: true, false
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc users chats me](gc_users_chats_me.html)	 - /api/v2/users/chats/me


