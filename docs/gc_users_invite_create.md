## gc users invite create

Send an activation email to the user

### Synopsis

Send an activation email to the user

```
gc users invite create [userId] [flags]
```

### Options

```
      --force string   Resend the invitation even if one is already outstanding Valid values: true, false
  -h, --help           help for create
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

* [gc users invite](gc_users_invite.html)	 - /api/v2/users/{userId}/invite


