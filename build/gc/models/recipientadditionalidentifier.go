package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecipientadditionalidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecipientadditionalidentifierDud struct { 
    VarType string `json:"type"`


    Value string `json:"value"`

}

// Recipientadditionalidentifier - Additional identifiers for describing messaging recipient.
type Recipientadditionalidentifier struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Recipientadditionalidentifier) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recipientadditionalidentifier) MarshalJSON() ([]byte, error) {
    type Alias Recipientadditionalidentifier

    if RecipientadditionalidentifierMarshalled {
        return []byte("{}"), nil
    }
    RecipientadditionalidentifierMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

