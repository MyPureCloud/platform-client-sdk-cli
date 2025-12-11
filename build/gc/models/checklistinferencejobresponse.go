package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistinferencejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistinferencejobresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Checklistinferencejobresponse
type Checklistinferencejobresponse struct { 
    // Id - ID of the inference job.
    Id string `json:"id"`


    // Status - Status of the checklist inference job.
    Status string `json:"status"`


    // VarError - Error information associated with a job in case of any errors.
    VarError Errorinfo `json:"error"`


    // AgentChecklistInfo - Agent checklist info.
    AgentChecklistInfo Agentchecklistinfo `json:"agentChecklistInfo"`


    // JobStartTime - Date when the inference job started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    JobStartTime time.Time `json:"jobStartTime"`


    // JobEndTime - Date when the inference job finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    JobEndTime time.Time `json:"jobEndTime"`


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
func (o *Checklistinferencejobresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistinferencejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Checklistinferencejobresponse

    if ChecklistinferencejobresponseMarshalled {
        return []byte("{}"), nil
    }
    ChecklistinferencejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        VarError Errorinfo `json:"error"`
        
        AgentChecklistInfo Agentchecklistinfo `json:"agentChecklistInfo"`
        
        JobStartTime time.Time `json:"jobStartTime"`
        
        JobEndTime time.Time `json:"jobEndTime"`
        
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

