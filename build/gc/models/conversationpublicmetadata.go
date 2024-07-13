package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationpublicmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationpublicmetadataDud struct { 
    


    

}

// Conversationpublicmetadata - Information about a public message.
type Conversationpublicmetadata struct { 
    // RootId - The id of the root public message.
    RootId string `json:"rootId"`


    // ReplyToId - The id of the message this public message is replying to.
    ReplyToId string `json:"replyToId"`

}

// String returns a JSON representation of the model
func (o *Conversationpublicmetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationpublicmetadata) MarshalJSON() ([]byte, error) {
    type Alias Conversationpublicmetadata

    if ConversationpublicmetadataMarshalled {
        return []byte("{}"), nil
    }
    ConversationpublicmetadataMarshalled = true

    return json.Marshal(&struct {
        
        RootId string `json:"rootId"`
        
        ReplyToId string `json:"replyToId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

