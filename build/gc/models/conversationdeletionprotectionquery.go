package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdeletionprotectionqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdeletionprotectionqueryDud struct { 
    

}

// Conversationdeletionprotectionquery
type Conversationdeletionprotectionquery struct { 
    // ConversationIds - This is a list of ConversationIds. The list cannot exceed 100 conversationids.
    ConversationIds []string `json:"conversationIds"`

}

// String returns a JSON representation of the model
func (o *Conversationdeletionprotectionquery) String() string {
    
    
     o.ConversationIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdeletionprotectionquery) MarshalJSON() ([]byte, error) {
    type Alias Conversationdeletionprotectionquery

    if ConversationdeletionprotectionqueryMarshalled {
        return []byte("{}"), nil
    }
    ConversationdeletionprotectionqueryMarshalled = true

    return json.Marshal(&struct { 
        ConversationIds []string `json:"conversationIds"`
        
        *Alias
    }{
        

        
        ConversationIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

