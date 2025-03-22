package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaqueryfilterDud struct { 
    


    


    

}

// Socialmediaqueryfilter
type Socialmediaqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Socialmediaqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Socialmediaquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Socialmediaqueryfilter) String() string {
    
     o.Clauses = []Socialmediaqueryclause{{}} 
     o.Predicates = []Socialmediaquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaqueryfilter

    if SocialmediaqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Socialmediaqueryclause `json:"clauses"`
        
        Predicates []Socialmediaquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Socialmediaqueryclause{{}},
        


        
        Predicates: []Socialmediaquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

