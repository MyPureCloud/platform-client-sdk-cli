package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Contactimportjobresponse
type Contactimportjobresponse struct { 
    


    // Status
    Status string `json:"status"`


    // StatusDetails - Detailed description for JobStatus.
    StatusDetails string `json:"statusDetails"`


    // ExecutionStep - Detailed description for the Job execution state
    ExecutionStep string `json:"executionStep"`


    // Metadata - Metadata for the job
    Metadata Contactimportjobmetadata `json:"metadata"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    


    // Settings - Settings
    Settings Addressableentityref `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Contactimportjobresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobresponse

    if ContactimportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        StatusDetails string `json:"statusDetails"`
        
        ExecutionStep string `json:"executionStep"`
        
        Metadata Contactimportjobmetadata `json:"metadata"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        Settings Addressableentityref `json:"settings"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

