package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationenrichmentlanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationenrichmentlanguageDud struct { 
    

}

// Conversationenrichmentlanguage - Detected language of this message.
type Conversationenrichmentlanguage struct { 
    // Language - The IETF detected language code of this message.
    Language string `json:"language"`

}

// String returns a JSON representation of the model
func (o *Conversationenrichmentlanguage) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationenrichmentlanguage) MarshalJSON() ([]byte, error) {
    type Alias Conversationenrichmentlanguage

    if ConversationenrichmentlanguageMarshalled {
        return []byte("{}"), nil
    }
    ConversationenrichmentlanguageMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

