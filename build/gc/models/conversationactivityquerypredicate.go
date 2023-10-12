package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityquerypredicateDud struct { 
    


    


    


    

}

// Conversationactivityquerypredicate
type Conversationactivityquerypredicate struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // Operator - Optional operator, default is matches
    Operator string `json:"operator"`


    // Value - Right hand side for dimension predicates
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityquerypredicate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityquerypredicate

    if ConversationactivityquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

