package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsettingsDud struct { 
    


    


    


    

}

// Evaluationsettings
type Evaluationsettings struct { 
    // RevisionsEnabled - Whether revisions are allowed for evaluations. When enabled, rescoring creates a new version of the evaluation and retracts the existing evaluation version. Does not apply for calibration evaluations.
    RevisionsEnabled bool `json:"revisionsEnabled"`


    // DisputesEnabled - Whether disputes are allowed for evaluations. Does not apply for calibration evaluations.
    DisputesEnabled bool `json:"disputesEnabled"`


    // DisputesAllowedPerEvaluation - The maximum number of disputes allowed for an evaluation.
    DisputesAllowedPerEvaluation int `json:"disputesAllowedPerEvaluation"`


    // DisputesAssignees - A list of assignees responsible for handling each dispute. This list size needs to be equal to disputesAllowedPerEvaluation.
    DisputesAssignees []Evaluationsettingsassignee `json:"disputesAssignees"`

}

// String returns a JSON representation of the model
func (o *Evaluationsettings) String() string {
    
    
    
     o.DisputesAssignees = []Evaluationsettingsassignee{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsettings) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsettings

    if EvaluationsettingsMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsettingsMarshalled = true

    return json.Marshal(&struct {
        
        RevisionsEnabled bool `json:"revisionsEnabled"`
        
        DisputesEnabled bool `json:"disputesEnabled"`
        
        DisputesAllowedPerEvaluation int `json:"disputesAllowedPerEvaluation"`
        
        DisputesAssignees []Evaluationsettingsassignee `json:"disputesAssignees"`
        *Alias
    }{

        


        


        


        
        DisputesAssignees: []Evaluationsettingsassignee{{}},
        

        Alias: (*Alias)(u),
    })
}

