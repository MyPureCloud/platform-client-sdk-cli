package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentformquestiongroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentformquestiongroupDud struct { 
    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Assessmentformquestiongroup
type Assessmentformquestiongroup struct { 
    // Id - The ID of the question group,
    Id string `json:"id"`


    // Name - The question group name
    Name string `json:"name"`


    // VarType - The question group type
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


    // Questions - The list of questions for this question group
    Questions []Assessmentformquestion `json:"questions"`


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`


    

}

// String returns a JSON representation of the model
func (o *Assessmentformquestiongroup) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Questions = []Assessmentformquestion{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentformquestiongroup) MarshalJSON() ([]byte, error) {
    type Alias Assessmentformquestiongroup

    if AssessmentformquestiongroupMarshalled {
        return []byte("{}"), nil
    }
    AssessmentformquestiongroupMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        DefaultAnswersToHighest bool `json:"defaultAnswersToHighest"`
        
        DefaultAnswersToNA bool `json:"defaultAnswersToNA"`
        
        NaEnabled bool `json:"naEnabled"`
        
        Weight float32 `json:"weight"`
        
        ManualWeight bool `json:"manualWeight"`
        
        Questions []Assessmentformquestion `json:"questions"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Questions: []Assessmentformquestion{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

