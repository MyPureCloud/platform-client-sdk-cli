package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityqueryclauseDud struct { 
    


    

}

// Conversationactivityqueryclause
type Conversationactivityqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityqueryclause) String() string {
    
     o.Predicates = []Conversationactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityqueryclause

    if ConversationactivityqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Conversationactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Conversationactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

