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
    // QueueId - The ID of the queue to be associated with the message. This will determine the fromAddress of the message, unless useUserFromAddress is true and the queue is configured to use the agent's Direct Routing address as the fromAddress.
    QueueId string `json:"queueId"`


    // ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.  For open messenger type, any string within the outbound.open.messaging.to.address.characters.max limit can be used. For whatsapp messenger type, use a Whatsapp ID of a phone number. E.g for a E.164 formatted phone number `+13175555555`, a Whatsapp ID would be 13175555555
    ToAddress string `json:"toAddress"`


    // ToAddressMessengerType - The messaging address messenger type.
    ToAddressMessengerType string `json:"toAddressMessengerType"`


    // UseExistingConversation - An override to use an existing conversation.  If set to true, an existing conversation will be used if there is one within the conversation window.  If set to false, create request fails if there is a conversation within the conversation window.
    UseExistingConversation bool `json:"useExistingConversation"`


    // ExternalContactId - The external contact with which the message will be associated.
    ExternalContactId string `json:"externalContactId"`


    // UseUserFromAddress - An override to attempt to use the user's configured direct routing address as the fromAddress.  If set to true, users configured address with 'directrouting' integration will be used as fromAddress.  If set to false or not set, the queueId will be used for determining fromAddress.
    UseUserFromAddress bool `json:"useUserFromAddress"`

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
        
        UseUserFromAddress bool `json:"useUserFromAddress"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

