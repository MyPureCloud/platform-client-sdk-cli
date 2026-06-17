package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidlistresponseDud struct { 
    

}

// Schedulebidlistresponse
type Schedulebidlistresponse struct { 
    // Entities
    Entities []Schedulebid `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Schedulebidlistresponse) String() string {
     o.Entities = []Schedulebid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidlistresponse

    if SchedulebidlistresponseMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidlistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Schedulebid `json:"entities"`
        *Alias
    }{

        
        Entities: []Schedulebid{{}},
        

        Alias: (*Alias)(u),
    })
}

