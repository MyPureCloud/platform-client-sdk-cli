package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationeventpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationeventpresenceDud struct { 
    


    


    

}

// Conversationeventpresence - A Presence event.
type Conversationeventpresence struct { 
    // VarType - Describes the type of Presence event.
    VarType string `json:"type"`


    // ReceivedMessage - A sign in prompt message for when the authenticate event is received.
    ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`


    // ReplyMessage - A reply message on authentication success.
    ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`

}

// String returns a JSON representation of the model
func (o *Conversationeventpresence) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationeventpresence) MarshalJSON() ([]byte, error) {
    type Alias Conversationeventpresence

    if ConversationeventpresenceMarshalled {
        return []byte("{}"), nil
    }
    ConversationeventpresenceMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`
        
        ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

