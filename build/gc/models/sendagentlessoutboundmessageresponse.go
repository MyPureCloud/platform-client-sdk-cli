package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SendagentlessoutboundmessageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SendagentlessoutboundmessageresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    User Addressableentityref `json:"user"`

}

// Sendagentlessoutboundmessageresponse
type Sendagentlessoutboundmessageresponse struct { 
    


    // ConversationId - The identifier of the conversation.
    ConversationId string `json:"conversationId"`


    // FromAddress - The sender of the message.
    FromAddress string `json:"fromAddress"`


    // ToAddress - The recipient of the message.
    ToAddress string `json:"toAddress"`


    // MessengerType - Type of messenger.
    MessengerType string `json:"messengerType"`


    // TextBody - The body of the text message.
    TextBody string `json:"textBody"`


    // MessagingTemplate - The messaging template sent
    MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`


    // Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    


    

}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessageresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sendagentlessoutboundmessageresponse) MarshalJSON() ([]byte, error) {
    type Alias Sendagentlessoutboundmessageresponse

    if SendagentlessoutboundmessageresponseMarshalled {
        return []byte("{}"), nil
    }
    SendagentlessoutboundmessageresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        ConversationId string `json:"conversationId"`
        
        FromAddress string `json:"fromAddress"`
        
        ToAddress string `json:"toAddress"`
        
        MessengerType string `json:"messengerType"`
        
        TextBody string `json:"textBody"`
        
        MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`
        
        Timestamp time.Time `json:"timestamp"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

