package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatreactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatreactionDud struct { 
    


    

}

// Chatreaction
type Chatreaction struct { 
    // Emoji - The emoji string for the reaction
    Emoji string `json:"emoji"`


    // Users - The users that reacted with an emoji
    Users []Addressableentityref `json:"users"`

}

// String returns a JSON representation of the model
func (o *Chatreaction) String() string {
    
     o.Users = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatreaction) MarshalJSON() ([]byte, error) {
    type Alias Chatreaction

    if ChatreactionMarshalled {
        return []byte("{}"), nil
    }
    ChatreactionMarshalled = true

    return json.Marshal(&struct {
        
        Emoji string `json:"emoji"`
        
        Users []Addressableentityref `json:"users"`
        *Alias
    }{

        


        
        Users: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

