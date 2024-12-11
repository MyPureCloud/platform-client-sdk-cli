package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobstatusupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobstatusupdaterequestDud struct { 
    


    

}

// Contactimportjobstatusupdaterequest
type Contactimportjobstatusupdaterequest struct { 
    // JobId - Job Id
    JobId string `json:"jobId"`


    // Status - Workflow status
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Contactimportjobstatusupdaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobstatusupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobstatusupdaterequest

    if ContactimportjobstatusupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobstatusupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

