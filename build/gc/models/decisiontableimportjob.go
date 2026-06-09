package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableimportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableimportjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    CreatedBy Addressableentityref `json:"createdBy"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Decisiontableimportjob - State of a decision table row import job
type Decisiontableimportjob struct { 
    


    // TableVersion - The table version to be replaced by this import
    TableVersion int `json:"tableVersion"`


    // Status - Current status of the import job
    Status string `json:"status"`


    // UploadUrl - Pre-signed URL to upload the import file (PUT)
    UploadUrl string `json:"uploadUrl"`


    // UploadHeaders - Headers required when uploading file with data to be imported to uploadUrl
    UploadHeaders map[string]string `json:"uploadHeaders"`


    // ImportMode - Whether rows are appended to existing rows or rows are replaced
    ImportMode string `json:"importMode"`


    // FileName - Original file name supplied when the job was created, including the file extension
    FileName string `json:"fileName"`


    


    // DateCreated - When the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - When the job was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateCompleted - When processing finished, successfully or not. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // DateExpires - When upload credentials expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`


    // RowMetrics - Row-level metrics populated incrementally during import processing
    RowMetrics Decisiontableimportrowmetrics `json:"rowMetrics"`


    // VarError - Present when the import job could not be successfully finished
    VarError Decisiontableimportjoberror `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Decisiontableimportjob) String() string {
    
    
    
     o.UploadHeaders = map[string]string{"": ""} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableimportjob) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableimportjob

    if DecisiontableimportjobMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableimportjobMarshalled = true

    return json.Marshal(&struct {
        
        TableVersion int `json:"tableVersion"`
        
        Status string `json:"status"`
        
        UploadUrl string `json:"uploadUrl"`
        
        UploadHeaders map[string]string `json:"uploadHeaders"`
        
        ImportMode string `json:"importMode"`
        
        FileName string `json:"fileName"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        RowMetrics Decisiontableimportrowmetrics `json:"rowMetrics"`
        
        VarError Decisiontableimportjoberror `json:"error"`
        *Alias
    }{

        


        


        


        


        
        UploadHeaders: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

