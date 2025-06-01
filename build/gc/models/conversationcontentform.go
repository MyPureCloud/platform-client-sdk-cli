package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentformDud struct { 
    


    


    


    


    


    


    


    

}

// Conversationcontentform - Form content object.
type Conversationcontentform struct { 
    // Introduction - The intro component, used to give an intro into what the form entails
    Introduction Conversationcontentintroduction `json:"introduction"`


    // FormPages - Form pages
    FormPages []Conversationformpage `json:"formPages"`


    // ReceivedMessage - The message prompt to fill out the form received.
    ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`


    // ReplyMessage - The reply message after the user has filled out the form received.
    ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`


    // ShowSummary - Show summary at end of form submission.
    ShowSummary bool `json:"showSummary"`


    // Response - Content of the payload included in the Form response.
    Response []Conversationformresponsecomponent `json:"response"`


    // OriginatingMessageId - Reference to the ID of the original message.
    OriginatingMessageId string `json:"originatingMessageId"`


    // CannedResponseId - The id of the canned response which was used to create the form.
    CannedResponseId string `json:"cannedResponseId"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentform) String() string {
    
     o.FormPages = []Conversationformpage{{}} 
    
    
    
     o.Response = []Conversationformresponsecomponent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentform) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentform

    if ConversationcontentformMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentformMarshalled = true

    return json.Marshal(&struct {
        
        Introduction Conversationcontentintroduction `json:"introduction"`
        
        FormPages []Conversationformpage `json:"formPages"`
        
        ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`
        
        ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`
        
        ShowSummary bool `json:"showSummary"`
        
        Response []Conversationformresponsecomponent `json:"response"`
        
        OriginatingMessageId string `json:"originatingMessageId"`
        
        CannedResponseId string `json:"cannedResponseId"`
        *Alias
    }{

        


        
        FormPages: []Conversationformpage{{}},
        


        


        


        


        
        Response: []Conversationformresponsecomponent{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

