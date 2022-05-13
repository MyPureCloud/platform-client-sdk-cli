package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentformDud struct { 
    Id string `json:"id"`


    DateModified time.Time `json:"dateModified"`


    ContextId string `json:"contextId"`


    SelfUri string `json:"selfUri"`


    Published bool `json:"published"`


    


    

}

// Assessmentform
type Assessmentform struct { 
    


    


    


    


    


    // PassPercent - The pass percent for the assessment form
    PassPercent int `json:"passPercent"`


    // QuestionGroups - A list of question groups
    QuestionGroups []Assessmentformquestiongroup `json:"questionGroups"`

}

// String returns a JSON representation of the model
func (o *Assessmentform) String() string {
    
     o.QuestionGroups = []Assessmentformquestiongroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentform) MarshalJSON() ([]byte, error) {
    type Alias Assessmentform

    if AssessmentformMarshalled {
        return []byte("{}"), nil
    }
    AssessmentformMarshalled = true

    return json.Marshal(&struct {
        
        PassPercent int `json:"passPercent"`
        
        QuestionGroups []Assessmentformquestiongroup `json:"questionGroups"`
        *Alias
    }{

        


        


        


        


        


        


        
        QuestionGroups: []Assessmentformquestiongroup{{}},
        

        Alias: (*Alias)(u),
    })
}

