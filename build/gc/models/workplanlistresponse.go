package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanlistresponseDud struct { 
    

}

// Workplanlistresponse
type Workplanlistresponse struct { 
    // Entities
    Entities []Workplanlistitemresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Workplanlistresponse) String() string {
    
    
     o.Entities = []Workplanlistitemresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanlistresponse

    if WorkplanlistresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanlistresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Workplanlistitemresponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Workplanlistitemresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

