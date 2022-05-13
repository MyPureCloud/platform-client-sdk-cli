package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateemailrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateemailrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Createemailrequest
type Createemailrequest struct { 
    // QueueId - The ID of the queue to use for routing the email conversation. This field is mutually exclusive with flowId
    QueueId string `json:"queueId"`


    // FlowId - The ID of the flow to use for routing email conversation. This field is mutually exclusive with queueId
    FlowId string `json:"flowId"`


    // Provider - The name of the provider that is sourcing the emails. The Provider \"PureCloud Email\" is reserved for native emails.
    Provider string `json:"provider"`


    // SkillIds - The list of skill ID's to use for routing.
    SkillIds []string `json:"skillIds"`


    // LanguageId - The ID of the language to use for routing.
    LanguageId string `json:"languageId"`


    // Priority - The priority to assign to the conversation for routing.
    Priority int `json:"priority"`


    // Attributes - The list of attributes to associate with the customer participant.
    Attributes map[string]string `json:"attributes"`


    // ToAddress - The email address of the recipient of the email.
    ToAddress string `json:"toAddress"`


    // ToName - The name of the recipient of the email.
    ToName string `json:"toName"`


    // FromAddress - The email address of the sender of the email.
    FromAddress string `json:"fromAddress"`


    // FromName - The name of the sender of the email.
    FromName string `json:"fromName"`


    // Subject - The subject of the email
    Subject string `json:"subject"`


    // Direction - Specify OUTBOUND to send an email on behalf of a queue, or INBOUND to create an external conversation. An external conversation is one where the provider is not PureCloud based.
    Direction string `json:"direction"`


    // HtmlBody - An HTML body content of the email.
    HtmlBody string `json:"htmlBody"`


    // TextBody - A text body content of the email.
    TextBody string `json:"textBody"`


    // ExternalContactId - The external contact with which the email should be associated. This field is only valid for OUTBOUND email.
    ExternalContactId string `json:"externalContactId"`

}

// String returns a JSON representation of the model
func (o *Createemailrequest) String() string {
    
    
    
     o.SkillIds = []string{""} 
    
    
     o.Attributes = map[string]string{"": ""} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createemailrequest) MarshalJSON() ([]byte, error) {
    type Alias Createemailrequest

    if CreateemailrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateemailrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        FlowId string `json:"flowId"`
        
        Provider string `json:"provider"`
        
        SkillIds []string `json:"skillIds"`
        
        LanguageId string `json:"languageId"`
        
        Priority int `json:"priority"`
        
        Attributes map[string]string `json:"attributes"`
        
        ToAddress string `json:"toAddress"`
        
        ToName string `json:"toName"`
        
        FromAddress string `json:"fromAddress"`
        
        FromName string `json:"fromName"`
        
        Subject string `json:"subject"`
        
        Direction string `json:"direction"`
        
        HtmlBody string `json:"htmlBody"`
        
        TextBody string `json:"textBody"`
        
        ExternalContactId string `json:"externalContactId"`
        *Alias
    }{

        


        


        


        
        SkillIds: []string{""},
        


        


        


        
        Attributes: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

