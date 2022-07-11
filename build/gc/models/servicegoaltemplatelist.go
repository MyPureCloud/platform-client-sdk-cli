package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicegoaltemplatelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicegoaltemplatelistDud struct { 
    

}

// Servicegoaltemplatelist
type Servicegoaltemplatelist struct { 
    // Entities
    Entities []Servicegoaltemplate `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatelist) String() string {
     o.Entities = []Servicegoaltemplate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicegoaltemplatelist) MarshalJSON() ([]byte, error) {
    type Alias Servicegoaltemplatelist

    if ServicegoaltemplatelistMarshalled {
        return []byte("{}"), nil
    }
    ServicegoaltemplatelistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Servicegoaltemplate `json:"entities"`
        *Alias
    }{

        
        Entities: []Servicegoaltemplate{{}},
        

        Alias: (*Alias)(u),
    })
}

