package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationinsightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationinsightDud struct { 
    


    


    

}

// Conversationinsight
type Conversationinsight struct { 
    // VarType - The type of insight
    VarType string `json:"type"`


    // Title - The reason for contact, resolution for the interaction, or follow-up action item
    Title string `json:"title"`


    // Description - Reasoning for the given insight
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Conversationinsight) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationinsight) MarshalJSON() ([]byte, error) {
    type Alias Conversationinsight

    if ConversationinsightMarshalled {
        return []byte("{}"), nil
    }
    ConversationinsightMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

