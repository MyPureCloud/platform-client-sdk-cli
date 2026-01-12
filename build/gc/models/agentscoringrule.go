package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentscoringruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentscoringruleDud struct { 
    Id string `json:"id"`


    ProgramId string `json:"programId"`


    


    


    


    


    


    Published bool `json:"published"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Agentscoringrule
type Agentscoringrule struct { 
    


    


    // SamplingType - Sampling type setting. Valid values: All, Percentage
    SamplingType string `json:"samplingType"`


    // SamplingPercentage - Percentage of interactions to evaluate (0.01-100.00). Required when samplingType is Percentage, null when All.
    SamplingPercentage float32 `json:"samplingPercentage"`


    // SubmissionType - Submission type for evaluations. Valid values: Automated, Manual
    SubmissionType string `json:"submissionType"`


    // EvaluationFormContextId - The evaluation form contextID to use for scoring.
    EvaluationFormContextId string `json:"evaluationFormContextId"`


    // Enabled - Whether the rule is enabled.
    Enabled bool `json:"enabled"`


    


    // Evaluator - The evaluator for evaluations created by this rule.
    Evaluator Addressableentityref `json:"evaluator"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Agentscoringrule) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentscoringrule) MarshalJSON() ([]byte, error) {
    type Alias Agentscoringrule

    if AgentscoringruleMarshalled {
        return []byte("{}"), nil
    }
    AgentscoringruleMarshalled = true

    return json.Marshal(&struct {
        
        SamplingType string `json:"samplingType"`
        
        SamplingPercentage float32 `json:"samplingPercentage"`
        
        SubmissionType string `json:"submissionType"`
        
        EvaluationFormContextId string `json:"evaluationFormContextId"`
        
        Enabled bool `json:"enabled"`
        
        Evaluator Addressableentityref `json:"evaluator"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

