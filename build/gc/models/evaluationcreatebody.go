package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationcreatebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationcreatebodyDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Evaluationcreatebody
type Evaluationcreatebody struct { 
    


    // EvaluationForm - Evaluation form used for evaluation (must be included for a successful request)
    EvaluationForm Evaluationcreateevalform `json:"evaluationForm"`


    // Evaluator - User ID of the evaluator (must be included for a successful request)
    Evaluator Evaluationcreateuser `json:"evaluator"`


    // Agent - User ID of the agent (must be included for a successful request)
    Agent Evaluationcreateuser `json:"agent"`


    // AgentHasRead
    AgentHasRead bool `json:"agentHasRead"`


    // Answers
    Answers Evaluationscoringset `json:"answers"`


    // Calibration
    Calibration Evaluationcreatecalibration `json:"calibration"`


    // EvaluationContextId
    EvaluationContextId string `json:"evaluationContextId"`


    // Conversation
    Conversation Evaluationcreateconversation `json:"conversation"`


    // ResourceType
    ResourceType string `json:"resourceType"`


    // EvaluationSource
    EvaluationSource Evaluationsource `json:"evaluationSource"`


    // Rescore
    Rescore bool `json:"rescore"`


    // Queue
    Queue Evaluationcreatequeue `json:"queue"`


    // ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReleaseDate time.Time `json:"releaseDate"`


    // Status
    Status string `json:"status"`


    // NeverRelease
    NeverRelease bool `json:"neverRelease"`


    // DateAssigneeChanged - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssigneeChanged time.Time `json:"dateAssigneeChanged"`


    // Assignee
    Assignee Evaluationcreateuser `json:"assignee"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationcreatebody) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationcreatebody) MarshalJSON() ([]byte, error) {
    type Alias Evaluationcreatebody

    if EvaluationcreatebodyMarshalled {
        return []byte("{}"), nil
    }
    EvaluationcreatebodyMarshalled = true

    return json.Marshal(&struct {
        
        EvaluationForm Evaluationcreateevalform `json:"evaluationForm"`
        
        Evaluator Evaluationcreateuser `json:"evaluator"`
        
        Agent Evaluationcreateuser `json:"agent"`
        
        AgentHasRead bool `json:"agentHasRead"`
        
        Answers Evaluationscoringset `json:"answers"`
        
        Calibration Evaluationcreatecalibration `json:"calibration"`
        
        EvaluationContextId string `json:"evaluationContextId"`
        
        Conversation Evaluationcreateconversation `json:"conversation"`
        
        ResourceType string `json:"resourceType"`
        
        EvaluationSource Evaluationsource `json:"evaluationSource"`
        
        Rescore bool `json:"rescore"`
        
        Queue Evaluationcreatequeue `json:"queue"`
        
        ReleaseDate time.Time `json:"releaseDate"`
        
        Status string `json:"status"`
        
        NeverRelease bool `json:"neverRelease"`
        
        DateAssigneeChanged time.Time `json:"dateAssigneeChanged"`
        
        Assignee Evaluationcreateuser `json:"assignee"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

