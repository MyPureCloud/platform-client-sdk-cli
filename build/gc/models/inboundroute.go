package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InboundrouteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InboundrouteDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Inboundroute
type Inboundroute struct { 
    


    // Name
    Name string `json:"name"`


    // Pattern - The search pattern that the mailbox name should match.
    Pattern string `json:"pattern"`


    // Queue - The queue to route the emails to.
    Queue Domainentityref `json:"queue"`


    // Priority - The priority to use for routing.
    Priority int `json:"priority"`


    // Skills - The skills to use for routing.
    Skills []Domainentityref `json:"skills"`


    // Language - The language to use for routing.
    Language Domainentityref `json:"language"`


    // FromName - The sender name to use for outgoing replies.
    FromName string `json:"fromName"`


    // FromEmail - The sender email to use for outgoing replies.
    FromEmail string `json:"fromEmail"`


    // Flow - The flow to use for processing the email.
    Flow Domainentityref `json:"flow"`


    // ReplyEmailAddress - The route to use for email replies.
    ReplyEmailAddress *Queueemailaddress `json:"replyEmailAddress"`


    // AutoBcc - The recipients that should be automatically blind copied on outbound emails associated with this InboundRoute.
    AutoBcc []Emailaddress `json:"autoBcc"`


    // SpamFlow - The flow to use for processing inbound emails that have been marked as spam.
    SpamFlow Domainentityref `json:"spamFlow"`


    // Signature - The configuration for the canned response signature that will be appended to outbound emails sent via this route
    Signature Signature `json:"signature"`


    // HistoryInclusion - The configuration to indicate how the history of a conversation has to be included in a draft
    HistoryInclusion string `json:"historyInclusion"`


    // AllowMultipleActions - Control if multiple actions are allowed on this route. When true the disconnect has to be done manually. When false a conversation will be disconnected by the system after every action
    AllowMultipleActions bool `json:"allowMultipleActions"`


    

}

// String returns a JSON representation of the model
func (o *Inboundroute) String() string {
    
    
    
    
     o.Skills = []Domainentityref{{}} 
    
    
    
    
    
     o.AutoBcc = []Emailaddress{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inboundroute) MarshalJSON() ([]byte, error) {
    type Alias Inboundroute

    if InboundrouteMarshalled {
        return []byte("{}"), nil
    }
    InboundrouteMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Pattern string `json:"pattern"`
        
        Queue Domainentityref `json:"queue"`
        
        Priority int `json:"priority"`
        
        Skills []Domainentityref `json:"skills"`
        
        Language Domainentityref `json:"language"`
        
        FromName string `json:"fromName"`
        
        FromEmail string `json:"fromEmail"`
        
        Flow Domainentityref `json:"flow"`
        
        ReplyEmailAddress *Queueemailaddress `json:"replyEmailAddress"`
        
        AutoBcc []Emailaddress `json:"autoBcc"`
        
        SpamFlow Domainentityref `json:"spamFlow"`
        
        Signature Signature `json:"signature"`
        
        HistoryInclusion string `json:"historyInclusion"`
        
        AllowMultipleActions bool `json:"allowMultipleActions"`
        *Alias
    }{

        


        


        


        


        


        
        Skills: []Domainentityref{{}},
        


        


        


        


        


        


        
        AutoBcc: []Emailaddress{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

