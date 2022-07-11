package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionenddetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionenddetailsDud struct { 
    


    

}

// Sessionenddetails
type Sessionenddetails struct { 
    // VarType - The type of termination handling that resulted in the session end. It can be either Exit or Disconnect
    VarType string `json:"type"`


    // Reason - The reason for termination action. It can be due to an error or normal flow execution
    Reason string `json:"reason"`

}

// String returns a JSON representation of the model
func (o *Sessionenddetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionenddetails) MarshalJSON() ([]byte, error) {
    type Alias Sessionenddetails

    if SessionenddetailsMarshalled {
        return []byte("{}"), nil
    }
    SessionenddetailsMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Reason string `json:"reason"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

