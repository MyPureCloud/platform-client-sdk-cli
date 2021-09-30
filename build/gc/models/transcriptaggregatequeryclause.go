package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscriptaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscriptaggregatequeryclauseDud struct { 
    


    

}

// Transcriptaggregatequeryclause
type Transcriptaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Transcriptaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Transcriptaggregatequerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcriptaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Transcriptaggregatequeryclause

    if TranscriptaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    TranscriptaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Transcriptaggregatequerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Transcriptaggregatequerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

