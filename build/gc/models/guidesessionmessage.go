package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionmessageDud struct { 
    


    

}

// Guidesessionmessage - A message in the conversation history provided to a guide session turn.
type Guidesessionmessage struct { 
    // Role - The role of the message author.
    Role string `json:"role"`


    // Content - The content of the message.
    Content string `json:"content"`

}

// String returns a JSON representation of the model
func (o *Guidesessionmessage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionmessage) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionmessage

    if GuidesessionmessageMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionmessageMarshalled = true

    return json.Marshal(&struct {
        
        Role string `json:"role"`
        
        Content string `json:"content"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

