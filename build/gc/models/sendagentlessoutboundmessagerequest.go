package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SendagentlessoutboundmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SendagentlessoutboundmessagerequestDud struct { 
    


    


    


    


    


    

}

// Sendagentlessoutboundmessagerequest
type Sendagentlessoutboundmessagerequest struct { 
    // FromAddress - The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned SMS phone number. For a WhatsApp messenger type use the provisioned WhatsApp integration’s ID
    FromAddress string `json:"fromAddress"`


    // ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.
    ToAddress string `json:"toAddress"`


    // ToAddressMessengerType - The recipient messaging address messenger type.
    ToAddressMessengerType string `json:"toAddressMessengerType"`


    // TextBody - The text of the message to send. This field is required in the case of SMS messenger type. Maximum character counts are: SMS - 765 characters, other channels - 2000 characters.
    TextBody string `json:"textBody"`


    // MessagingTemplate - The messaging template to use in the case of WhatsApp messenger type. This field is required when using WhatsApp messenger type
    MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`


    // UseExistingActiveConversation - Use an existing active conversation to send the agentless outbound message. Set this parameter to 'true' to use active conversation. Default value: false
    UseExistingActiveConversation bool `json:"useExistingActiveConversation"`

}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessagerequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sendagentlessoutboundmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Sendagentlessoutboundmessagerequest

    if SendagentlessoutboundmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    SendagentlessoutboundmessagerequestMarshalled = true

    return json.Marshal(&struct {
        
        FromAddress string `json:"fromAddress"`
        
        ToAddress string `json:"toAddress"`
        
        ToAddressMessengerType string `json:"toAddressMessengerType"`
        
        TextBody string `json:"textBody"`
        
        MessagingTemplate Messagingtemplaterequest `json:"messagingTemplate"`
        
        UseExistingActiveConversation bool `json:"useExistingActiveConversation"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

