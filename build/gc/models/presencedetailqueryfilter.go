package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresencedetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresencedetailqueryfilterDud struct { 
    


    


    

}

// Presencedetailqueryfilter
type Presencedetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Presencedetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Presencedetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Presencedetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Presencedetailqueryclause{{}} 
    
    
    
     o.Predicates = []Presencedetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presencedetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Presencedetailqueryfilter

    if PresencedetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    PresencedetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Presencedetailqueryclause `json:"clauses"`
        
        Predicates []Presencedetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Presencedetailqueryclause{{}},
        

        

        
        Predicates: []Presencedetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

