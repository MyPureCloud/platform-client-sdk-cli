package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentlistpickerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentlistpickerDud struct { 
    


    


    

}

// Conversationcontentlistpicker - List Picker object for presenting multiple sections of selectable items.
type Conversationcontentlistpicker struct { 
    // Sections - An array of sections in the List Picker.
    Sections []Conversationcontentlistpickersection `json:"sections"`


    // ReplyMessage - The message displayed in the received message bubble.
    ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`


    // ReceivedMessage - The message displayed in the reply message bubble.
    ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentlistpicker) String() string {
     o.Sections = []Conversationcontentlistpickersection{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentlistpicker) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentlistpicker

    if ConversationcontentlistpickerMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentlistpickerMarshalled = true

    return json.Marshal(&struct {
        
        Sections []Conversationcontentlistpickersection `json:"sections"`
        
        ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`
        
        ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`
        *Alias
    }{

        
        Sections: []Conversationcontentlistpickersection{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

