package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventdefinitionDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Description string `json:"description"`

}

// Eventdefinition
type Eventdefinition struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Eventdefinition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventdefinition) MarshalJSON() ([]byte, error) {
    type Alias Eventdefinition

    if EventdefinitionMarshalled {
        return []byte("{}"), nil
    }
    EventdefinitionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

