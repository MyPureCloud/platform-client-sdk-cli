package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigninteractionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigninteractionDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Campaigninteraction
type Campaigninteraction struct { 
    // Id
    Id string `json:"id"`


    // Campaign
    Campaign Domainentityref `json:"campaign"`


    // Agent
    Agent Domainentityref `json:"agent"`


    // Contact
    Contact Domainentityref `json:"contact"`


    // DestinationAddress
    DestinationAddress string `json:"destinationAddress"`


    // ActivePreviewCall - Boolean value if there is an active preview call on the interaction
    ActivePreviewCall bool `json:"activePreviewCall"`


    // LastActivePreviewWrapupTime - The time when the last preview of the interaction was wrapped up. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    LastActivePreviewWrapupTime time.Time `json:"lastActivePreviewWrapupTime"`


    // CreationTime - The time when dialer created the interaction. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreationTime time.Time `json:"creationTime"`


    // CallPlacedTime - The time when the agent or system places the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CallPlacedTime time.Time `json:"callPlacedTime"`


    // CallRoutedTime - The time when the agent was connected to the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CallRoutedTime time.Time `json:"callRoutedTime"`


    // PreviewConnectedTime - The time when the customer and routing participant are connected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PreviewConnectedTime time.Time `json:"previewConnectedTime"`


    // Queue
    Queue Domainentityref `json:"queue"`


    // Script
    Script Domainentityref `json:"script"`


    // Disposition - Describes what happened with call analysis for instance: disposition.classification.callable.person, disposition.classification.callable.noanswer
    Disposition string `json:"disposition"`


    // CallerName
    CallerName string `json:"callerName"`


    // CallerAddress
    CallerAddress string `json:"callerAddress"`


    // PreviewPopDeliveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PreviewPopDeliveredTime time.Time `json:"previewPopDeliveredTime"`


    // Conversation
    Conversation Conversationbasic `json:"conversation"`


    // DialerSystemParticipantId - conversation participant id that is the dialer system participant to monitor the call from dialer perspective
    DialerSystemParticipantId string `json:"dialerSystemParticipantId"`


    // DialingMode
    DialingMode string `json:"dialingMode"`


    // Skills - Any skills that are attached to the call for routing
    Skills []Domainentityref `json:"skills"`

}

// String returns a JSON representation of the model
func (o *Campaigninteraction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Skills = []Domainentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigninteraction) MarshalJSON() ([]byte, error) {
    type Alias Campaigninteraction

    if CampaigninteractionMarshalled {
        return []byte("{}"), nil
    }
    CampaigninteractionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Campaign Domainentityref `json:"campaign"`
        
        Agent Domainentityref `json:"agent"`
        
        Contact Domainentityref `json:"contact"`
        
        DestinationAddress string `json:"destinationAddress"`
        
        ActivePreviewCall bool `json:"activePreviewCall"`
        
        LastActivePreviewWrapupTime time.Time `json:"lastActivePreviewWrapupTime"`
        
        CreationTime time.Time `json:"creationTime"`
        
        CallPlacedTime time.Time `json:"callPlacedTime"`
        
        CallRoutedTime time.Time `json:"callRoutedTime"`
        
        PreviewConnectedTime time.Time `json:"previewConnectedTime"`
        
        Queue Domainentityref `json:"queue"`
        
        Script Domainentityref `json:"script"`
        
        Disposition string `json:"disposition"`
        
        CallerName string `json:"callerName"`
        
        CallerAddress string `json:"callerAddress"`
        
        PreviewPopDeliveredTime time.Time `json:"previewPopDeliveredTime"`
        
        Conversation Conversationbasic `json:"conversation"`
        
        DialerSystemParticipantId string `json:"dialerSystemParticipantId"`
        
        DialingMode string `json:"dialingMode"`
        
        Skills []Domainentityref `json:"skills"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Skills: []Domainentityref{{}},
        

        
        Alias: (*Alias)(u),
    })
}

