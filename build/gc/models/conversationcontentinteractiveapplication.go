package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentinteractiveapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentinteractiveapplicationDud struct { 
    


    


    


    

}

// Conversationcontentinteractiveapplication - InteractiveApplication content object.
type Conversationcontentinteractiveapplication struct { 
    // Name - The name of the message app.
    Name string `json:"name"`


    // Url - Contains the data that is sent to the message app.
    Url string `json:"url"`


    // ReceivedMessage - The message displayed in the received message bubble.
    ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`


    // ReplyMessage - The message displayed in the reply message bubble.
    ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentinteractiveapplication) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentinteractiveapplication) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentinteractiveapplication

    if ConversationcontentinteractiveapplicationMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentinteractiveapplicationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        
        ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`
        
        ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

