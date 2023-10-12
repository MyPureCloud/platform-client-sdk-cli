package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityqueryfilterDud struct { 
    


    


    

}

// Conversationactivityqueryfilter
type Conversationactivityqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Conversationactivityqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityqueryfilter) String() string {
    
     o.Clauses = []Conversationactivityqueryclause{{}} 
     o.Predicates = []Conversationactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityqueryfilter

    if ConversationactivityqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Conversationactivityqueryclause `json:"clauses"`
        
        Predicates []Conversationactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Conversationactivityqueryclause{{}},
        


        
        Predicates: []Conversationactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

