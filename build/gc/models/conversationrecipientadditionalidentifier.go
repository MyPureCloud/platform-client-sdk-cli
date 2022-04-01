package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationrecipientadditionalidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationrecipientadditionalidentifierDud struct { 
    VarType string `json:"type"`


    Value string `json:"value"`

}

// Conversationrecipientadditionalidentifier - Additional identifiers for describing messaging recipient.
type Conversationrecipientadditionalidentifier struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Conversationrecipientadditionalidentifier) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationrecipientadditionalidentifier) MarshalJSON() ([]byte, error) {
    type Alias Conversationrecipientadditionalidentifier

    if ConversationrecipientadditionalidentifierMarshalled {
        return []byte("{}"), nil
    }
    ConversationrecipientadditionalidentifierMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

