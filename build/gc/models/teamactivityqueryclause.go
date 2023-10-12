package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivityqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivityqueryclauseDud struct { 
    


    

}

// Teamactivityqueryclause
type Teamactivityqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Teamactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Teamactivityqueryclause) String() string {
    
     o.Predicates = []Teamactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivityqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Teamactivityqueryclause

    if TeamactivityqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    TeamactivityqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Teamactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Teamactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

