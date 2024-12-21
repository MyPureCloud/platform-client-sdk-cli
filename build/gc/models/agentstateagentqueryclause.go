package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateagentqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateagentqueryclauseDud struct { 
    


    

}

// Agentstateagentqueryclause
type Agentstateagentqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Describes a <dimension> = <value> filter used to perform matching
    Predicates []Agentstateagentquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Agentstateagentqueryclause) String() string {
    
     o.Predicates = []Agentstateagentquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateagentqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Agentstateagentqueryclause

    if AgentstateagentqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    AgentstateagentqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Agentstateagentquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Agentstateagentquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

