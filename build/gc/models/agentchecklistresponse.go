package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentchecklistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentchecklistresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentchecklistresponse
type Agentchecklistresponse struct { 
    // Id - ID of the checklist.
    Id string `json:"id"`


    // Name - Name of the checklist.
    Name string `json:"name"`


    // ChecklistItems - List of individual Checklist Items.
    ChecklistItems []Checklistitem `json:"checklistItems"`


    // ActivationTriggers - List of triggers that activated this checklist.
    ActivationTriggers []Activationtrigger `json:"activationTriggers"`


    // Status - The evaluation status of the checklist.
    Status string `json:"status"`


    // ExitReason - Exit reason provided at the time of finalizing the checklist.
    ExitReason string `json:"exitReason"`


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


    // EvaluationStartDate - Date when the checklist evaluation began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EvaluationStartDate time.Time `json:"evaluationStartDate"`


    // EvaluationLastModifiedDate - Date when the checklist was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EvaluationLastModifiedDate time.Time `json:"evaluationLastModifiedDate"`


    // EvaluationFinalizedDate - Date when the checklist was finalized. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EvaluationFinalizedDate time.Time `json:"evaluationFinalizedDate"`


    // EvaluationFinalizedWithAcwDate - Date when the checklist was finalized with ACW. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EvaluationFinalizedWithAcwDate time.Time `json:"evaluationFinalizedWithAcwDate"`


    

}

// String returns a JSON representation of the model
func (o *Agentchecklistresponse) String() string {
    
    
     o.ChecklistItems = []Checklistitem{{}} 
     o.ActivationTriggers = []Activationtrigger{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentchecklistresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentchecklistresponse

    if AgentchecklistresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentchecklistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ChecklistItems []Checklistitem `json:"checklistItems"`
        
        ActivationTriggers []Activationtrigger `json:"activationTriggers"`
        
        Status string `json:"status"`
        
        ExitReason string `json:"exitReason"`
        
        Language string `json:"language"`
        
        AgentId string `json:"agentId"`
        
        ParticipantId string `json:"participantId"`
        
        QueueId string `json:"queueId"`
        
        AssistantId string `json:"assistantId"`
        
        MediaType string `json:"mediaType"`
        
        Direction string `json:"direction"`
        
        EvaluationStartDate time.Time `json:"evaluationStartDate"`
        
        EvaluationLastModifiedDate time.Time `json:"evaluationLastModifiedDate"`
        
        EvaluationFinalizedDate time.Time `json:"evaluationFinalizedDate"`
        
        EvaluationFinalizedWithAcwDate time.Time `json:"evaluationFinalizedWithAcwDate"`
        *Alias
    }{

        


        


        
        ChecklistItems: []Checklistitem{{}},
        


        
        ActivationTriggers: []Activationtrigger{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

