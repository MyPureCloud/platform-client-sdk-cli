package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationquestiongroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationquestiongroupDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Evaluationquestiongroup
type Evaluationquestiongroup struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // DefaultAnswersToHighest
    DefaultAnswersToHighest bool `json:"defaultAnswersToHighest"`


    // DefaultAnswersToNA
    DefaultAnswersToNA bool `json:"defaultAnswersToNA"`


    // NaEnabled
    NaEnabled bool `json:"naEnabled"`


    // Weight
    Weight float32 `json:"weight"`


    // ManualWeight
    ManualWeight bool `json:"manualWeight"`


    // Questions
    Questions []Evaluationquestion `json:"questions"`


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroup) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Questions = []Evaluationquestion{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationquestiongroup) MarshalJSON() ([]byte, error) {
    type Alias Evaluationquestiongroup

    if EvaluationquestiongroupMarshalled {
        return []byte("{}"), nil
    }
    EvaluationquestiongroupMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        DefaultAnswersToHighest bool `json:"defaultAnswersToHighest"`
        
        DefaultAnswersToNA bool `json:"defaultAnswersToNA"`
        
        NaEnabled bool `json:"naEnabled"`
        
        Weight float32 `json:"weight"`
        
        ManualWeight bool `json:"manualWeight"`
        
        Questions []Evaluationquestion `json:"questions"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Questions: []Evaluationquestion{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

