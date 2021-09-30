package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyaggregatequeryclauseDud struct { 
    


    

}

// Surveyaggregatequeryclause
type Surveyaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Surveyaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Surveyaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Surveyaggregatequeryclause

    if SurveyaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    SurveyaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Surveyaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Surveyaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

