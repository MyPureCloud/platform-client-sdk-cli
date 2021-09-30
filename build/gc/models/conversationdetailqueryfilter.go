package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdetailqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdetailqueryfilterDud struct { 
    


    


    

}

// Conversationdetailqueryfilter
type Conversationdetailqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Conversationdetailqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailqueryfilter) String() string {
    
    
    
    
    
    
     o.Clauses = []Conversationdetailqueryclause{{}} 
    
    
    
     o.Predicates = []Conversationdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdetailqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Conversationdetailqueryfilter

    if ConversationdetailqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    ConversationdetailqueryfilterMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Clauses []Conversationdetailqueryclause `json:"clauses"`
        
        Predicates []Conversationdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Clauses: []Conversationdetailqueryclause{{}},
        

        

        
        Predicates: []Conversationdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

