package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyquestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyquestionDud struct { 
    


    ContextId string `json:"contextId"`


    


    


    


    


    


    


    


    

}

// Surveyquestion
type Surveyquestion struct { 
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


    // VisibilityCondition
    VisibilityCondition Visibilitycondition `json:"visibilityCondition"`


    // AnswerOptions - Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
    AnswerOptions []Answeroption `json:"answerOptions"`


    // MaxResponseCharacters - How many characters are allowed in the text response to this question. Used by NPS and Free Text question types.
    MaxResponseCharacters int `json:"maxResponseCharacters"`


    // ExplanationPrompt - Prompt for details explaining the chosen NPS score. Used by NPS questions.
    ExplanationPrompt string `json:"explanationPrompt"`

}

// String returns a JSON representation of the model
func (o *Surveyquestion) String() string {
    
    
    
    
    
    
     o.AnswerOptions = []Answeroption{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyquestion) MarshalJSON() ([]byte, error) {
    type Alias Surveyquestion

    if SurveyquestionMarshalled {
        return []byte("{}"), nil
    }
    SurveyquestionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        HelpText string `json:"helpText"`
        
        VarType string `json:"type"`
        
        NaEnabled bool `json:"naEnabled"`
        
        VisibilityCondition Visibilitycondition `json:"visibilityCondition"`
        
        AnswerOptions []Answeroption `json:"answerOptions"`
        
        MaxResponseCharacters int `json:"maxResponseCharacters"`
        
        ExplanationPrompt string `json:"explanationPrompt"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        AnswerOptions: []Answeroption{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

