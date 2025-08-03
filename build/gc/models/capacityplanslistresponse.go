package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanslistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanslistresponseDud struct { 
    

}

// Capacityplanslistresponse
type Capacityplanslistresponse struct { 
    // Entities
    Entities []Capacityplanlistitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Capacityplanslistresponse) String() string {
     o.Entities = []Capacityplanlistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanslistresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanslistresponse

    if CapacityplanslistresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanslistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Capacityplanlistitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Capacityplanlistitem{{}},
        

        Alias: (*Alias)(u),
    })
}

