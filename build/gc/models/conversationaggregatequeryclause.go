package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationaggregatequeryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationaggregatequeryclauseDud struct { 
    


    

}

// Conversationaggregatequeryclause
type Conversationaggregatequeryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryclause) String() string {
    
     o.Predicates = []Conversationaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationaggregatequeryclause) MarshalJSON() ([]byte, error) {
    type Alias Conversationaggregatequeryclause

    if ConversationaggregatequeryclauseMarshalled {
        return []byte("{}"), nil
    }
    ConversationaggregatequeryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Conversationaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Conversationaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

