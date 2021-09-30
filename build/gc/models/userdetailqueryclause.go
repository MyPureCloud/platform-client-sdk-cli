package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserdetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserdetailqueryclauseDud struct { 
    


    

}

// Userdetailqueryclause
type Userdetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Userdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Userdetailqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Userdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userdetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Userdetailqueryclause

    if UserdetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    UserdetailqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Userdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Userdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

