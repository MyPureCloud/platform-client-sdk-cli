package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradepreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradepreviewresponseDud struct { 
    

}

// Shifttradepreviewresponse
type Shifttradepreviewresponse struct { 
    // Activities - List of activities that will make up the new shift if this shift trade is approved
    Activities []Shifttradeactivitypreviewresponse `json:"activities"`

}

// String returns a JSON representation of the model
func (o *Shifttradepreviewresponse) String() string {
    
    
     o.Activities = []Shifttradeactivitypreviewresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradepreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradepreviewresponse

    if ShifttradepreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradepreviewresponseMarshalled = true

    return json.Marshal(&struct { 
        Activities []Shifttradeactivitypreviewresponse `json:"activities"`
        
        *Alias
    }{
        

        
        Activities: []Shifttradeactivitypreviewresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

