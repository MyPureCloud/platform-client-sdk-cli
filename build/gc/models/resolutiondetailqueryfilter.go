package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutiondetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutiondetailqueryfilterDud struct { 
    


    


    

}

// Resolutiondetailqueryfilter
type Resolutiondetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Resolutiondetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Resolutiondetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Resolutiondetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Resolutiondetailqueryclause{{}} 
    
    
    
     o.Predicates = []Resolutiondetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutiondetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Resolutiondetailqueryfilter

    if ResolutiondetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    ResolutiondetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Resolutiondetailqueryclause `json:"clauses"`
        
        Predicates []Resolutiondetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Resolutiondetailqueryclause{{}},
        

        

        
        Predicates: []Resolutiondetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

