package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MeteredassignmentbyagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MeteredassignmentbyagentDud struct { 
    


    


    


    


    


    


    


    


    

}

// Meteredassignmentbyagent
type Meteredassignmentbyagent struct { 
    // EvaluationContextId
    EvaluationContextId string `json:"evaluationContextId"`


    // Evaluators
    Evaluators []User `json:"evaluators"`


    // MaxNumberEvaluations
    MaxNumberEvaluations int `json:"maxNumberEvaluations"`


    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // TimeInterval
    TimeInterval Timeinterval `json:"timeInterval"`


    // TimeZone
    TimeZone string `json:"timeZone"`


    // Percentage - Percentage of eligible conversations to sample for evaluation, from 0.00 to 100.00. If 0.00 or not set, percentage sampling is not used and maxNumberEvaluations applies.
    Percentage float32 `json:"percentage"`


    // Minimum - Minimum number of evaluations to create when percentage sampling is used.
    Minimum int `json:"minimum"`


    // AgentConnectedTimeCriteria - Only create evaluations for conversations where the agent's connected time satisfies this duration condition. Connected time runs from when the agent connected until the conversation ended, excluding alerting, dialing and wrap-up.
    AgentConnectedTimeCriteria Durationcondition `json:"agentConnectedTimeCriteria"`

}

// String returns a JSON representation of the model
func (o *Meteredassignmentbyagent) String() string {
    
     o.Evaluators = []User{{}} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Meteredassignmentbyagent) MarshalJSON() ([]byte, error) {
    type Alias Meteredassignmentbyagent

    if MeteredassignmentbyagentMarshalled {
        return []byte("{}"), nil
    }
    MeteredassignmentbyagentMarshalled = true

    return json.Marshal(&struct {
        
        EvaluationContextId string `json:"evaluationContextId"`
        
        Evaluators []User `json:"evaluators"`
        
        MaxNumberEvaluations int `json:"maxNumberEvaluations"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        TimeInterval Timeinterval `json:"timeInterval"`
        
        TimeZone string `json:"timeZone"`
        
        Percentage float32 `json:"percentage"`
        
        Minimum int `json:"minimum"`
        
        AgentConnectedTimeCriteria Durationcondition `json:"agentConnectedTimeCriteria"`
        *Alias
    }{

        


        
        Evaluators: []User{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

