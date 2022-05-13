package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationaggregatequeryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationaggregatequeryfilterDud struct { 
    


    


    

}

// Conversationaggregatequeryfilter
type Conversationaggregatequeryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Conversationaggregatequeryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationaggregatequerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryfilter) String() string {
    
     o.Clauses = []Conversationaggregatequeryclause{{}} 
     o.Predicates = []Conversationaggregatequerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationaggregatequeryfilter) MarshalJSON() ([]byte, error) {
    type Alias Conversationaggregatequeryfilter

    if ConversationaggregatequeryfilterMarshalled {
        return []byte("{}"), nil
    }
    ConversationaggregatequeryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Conversationaggregatequeryclause `json:"clauses"`
        
        Predicates []Conversationaggregatequerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Conversationaggregatequeryclause{{}},
        


        
        Predicates: []Conversationaggregatequerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

