package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantmetricsDud struct { 
    


    


    


    


    


    


    


    

}

// Participantmetrics
type Participantmetrics struct { 
    // AgentDurationPercentage - Percentage of Agent duration in the conversation
    AgentDurationPercentage float64 `json:"agentDurationPercentage"`


    // CustomerDurationPercentage - Percentage of Customer duration in the conversation
    CustomerDurationPercentage float64 `json:"customerDurationPercentage"`


    // SilenceDurationPercentage - Percentage of Silence duration in the conversation
    SilenceDurationPercentage float64 `json:"silenceDurationPercentage"`


    // IvrDurationPercentage - Percentage of IVR duration in the conversation
    IvrDurationPercentage float64 `json:"ivrDurationPercentage"`


    // AcdDurationPercentage - Percentage of ACD duration in the conversation
    AcdDurationPercentage float64 `json:"acdDurationPercentage"`


    // OvertalkDurationPercentage - Percentage of Overtalk duration in the conversation
    OvertalkDurationPercentage float64 `json:"overtalkDurationPercentage"`


    // OtherDurationPercentage - Percentage of Other events duration in the conversation
    OtherDurationPercentage float64 `json:"otherDurationPercentage"`


    // OvertalkCount - Number of Overtalks in the conversation
    OvertalkCount int `json:"overtalkCount"`

}

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participantmetrics) MarshalJSON() ([]byte, error) {
    type Alias Participantmetrics

    if ParticipantmetricsMarshalled {
        return []byte("{}"), nil
    }
    ParticipantmetricsMarshalled = true

    return json.Marshal(&struct { 
        AgentDurationPercentage float64 `json:"agentDurationPercentage"`
        
        CustomerDurationPercentage float64 `json:"customerDurationPercentage"`
        
        SilenceDurationPercentage float64 `json:"silenceDurationPercentage"`
        
        IvrDurationPercentage float64 `json:"ivrDurationPercentage"`
        
        AcdDurationPercentage float64 `json:"acdDurationPercentage"`
        
        OvertalkDurationPercentage float64 `json:"overtalkDurationPercentage"`
        
        OtherDurationPercentage float64 `json:"otherDurationPercentage"`
        
        OvertalkCount int `json:"overtalkCount"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

