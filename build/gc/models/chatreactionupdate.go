package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatreactionupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatreactionupdateDud struct { 
    

}

// Chatreactionupdate
type Chatreactionupdate struct { 
    // Reactions - Reactions to update
    Reactions []string `json:"reactions"`

}

// String returns a JSON representation of the model
func (o *Chatreactionupdate) String() string {
     o.Reactions = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatreactionupdate) MarshalJSON() ([]byte, error) {
    type Alias Chatreactionupdate

    if ChatreactionupdateMarshalled {
        return []byte("{}"), nil
    }
    ChatreactionupdateMarshalled = true

    return json.Marshal(&struct {
        
        Reactions []string `json:"reactions"`
        *Alias
    }{

        
        Reactions: []string{""},
        

        Alias: (*Alias)(u),
    })
}

