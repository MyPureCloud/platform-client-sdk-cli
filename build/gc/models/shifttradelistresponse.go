package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradelistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradelistresponseDud struct { 
    

}

// Shifttradelistresponse
type Shifttradelistresponse struct { 
    // Entities
    Entities []Shifttraderesponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Shifttradelistresponse) String() string {
    
    
     o.Entities = []Shifttraderesponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradelistresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradelistresponse

    if ShifttradelistresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradelistresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Shifttraderesponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Shifttraderesponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

