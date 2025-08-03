package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanstaffinggroupmetricchangehistorylistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanstaffinggroupmetricchangehistorylistresponseDud struct { 
    

}

// Capacityplanstaffinggroupmetricchangehistorylistresponse
type Capacityplanstaffinggroupmetricchangehistorylistresponse struct { 
    // Entities
    Entities []Staffinggroupmetricchangeresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Capacityplanstaffinggroupmetricchangehistorylistresponse) String() string {
     o.Entities = []Staffinggroupmetricchangeresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanstaffinggroupmetricchangehistorylistresponse) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanstaffinggroupmetricchangehistorylistresponse

    if CapacityplanstaffinggroupmetricchangehistorylistresponseMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanstaffinggroupmetricchangehistorylistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Staffinggroupmetricchangeresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Staffinggroupmetricchangeresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

