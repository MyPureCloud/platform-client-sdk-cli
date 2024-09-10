## gc journey actiontemplates list

Retrieve all action templates.

### Synopsis

Retrieve all action templates.

```
gc journey actiontemplates list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --mediaType string         Media type Valid values: webchat, webMessagingOffer, contentOffer, integrationAction, architectFlow, openAction
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --queryFields queryValue   ActionTemplate field(s) to query on. Requires queryValue to also be set.
      --queryValue queryFields   Value to query on. Requires queryFields to also be set.
      --sortBy -                 Field(s) to sort by. Prefix with - for descending (e.g. sortBy=name,-createdDate).
      --state string             Action template state. Valid values: Active, Inactive, Deleted
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

* [gc journey actiontemplates](gc_journey_actiontemplates.html)	 - /api/v2/journey/actiontemplates


