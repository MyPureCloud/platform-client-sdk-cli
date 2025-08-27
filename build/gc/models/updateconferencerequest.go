package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateconferencerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateconferencerequestDud struct { 
    

}

// Updateconferencerequest
type Updateconferencerequest struct { 
    // ConversationIds - A list of conversations to merge into this conversation to create a conference.
    ConversationIds []string `json:"conversationIds"`

}

// String returns a JSON representation of the model
func (o *Updateconferencerequest) String() string {
     o.ConversationIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateconferencerequest) MarshalJSON() ([]byte, error) {
    type Alias Updateconferencerequest

    if UpdateconferencerequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateconferencerequestMarshalled = true

    return json.Marshal(&struct {
        
        ConversationIds []string `json:"conversationIds"`
        *Alias
    }{

        
        ConversationIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

