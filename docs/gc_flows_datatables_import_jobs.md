## gc flows datatables import jobs

/api/v2/flows/datatables/{datatableId}/import/jobs

### Synopsis

/api/v2/flows/datatables/{datatableId}/import/jobs

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

* [gc flows datatables import](gc_flows_datatables_import.html)	 - /api/v2/flows/datatables/{datatableId}/import
* [gc flows datatables import jobs create](gc_flows_datatables_import_jobs_create.html)	 - Begin an import process for importing rows into a datatable
* [gc flows datatables import jobs getstateinformation](gc_flows_datatables_import_jobs_getstateinformation.html)	 - Returns the state information about an import job
* [gc flows datatables import jobs list](gc_flows_datatables_import_jobs_list.html)	 - Get all recent import jobs


