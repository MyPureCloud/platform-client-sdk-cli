package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistactivationpayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistactivationpayloadDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Checklistactivationpayload
type Checklistactivationpayload struct { 
    // ActivationTriggerType - Trigger type that activated this checklist.
    ActivationTriggerType string `json:"activationTriggerType"`


    // IntentId - The intent ID if checklist was triggered by an intent.
    IntentId string `json:"intentId"`


    // IntentName - The intent name if checklist was triggered by an intent.
    IntentName string `json:"intentName"`


    // Language - Language associated with the checklist.
    Language string `json:"language"`


    // AgentId - Agent ID.
    AgentId string `json:"agentId"`


    // ParticipantId - Participant ID.
    ParticipantId string `json:"participantId"`


    // QueueId - Queue ID.
    QueueId string `json:"queueId"`


    // AssistantId - Assistant ID.
    AssistantId string `json:"assistantId"`


    // MediaType - Media type.
    MediaType string `json:"mediaType"`


    // Direction - Direction of the conversation.
    Direction string `json:"direction"`

}

// String returns a JSON representation of the model
func (o *Checklistactivationpayload) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistactivationpayload) MarshalJSON() ([]byte, error) {
    type Alias Checklistactivationpayload

    if ChecklistactivationpayloadMarshalled {
        return []byte("{}"), nil
    }
    ChecklistactivationpayloadMarshalled = true

    return json.Marshal(&struct {
        
        ActivationTriggerType string `json:"activationTriggerType"`
        
        IntentId string `json:"intentId"`
        
        IntentName string `json:"intentName"`
        
        Language string `json:"language"`
        
        AgentId string `json:"agentId"`
        
        ParticipantId string `json:"participantId"`
        
        QueueId string `json:"queueId"`
        
        AssistantId string `json:"assistantId"`
        
        MediaType string `json:"mediaType"`
        
        Direction string `json:"direction"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

