package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationcreateconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationcreateconversationDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Evaluationcreateconversation
type Evaluationcreateconversation struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationcreateconversation) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationcreateconversation) MarshalJSON() ([]byte, error) {
    type Alias Evaluationcreateconversation

    if EvaluationcreateconversationMarshalled {
        return []byte("{}"), nil
    }
    EvaluationcreateconversationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

