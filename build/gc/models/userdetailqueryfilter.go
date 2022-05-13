package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserdetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserdetailqueryfilterDud struct { 
    


    


    

}

// Userdetailqueryfilter
type Userdetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Userdetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Userdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Userdetailqueryfilter) String() string {
    
     o.Clauses = []Userdetailqueryclause{{}} 
     o.Predicates = []Userdetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userdetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Userdetailqueryfilter

    if UserdetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    UserdetailqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Userdetailqueryclause `json:"clauses"`
        
        Predicates []Userdetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Userdetailqueryclause{{}},
        


        
        Predicates: []Userdetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

