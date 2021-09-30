package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradematchessummaryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradematchessummaryresponseDud struct { 
    

}

// Shifttradematchessummaryresponse
type Shifttradematchessummaryresponse struct { 
    // Entities
    Entities []Weekshifttradematchessummaryresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchessummaryresponse) String() string {
    
    
     o.Entities = []Weekshifttradematchessummaryresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradematchessummaryresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradematchessummaryresponse

    if ShifttradematchessummaryresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradematchessummaryresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Weekshifttradematchessummaryresponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Weekshifttradematchessummaryresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

