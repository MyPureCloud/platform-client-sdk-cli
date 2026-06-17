package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegisterarchitectvalidatejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegisterarchitectvalidatejobresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Registerarchitectvalidatejobresponse
type Registerarchitectvalidatejobresponse struct { 
    


    // Status - The status of the validate job.
    Status string `json:"status"`


    // TotalFlows - The number of flows submitted for validation.
    TotalFlows int `json:"totalFlows"`


    

}

// String returns a JSON representation of the model
func (o *Registerarchitectvalidatejobresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Registerarchitectvalidatejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Registerarchitectvalidatejobresponse

    if RegisterarchitectvalidatejobresponseMarshalled {
        return []byte("{}"), nil
    }
    RegisterarchitectvalidatejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        TotalFlows int `json:"totalFlows"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

