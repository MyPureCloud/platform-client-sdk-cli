package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResolutiondetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResolutiondetailqueryclauseDud struct { 
    


    

}

// Resolutiondetailqueryclause
type Resolutiondetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Resolutiondetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Resolutiondetailqueryclause) String() string {
    
     o.Predicates = []Resolutiondetailquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resolutiondetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Resolutiondetailqueryclause

    if ResolutiondetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    ResolutiondetailqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Resolutiondetailquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Resolutiondetailquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

