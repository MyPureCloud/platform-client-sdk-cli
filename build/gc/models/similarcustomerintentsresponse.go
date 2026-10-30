package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarcustomerintentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarcustomerintentsresponseDud struct { 
    

}

// Similarcustomerintentsresponse
type Similarcustomerintentsresponse struct { 
    // Entities
    Entities []Similarcustomerintentresult `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Similarcustomerintentsresponse) String() string {
     o.Entities = []Similarcustomerintentresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarcustomerintentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Similarcustomerintentsresponse

    if SimilarcustomerintentsresponseMarshalled {
        return []byte("{}"), nil
    }
    SimilarcustomerintentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Similarcustomerintentresult `json:"entities"`
        *Alias
    }{

        
        Entities: []Similarcustomerintentresult{{}},
        

        Alias: (*Alias)(u),
    })
}

