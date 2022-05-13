package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveydetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveydetailqueryclauseDud struct { 
    


    

}

// Surveydetailqueryclause
type Surveydetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Surveydetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Surveydetailqueryclause) String() string {
    
     o.Predicates = []Surveydetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveydetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Surveydetailqueryclause

    if SurveydetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    SurveydetailqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Surveydetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Surveydetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

