package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobstatusupdateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobstatusupdateresponseDud struct { 
    


    

}

// Contactimportjobstatusupdateresponse
type Contactimportjobstatusupdateresponse struct { 
    // JobId
    JobId string `json:"jobId"`


    // Status
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Contactimportjobstatusupdateresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobstatusupdateresponse) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobstatusupdateresponse

    if ContactimportjobstatusupdateresponseMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobstatusupdateresponseMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

