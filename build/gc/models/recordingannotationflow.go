package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingannotationflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingannotationflowDud struct { 
    


    

}

// Recordingannotationflow
type Recordingannotationflow struct { 
    // Name - The flow name
    Name string `json:"name"`


    // Id - The flow Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Recordingannotationflow) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingannotationflow) MarshalJSON() ([]byte, error) {
    type Alias Recordingannotationflow

    if RecordingannotationflowMarshalled {
        return []byte("{}"), nil
    }
    RecordingannotationflowMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

