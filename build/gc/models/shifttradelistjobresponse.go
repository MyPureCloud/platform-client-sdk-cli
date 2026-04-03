package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradelistjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradelistjobresponseDud struct { 
    

}

// Shifttradelistjobresponse
type Shifttradelistjobresponse struct { 
    // Entities
    Entities []Shifttraderesponseitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Shifttradelistjobresponse) String() string {
     o.Entities = []Shifttraderesponseitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradelistjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Shifttradelistjobresponse

    if ShifttradelistjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ShifttradelistjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Shifttraderesponseitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Shifttraderesponseitem{{}},
        

        Alias: (*Alias)(u),
    })
}

