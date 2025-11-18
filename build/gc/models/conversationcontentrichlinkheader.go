package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentrichlinkheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentrichlinkheaderDud struct { 
    


    

}

// Conversationcontentrichlinkheader - Header for a Rich Link
type Conversationcontentrichlinkheader struct { 
    // VarType - Describes the Rich Link header type.
    VarType string `json:"type"`


    // Value - Rich Link header value.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentrichlinkheader) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentrichlinkheader) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentrichlinkheader

    if ConversationcontentrichlinkheaderMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentrichlinkheaderMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

