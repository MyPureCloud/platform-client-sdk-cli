package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresencedetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresencedetailqueryclauseDud struct { 
    


    

}

// Presencedetailqueryclause
type Presencedetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Presencedetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Presencedetailqueryclause) String() string {
    
     o.Predicates = []Presencedetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presencedetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Presencedetailqueryclause

    if PresencedetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    PresencedetailqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Presencedetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Presencedetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

