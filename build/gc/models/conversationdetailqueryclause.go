package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdetailqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdetailqueryclauseDud struct { 
    


    

}

// Conversationdetailqueryclause
type Conversationdetailqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Conversationdetailquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailqueryclause) String() string {
    
    
    
    
    
    
     o.Predicates = []Conversationdetailquerypredicate{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdetailqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Conversationdetailqueryclause

    if ConversationdetailqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    ConversationdetailqueryclauseMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Predicates []Conversationdetailquerypredicate `json:"predicates"`
        
        *Alias
    }{
        

        

        

        
        Predicates: []Conversationdetailquerypredicate{{}},
        

        
        Alias: (*Alias)(u),
    })
}

