package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanrotationlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanrotationlistresponseDud struct { 
    

}

// Workplanrotationlistresponse
type Workplanrotationlistresponse struct { 
    // Entities
    Entities []Workplanrotationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationlistresponse) String() string {
    
    
     o.Entities = []Workplanrotationresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanrotationlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanrotationlistresponse

    if WorkplanrotationlistresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanrotationlistresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Workplanrotationresponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Workplanrotationresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

