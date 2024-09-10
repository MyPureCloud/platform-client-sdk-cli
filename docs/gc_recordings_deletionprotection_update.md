## gc recordings deletionprotection update

Apply or revoke recording protection for conversations

### Synopsis

Apply or revoke recording protection for conversations

```
gc recordings deletionprotection update [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
  -h, --help               help for update
  -b, --printrequestbody   Print the request body format of the API.
      --protect string     Check for apply, uncheck for revoke (each action requires the respective permission) Valid values: true, false
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

* [gc recordings deletionprotection](gc_recordings_deletionprotection.html)	 - /api/v2/recordings/deletionprotection


