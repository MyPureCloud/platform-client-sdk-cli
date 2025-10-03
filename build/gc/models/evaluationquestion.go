package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationquestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationquestionDud struct { 
    


    ContextId string `json:"contextId"`


    


    


    


    


    


    


    


    


    


    

}

// Evaluationquestion
type Evaluationquestion struct { 
    // Id
    Id string `json:"id"`


    


    // Text
    Text string `json:"text"`


    // HelpText
    HelpText string `json:"helpText"`


    // VarType
    VarType string `json:"type"`


    // NaEnabled
    NaEnabled bool `json:"naEnabled"`


    // CommentsRequired
    CommentsRequired bool `json:"commentsRequired"`


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`


    // AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
    AnswerOptions []Answeroption `json:"answerOptions"`


    // MultipleSelectOptionQuestions - Only used by Multiple Select type questions. A list of multiple choice questions representing selectable options.
    MultipleSelectOptionQuestions []Evaluationquestion `json:"multipleSelectOptionQuestions"`


    // IsCritical
    IsCritical bool `json:"isCritical"`


    // IsKill
    IsKill bool `json:"isKill"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestion) String() string {
    
    
    
    
    
    
    
     o.AnswerOptions = []Answeroption{{}} 
     o.MultipleSelectOptionQuestions = []Evaluationquestion{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationquestion) MarshalJSON() ([]byte, error) {
    type Alias Evaluationquestion

    if EvaluationquestionMarshalled {
        return []byte("{}"), nil
    }
    EvaluationquestionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        HelpText string `json:"helpText"`
        
        VarType string `json:"type"`
        
        NaEnabled bool `json:"naEnabled"`
        
        CommentsRequired bool `json:"commentsRequired"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        
        AnswerOptions []Answeroption `json:"answerOptions"`
        
        MultipleSelectOptionQuestions []Evaluationquestion `json:"multipleSelectOptionQuestions"`
        
        IsCritical bool `json:"isCritical"`
        
        IsKill bool `json:"isKill"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        AnswerOptions: []Answeroption{{}},
        


        
        MultipleSelectOptionQuestions: []Evaluationquestion{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

