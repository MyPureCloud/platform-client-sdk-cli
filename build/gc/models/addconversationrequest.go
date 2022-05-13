package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddconversationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddconversationrequestDud struct { 
    

}

// Addconversationrequest - Update coaching appointment request
type Addconversationrequest struct { 
    // ConversationId - The id of the conversation to add
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Addconversationrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addconversationrequest) MarshalJSON() ([]byte, error) {
    type Alias Addconversationrequest

    if AddconversationrequestMarshalled {
        return []byte("{}"), nil
    }
    AddconversationrequestMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

