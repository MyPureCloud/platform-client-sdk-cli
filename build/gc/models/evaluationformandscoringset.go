package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationformandscoringsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationformandscoringsetDud struct { 
    


    

}

// Evaluationformandscoringset
type Evaluationformandscoringset struct { 
    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // Answers
    Answers Evaluationscoringset `json:"answers"`

}

// String returns a JSON representation of the model
func (o *Evaluationformandscoringset) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationformandscoringset) MarshalJSON() ([]byte, error) {
    type Alias Evaluationformandscoringset

    if EvaluationformandscoringsetMarshalled {
        return []byte("{}"), nil
    }
    EvaluationformandscoringsetMarshalled = true

    return json.Marshal(&struct { 
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        Answers Evaluationscoringset `json:"answers"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

