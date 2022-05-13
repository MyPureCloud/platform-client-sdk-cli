package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateoutboundmessagingconversationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateoutboundmessagingconversationrequestDud struct { 
    


    


    


    


    

}

// Createoutboundmessagingconversationrequest
type Createoutboundmessagingconversationrequest struct { 
    // QueueId - The ID of the queue to be associated with the message. This will determine the fromAddress of the message.
    QueueId string `json:"queueId"`


    // ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234
    ToAddress string `json:"toAddress"`


    // ToAddressMessengerType - The messaging address messenger type.
    ToAddressMessengerType string `json:"toAddressMessengerType"`


    // UseExistingConversation - An override to use an existing conversation.  If set to true, an existing conversation will be used if there is one within the conversation window.  If set to false, create request fails if there is a conversation within the conversation window.
    UseExistingConversation bool `json:"useExistingConversation"`


    // ExternalContactId - The external contact with which the message will be associated.
    ExternalContactId string `json:"externalContactId"`

}

// String returns a JSON representation of the model
func (o *Createoutboundmessagingconversationrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createoutboundmessagingconversationrequest) MarshalJSON() ([]byte, error) {
    type Alias Createoutboundmessagingconversationrequest

    if CreateoutboundmessagingconversationrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateoutboundmessagingconversationrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        ToAddress string `json:"toAddress"`
        
        ToAddressMessengerType string `json:"toAddressMessengerType"`
        
        UseExistingConversation bool `json:"useExistingConversation"`
        
        ExternalContactId string `json:"externalContactId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

