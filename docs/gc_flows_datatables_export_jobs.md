## gc flows datatables export jobs

/api/v2/flows/datatables/{datatableId}/export/jobs

### Synopsis

/api/v2/flows/datatables/{datatableId}/export/jobs

### Options

```
  -h, --help   help for jobs
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

* [gc flows datatables export](gc_flows_datatables_export.html)	 - /api/v2/flows/datatables/{datatableId}/export
* [gc flows datatables export jobs create](gc_flows_datatables_export_jobs_create.html)	 - Begin an export process for exporting all rows from a datatable
* [gc flows datatables export jobs get](gc_flows_datatables_export_jobs_get.html)	 - Returns the state information about an export job


