package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegisterarchitectexportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegisterarchitectexportjobresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Registerarchitectexportjobresponse
type Registerarchitectexportjobresponse struct { 
    


    // Status - The status of the export job.
    Status string `json:"status"`


    // TotalFlows - The number of flows submitted for export.
    TotalFlows int `json:"totalFlows"`


    

}

// String returns a JSON representation of the model
func (o *Registerarchitectexportjobresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Registerarchitectexportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Registerarchitectexportjobresponse

    if RegisterarchitectexportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    RegisterarchitectexportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        TotalFlows int `json:"totalFlows"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

