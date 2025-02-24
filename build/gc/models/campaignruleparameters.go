package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleparametersDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Campaignruleparameters
type Campaignruleparameters struct { 
    // Operator - The operator for comparison. Required for a CampaignRuleCondition.
    Operator string `json:"operator"`


    // Value - The value for comparison. Required for a CampaignRuleCondition.
    Value string `json:"value"`


    // Priority - The priority to set a campaign to. Required for the 'setCampaignPriority' action.
    Priority string `json:"priority"`


    // DialingMode - The dialing mode to set a campaign to. Required for the 'setCampaignDialingMode' action.
    DialingMode string `json:"dialingMode"`


    // AbandonRate - The abandon rate to set a campaign to. Required for the 'setCampaignAbandonRate' action.
    AbandonRate float32 `json:"abandonRate"`


    // OutboundLineCount - The  number of outbound lines to set a campaign to. Required for the 'setCampaignNumberOfLines' action.
    OutboundLineCount int `json:"outboundLineCount"`


    // RelativeWeight - The relative weight to set a campaign to. Required for the 'setCampaignWeight' action.
    RelativeWeight int `json:"relativeWeight"`


    // MaxCallsPerAgent - The maximum number of calls per agent to set a campaign to. Required for the 'setCampaignMaxCallsPerAgent' action.
    MaxCallsPerAgent float32 `json:"maxCallsPerAgent"`


    // Queue - The queue a campaign to. Required for the 'changeCampaignQueue' action.
    Queue Domainentityref `json:"queue"`


    // MessagesPerMinute - The number of messages per minute to set a messaging campaign to.
    MessagesPerMinute int `json:"messagesPerMinute"`


    // SmsMessagesPerMinute - The number of messages per minute to set a SMS messaging campaign to.
    SmsMessagesPerMinute int `json:"smsMessagesPerMinute"`


    // EmailMessagesPerMinute - The number of messages per minute to set a Email messaging campaign to.
    EmailMessagesPerMinute int `json:"emailMessagesPerMinute"`


    // SmsContentTemplate - The content template to set a SMS campaign to.
    SmsContentTemplate Domainentityref `json:"smsContentTemplate"`


    // EmailContentTemplate - The content template to set a Email campaign to.
    EmailContentTemplate Domainentityref `json:"emailContentTemplate"`

}

// String returns a JSON representation of the model
func (o *Campaignruleparameters) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleparameters

    if CampaignruleparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleparametersMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        Priority string `json:"priority"`
        
        DialingMode string `json:"dialingMode"`
        
        AbandonRate float32 `json:"abandonRate"`
        
        OutboundLineCount int `json:"outboundLineCount"`
        
        RelativeWeight int `json:"relativeWeight"`
        
        MaxCallsPerAgent float32 `json:"maxCallsPerAgent"`
        
        Queue Domainentityref `json:"queue"`
        
        MessagesPerMinute int `json:"messagesPerMinute"`
        
        SmsMessagesPerMinute int `json:"smsMessagesPerMinute"`
        
        EmailMessagesPerMinute int `json:"emailMessagesPerMinute"`
        
        SmsContentTemplate Domainentityref `json:"smsContentTemplate"`
        
        EmailContentTemplate Domainentityref `json:"emailContentTemplate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

