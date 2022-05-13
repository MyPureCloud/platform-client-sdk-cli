package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatatableimportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatatableimportjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Datatableimportjob - State information for an import job of rows to a datatable
type Datatableimportjob struct { 
    


    // Name
    Name string `json:"name"`


    // Owner - The PureCloud user who started the import job
    Owner Addressableentityref `json:"owner"`


    // Status - The status of the import job
    Status string `json:"status"`


    // DateCreated - The timestamp of when the import began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateCompleted - The timestamp of when the import stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // UploadURI - The URL of the location at which the caller can upload the file to be imported
    UploadURI string `json:"uploadURI"`


    // ImportMode - The indication of whether the processing should remove rows that don't appear in the import file
    ImportMode string `json:"importMode"`


    // ErrorInformation - Any error information, or null of the processing is not in an error state
    ErrorInformation Errorbody `json:"errorInformation"`


    // CountRecordsUpdated - The current count of the number of records processed
    CountRecordsUpdated int `json:"countRecordsUpdated"`


    // CountRecordsDeleted - The current count of the number of records deleted
    CountRecordsDeleted int `json:"countRecordsDeleted"`


    // CountRecordsFailed - The current count of the number of records that failed to import
    CountRecordsFailed int `json:"countRecordsFailed"`


    

}

// String returns a JSON representation of the model
func (o *Datatableimportjob) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datatableimportjob) MarshalJSON() ([]byte, error) {
    type Alias Datatableimportjob

    if DatatableimportjobMarshalled {
        return []byte("{}"), nil
    }
    DatatableimportjobMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Owner Addressableentityref `json:"owner"`
        
        Status string `json:"status"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        UploadURI string `json:"uploadURI"`
        
        ImportMode string `json:"importMode"`
        
        ErrorInformation Errorbody `json:"errorInformation"`
        
        CountRecordsUpdated int `json:"countRecordsUpdated"`
        
        CountRecordsDeleted int `json:"countRecordsDeleted"`
        
        CountRecordsFailed int `json:"countRecordsFailed"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

