package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekshifttradelistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekshifttradelistresponseDud struct { 
    

}

// Weekshifttradelistresponse
type Weekshifttradelistresponse struct { 
    // Entities
    Entities []Weekshifttraderesponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Weekshifttradelistresponse) String() string {
    
    
     o.Entities = []Weekshifttraderesponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekshifttradelistresponse) MarshalJSON() ([]byte, error) {
    type Alias Weekshifttradelistresponse

    if WeekshifttradelistresponseMarshalled {
        return []byte("{}"), nil
    }
    WeekshifttradelistresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Weekshifttraderesponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Weekshifttraderesponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

