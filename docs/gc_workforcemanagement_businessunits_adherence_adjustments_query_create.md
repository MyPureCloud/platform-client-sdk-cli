## gc workforcemanagement businessunits adherence adjustments query create

Query adherence adjustments for a business unit. Results will be returned using cursor pagination

### Synopsis

Query adherence adjustments for a business unit. Results will be returned using cursor pagination

```
gc workforcemanagement businessunits adherence adjustments query create [businessUnitId] [flags]
```

### Options

```
      --after string       The cursor that points to the end of the set of entities that has been returned.
      --before string      The cursor that points to the start of the set of entities that has been returned.
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
  -h, --help               help for create
      --pageSize string    The page size for the listing. The maximum page size is 500.
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

* [gc workforcemanagement businessunits adherence adjustments query](gc_workforcemanagement_businessunits_adherence_adjustments_query.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/adjustments/query


