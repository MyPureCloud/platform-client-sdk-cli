package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationreasonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationreasonDud struct { 
    


    

}

// Conversationreason - Reasons for a failed message receipt.
type Conversationreason struct { 
    // Code - The reason code for the failed message receipt.
    Code string `json:"code"`


    // Message - Description of the reason for the failed message receipt.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Conversationreason) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationreason) MarshalJSON() ([]byte, error) {
    type Alias Conversationreason

    if ConversationreasonMarshalled {
        return []byte("{}"), nil
    }
    ConversationreasonMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

