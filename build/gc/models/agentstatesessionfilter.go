package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatesessionfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatesessionfilterDud struct { 
    


    


    

}

// Agentstatesessionfilter
type Agentstatesessionfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Agentstatesessionqueryclause `json:"clauses"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstatesessionquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstatesessionfilter) String() string {
    
     o.Clauses = []Agentstatesessionqueryclause{{}} 
     o.Predicates = []Agentstatesessionquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatesessionfilter) MarshalJSON() ([]byte, error) {
    type Alias Agentstatesessionfilter

    if AgentstatesessionfilterMarshalled {
        return []byte("{}"), nil
    }
    AgentstatesessionfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Agentstatesessionqueryclause `json:"clauses"`
        
        Predicates []Agentstatesessionquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Agentstatesessionqueryclause{{}},
        


        
        Predicates: []Agentstatesessionquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

