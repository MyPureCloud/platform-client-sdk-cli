package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyquestiongroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyquestiongroupDud struct { 
    


    ContextId string `json:"contextId"`


    


    


    


    


    

}

// Surveyquestiongroup
type Surveyquestiongroup struct { 
    // Id
    Id string `json:"id"`


    


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // NaEnabled
    NaEnabled bool `json:"naEnabled"`


    // Questions
    Questions []Surveyquestion `json:"questions"`


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`

}

// String returns a JSON representation of the model
func (o *Surveyquestiongroup) String() string {
    
    
    
    
     o.Questions = []Surveyquestion{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyquestiongroup) MarshalJSON() ([]byte, error) {
    type Alias Surveyquestiongroup

    if SurveyquestiongroupMarshalled {
        return []byte("{}"), nil
    }
    SurveyquestiongroupMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        NaEnabled bool `json:"naEnabled"`
        
        Questions []Surveyquestion `json:"questions"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        *Alias
    }{

        


        


        


        


        


        
        Questions: []Surveyquestion{{}},
        


        

        Alias: (*Alias)(u),
    })
}

