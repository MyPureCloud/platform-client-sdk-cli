package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentlessemailsendresponsedtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentlessemailsendresponsedtoDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentlessemailsendresponsedto
type Agentlessemailsendresponsedto struct { 
    


    // ConversationId - The identifier of the conversation.
    ConversationId string `json:"conversationId"`


    // SenderType - The identifier of the external participant of the given conversation.
    SenderType string `json:"senderType"`


    // FromAddress - The sender of the message.
    FromAddress Emailaddress `json:"fromAddress"`


    // ToAddresses - The recipient of the message. We currently support one recipient only.
    ToAddresses []Emailaddress `json:"toAddresses"`


    // ReplyToAddress - The address to use for reply.
    ReplyToAddress Emailaddress `json:"replyToAddress"`


    // Subject - The subject of the message.
    Subject string `json:"subject"`


    // DateCreated - The message creation timestamp. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    

}

// String returns a JSON representation of the model
func (o *Agentlessemailsendresponsedto) String() string {
    
    
    
     o.ToAddresses = []Emailaddress{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentlessemailsendresponsedto) MarshalJSON() ([]byte, error) {
    type Alias Agentlessemailsendresponsedto

    if AgentlessemailsendresponsedtoMarshalled {
        return []byte("{}"), nil
    }
    AgentlessemailsendresponsedtoMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        SenderType string `json:"senderType"`
        
        FromAddress Emailaddress `json:"fromAddress"`
        
        ToAddresses []Emailaddress `json:"toAddresses"`
        
        ReplyToAddress Emailaddress `json:"replyToAddress"`
        
        Subject string `json:"subject"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        
        ToAddresses: []Emailaddress{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

