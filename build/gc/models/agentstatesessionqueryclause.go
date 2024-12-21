package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatesessionqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatesessionqueryclauseDud struct { 
    


    

}

// Agentstatesessionqueryclause
type Agentstatesessionqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstatesessionquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstatesessionqueryclause) String() string {
    
     o.Predicates = []Agentstatesessionquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatesessionqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Agentstatesessionqueryclause

    if AgentstatesessionqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgentstatesessionqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agentstatesessionquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agentstatesessionquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

