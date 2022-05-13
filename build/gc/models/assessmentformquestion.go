package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentformquestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentformquestionDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Assessmentformquestion
type Assessmentformquestion struct { 
    // Id
    Id string `json:"id"`


    // VarType
    VarType string `json:"type"`


    // Text - The question text
    Text string `json:"text"`


    // HelpText
    HelpText string `json:"helpText"`


    // NaEnabled
    NaEnabled bool `json:"naEnabled"`


    // CommentsRequired
    CommentsRequired bool `json:"commentsRequired"`


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`


    // AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
    AnswerOptions []Answeroption `json:"answerOptions"`


    // MaxResponseCharacters - How many characters are allowed in the text response to this question. Used by Free Text question types.
    MaxResponseCharacters int `json:"maxResponseCharacters"`


    // IsKill - Does an incorrect answer to this question mark the form as having a failed kill question. Only used by Multiple Choice type questions.
    IsKill bool `json:"isKill"`


    // IsCritical - Does this question contribute to the critical score. Only used by Multiple Choice type questions.
    IsCritical bool `json:"isCritical"`

}

// String returns a JSON representation of the model
func (o *Assessmentformquestion) String() string {
    
    
    
    
    
    
    
     o.AnswerOptions = []Answeroption{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentformquestion) MarshalJSON() ([]byte, error) {
    type Alias Assessmentformquestion

    if AssessmentformquestionMarshalled {
        return []byte("{}"), nil
    }
    AssessmentformquestionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        HelpText string `json:"helpText"`
        
        NaEnabled bool `json:"naEnabled"`
        
        CommentsRequired bool `json:"commentsRequired"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        
        AnswerOptions []Answeroption `json:"answerOptions"`
        
        MaxResponseCharacters int `json:"maxResponseCharacters"`
        
        IsKill bool `json:"isKill"`
        
        IsCritical bool `json:"isCritical"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        AnswerOptions: []Answeroption{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

