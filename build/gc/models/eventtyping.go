package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventtypingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventtypingDud struct { 
    


    Duration int `json:"duration"`

}

// Eventtyping - A Typing event.
type Eventtyping struct { 
    // VarType - Describes the type of Typing event.
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Eventtyping) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventtyping) MarshalJSON() ([]byte, error) {
    type Alias Eventtyping

    if EventtypingMarshalled {
        return []byte("{}"), nil
    }
    EventtypingMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

