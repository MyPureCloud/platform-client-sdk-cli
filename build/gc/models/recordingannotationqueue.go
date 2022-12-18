package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingannotationqueueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingannotationqueueDud struct { 
    


    

}

// Recordingannotationqueue
type Recordingannotationqueue struct { 
    // Name - The queue name
    Name string `json:"name"`


    // Id - The queue Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Recordingannotationqueue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingannotationqueue) MarshalJSON() ([]byte, error) {
    type Alias Recordingannotationqueue

    if RecordingannotationqueueMarshalled {
        return []byte("{}"), nil
    }
    RecordingannotationqueueMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

