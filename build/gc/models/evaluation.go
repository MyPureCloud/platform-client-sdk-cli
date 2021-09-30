package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Evaluation
type Evaluation struct { 
    


    // Name
    Name string `json:"name"`


    // Conversation
    Conversation Conversation `json:"conversation"`


    // EvaluationForm - Evaluation form used for evaluation.
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // Evaluator
    Evaluator User `json:"evaluator"`


    // Agent
    Agent User `json:"agent"`


    // Calibration
    Calibration Calibration `json:"calibration"`


    // Status
    Status string `json:"status"`


    // Answers
    Answers Evaluationscoringset `json:"answers"`


    // AgentHasRead
    AgentHasRead bool `json:"agentHasRead"`


    // ReleaseDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReleaseDate time.Time `json:"releaseDate"`


    // AssignedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AssignedDate time.Time `json:"assignedDate"`


    // ChangedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ChangedDate time.Time `json:"changedDate"`


    // Queue
    Queue Queue `json:"queue"`


    // MediaType - List of different communication types used in conversation.
    MediaType []string `json:"mediaType"`


    // Rescore - Is only true when evaluation is re-scored.
    Rescore bool `json:"rescore"`


    // ConversationDate - Date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationDate time.Time `json:"conversationDate"`


    // ConversationEndDate - End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationEndDate time.Time `json:"conversationEndDate"`


    // NeverRelease - Signifies if the evaluation is never to be released. This cannot be set true if release date is also set.
    NeverRelease bool `json:"neverRelease"`


    // ResourceId - Only used for email evaluations. Will be null for all other evaluations.
    ResourceId string `json:"resourceId"`


    // ResourceType - The type of resource. Only used for email evaluations. Will be null for evaluations on all other resources.
    ResourceType string `json:"resourceType"`


    // Redacted - Is only true when the user making the request does not have sufficient permissions to see evaluation
    Redacted bool `json:"redacted"`


    // IsScoringIndex
    IsScoringIndex bool `json:"isScoringIndex"`


    // AuthorizedActions - List of user authorized actions on evaluation. Possible values: edit, editScore, editAgentSignoff, delete, viewAudit
    AuthorizedActions []string `json:"authorizedActions"`


    

}

// String returns a JSON representation of the model
func (o *Evaluation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MediaType = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AuthorizedActions = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluation) MarshalJSON() ([]byte, error) {
    type Alias Evaluation

    if EvaluationMarshalled {
        return []byte("{}"), nil
    }
    EvaluationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Conversation Conversation `json:"conversation"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        Evaluator User `json:"evaluator"`
        
        Agent User `json:"agent"`
        
        Calibration Calibration `json:"calibration"`
        
        Status string `json:"status"`
        
        Answers Evaluationscoringset `json:"answers"`
        
        AgentHasRead bool `json:"agentHasRead"`
        
        ReleaseDate time.Time `json:"releaseDate"`
        
        AssignedDate time.Time `json:"assignedDate"`
        
        ChangedDate time.Time `json:"changedDate"`
        
        Queue Queue `json:"queue"`
        
        MediaType []string `json:"mediaType"`
        
        Rescore bool `json:"rescore"`
        
        ConversationDate time.Time `json:"conversationDate"`
        
        ConversationEndDate time.Time `json:"conversationEndDate"`
        
        NeverRelease bool `json:"neverRelease"`
        
        ResourceId string `json:"resourceId"`
        
        ResourceType string `json:"resourceType"`
        
        Redacted bool `json:"redacted"`
        
        IsScoringIndex bool `json:"isScoringIndex"`
        
        AuthorizedActions []string `json:"authorizedActions"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        MediaType: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        AuthorizedActions: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

