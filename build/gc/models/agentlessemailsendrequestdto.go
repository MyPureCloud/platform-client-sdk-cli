package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentlessemailsendrequestdtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentlessemailsendrequestdtoDud struct { 
    


    


    


    


    


    


    


    

}

// Agentlessemailsendrequestdto
type Agentlessemailsendrequestdto struct { 
    // SenderType - The direction of the message.
    SenderType string `json:"senderType"`


    // ConversationId - The identifier of the conversation.
    ConversationId string `json:"conversationId"`


    // FromAddress - The sender of the message.
    FromAddress Emailaddress `json:"fromAddress"`


    // ToAddresses - The recipient(s) of the message.
    ToAddresses []Emailaddress `json:"toAddresses"`


    // ReplyToAddress - The address to use for reply.
    ReplyToAddress Emailaddress `json:"replyToAddress"`


    // Subject - The subject of the message.
    Subject string `json:"subject"`


    // TextBody - The Content of the message, in plain text.
    TextBody string `json:"textBody"`


    // HtmlBody - The Content of the message, in HTML. Links, images and styles are allowed
    HtmlBody string `json:"htmlBody"`

}

// String returns a JSON representation of the model
func (o *Agentlessemailsendrequestdto) String() string {
    
    
    
     o.ToAddresses = []Emailaddress{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentlessemailsendrequestdto) MarshalJSON() ([]byte, error) {
    type Alias Agentlessemailsendrequestdto

    if AgentlessemailsendrequestdtoMarshalled {
        return []byte("{}"), nil
    }
    AgentlessemailsendrequestdtoMarshalled = true

    return json.Marshal(&struct {
        
        SenderType string `json:"senderType"`
        
        ConversationId string `json:"conversationId"`
        
        FromAddress Emailaddress `json:"fromAddress"`
        
        ToAddresses []Emailaddress `json:"toAddresses"`
        
        ReplyToAddress Emailaddress `json:"replyToAddress"`
        
        Subject string `json:"subject"`
        
        TextBody string `json:"textBody"`
        
        HtmlBody string `json:"htmlBody"`
        *Alias
    }{

        


        


        


        
        ToAddresses: []Emailaddress{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

