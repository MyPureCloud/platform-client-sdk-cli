package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationeventtypingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationeventtypingDud struct { 
    


    

}

// Conversationeventtyping - A Typing event.
type Conversationeventtyping struct { 
    // VarType - Describes the type of Typing event.
    VarType string `json:"type"`


    // Duration - The duration of the Typing event in milliseconds.
    Duration int `json:"duration"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtyping) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationeventtyping) MarshalJSON() ([]byte, error) {
    type Alias Conversationeventtyping

    if ConversationeventtypingMarshalled {
        return []byte("{}"), nil
    }
    ConversationeventtypingMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Duration int `json:"duration"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

