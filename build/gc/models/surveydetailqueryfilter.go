package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveydetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveydetailqueryfilterDud struct { 
    


    


    

}

// Surveydetailqueryfilter
type Surveydetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Surveydetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Surveydetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Surveydetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Surveydetailqueryclause{{}} 
    
    
    
     o.Predicates = []Surveydetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveydetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Surveydetailqueryfilter

    if SurveydetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    SurveydetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Surveydetailqueryclause `json:"clauses"`
        
        Predicates []Surveydetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Surveydetailqueryclause{{}},
        

        

        
        Predicates: []Surveydetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

