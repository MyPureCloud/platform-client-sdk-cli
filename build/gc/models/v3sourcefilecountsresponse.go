package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefilecountsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefilecountsresponseDud struct { 
    

}

// V3sourcefilecountsresponse
type V3sourcefilecountsresponse struct { 
    // Entities - Per-source file counts.
    Entities []V3sourcefilecountentry `json:"entities"`

}

// String returns a JSON representation of the model
func (o *V3sourcefilecountsresponse) String() string {
     o.Entities = []V3sourcefilecountentry{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefilecountsresponse) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefilecountsresponse

    if V3sourcefilecountsresponseMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefilecountsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3sourcefilecountentry `json:"entities"`
        *Alias
    }{

        
        Entities: []V3sourcefilecountentry{{}},
        

        Alias: (*Alias)(u),
    })
}

