package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaqueryclauseDud struct { 
    


    

}

// Socialmediaqueryclause
type Socialmediaqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Socialmediaquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Socialmediaqueryclause) String() string {
    
     o.Predicates = []Socialmediaquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaqueryclause

    if SocialmediaqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Socialmediaquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Socialmediaquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

