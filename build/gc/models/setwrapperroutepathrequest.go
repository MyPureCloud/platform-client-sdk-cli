package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetwrapperroutepathrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetwrapperroutepathrequestDud struct { 
    

}

// Setwrapperroutepathrequest
type Setwrapperroutepathrequest struct { 
    // Values
    Values []Routepathrequest `json:"values"`

}

// String returns a JSON representation of the model
func (o *Setwrapperroutepathrequest) String() string {
    
    
     o.Values = []Routepathrequest{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setwrapperroutepathrequest) MarshalJSON() ([]byte, error) {
    type Alias Setwrapperroutepathrequest

    if SetwrapperroutepathrequestMarshalled {
        return []byte("{}"), nil
    }
    SetwrapperroutepathrequestMarshalled = true

    return json.Marshal(&struct { 
        Values []Routepathrequest `json:"values"`
        
        *Alias
    }{
        

        
        Values: []Routepathrequest{{}},
        

        
        Alias: (*Alias)(u),
    })
}

