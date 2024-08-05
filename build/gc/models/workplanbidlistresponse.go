package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidlistresponseDud struct { 
    

}

// Workplanbidlistresponse
type Workplanbidlistresponse struct { 
    // Entities
    Entities []Workplanbid `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Workplanbidlistresponse) String() string {
     o.Entities = []Workplanbid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidlistresponse

    if WorkplanbidlistresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidlistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workplanbid `json:"entities"`
        *Alias
    }{

        
        Entities: []Workplanbid{{}},
        

        Alias: (*Alias)(u),
    })
}

