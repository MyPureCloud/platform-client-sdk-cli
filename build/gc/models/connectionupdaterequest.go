package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectionupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectionupdaterequestDud struct { 
    


    


    

}

// Connectionupdaterequest
type Connectionupdaterequest struct { 
    // Code - Unique code that allows to be connected
    Code string `json:"code"`


    // VarError - Name of the received error
    VarError string `json:"error"`


    // ErrorDescription - Detailed description of the error
    ErrorDescription string `json:"errorDescription"`

}

// String returns a JSON representation of the model
func (o *Connectionupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectionupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Connectionupdaterequest

    if ConnectionupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    ConnectionupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        VarError string `json:"error"`
        
        ErrorDescription string `json:"errorDescription"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

