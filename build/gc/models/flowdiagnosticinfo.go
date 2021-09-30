package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowdiagnosticinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowdiagnosticinfoDud struct { 
    

}

// Flowdiagnosticinfo
type Flowdiagnosticinfo struct { 
    // LastActionId - The step number of the survey invite flow where the error occurred.
    LastActionId int `json:"lastActionId"`

}

// String returns a JSON representation of the model
func (o *Flowdiagnosticinfo) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowdiagnosticinfo) MarshalJSON() ([]byte, error) {
    type Alias Flowdiagnosticinfo

    if FlowdiagnosticinfoMarshalled {
        return []byte("{}"), nil
    }
    FlowdiagnosticinfoMarshalled = true

    return json.Marshal(&struct { 
        LastActionId int `json:"lastActionId"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

