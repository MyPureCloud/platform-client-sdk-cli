package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgenetworkdiagnosticresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgenetworkdiagnosticresponseDud struct { 
    


    

}

// Edgenetworkdiagnosticresponse
type Edgenetworkdiagnosticresponse struct { 
    // CommandCorrelationId - UUID of each executed command on edge
    CommandCorrelationId string `json:"commandCorrelationId"`


    // Diagnostics - Response string of executed command from edge
    Diagnostics string `json:"diagnostics"`

}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgenetworkdiagnosticresponse) MarshalJSON() ([]byte, error) {
    type Alias Edgenetworkdiagnosticresponse

    if EdgenetworkdiagnosticresponseMarshalled {
        return []byte("{}"), nil
    }
    EdgenetworkdiagnosticresponseMarshalled = true

    return json.Marshal(&struct {
        
        CommandCorrelationId string `json:"commandCorrelationId"`
        
        Diagnostics string `json:"diagnostics"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

