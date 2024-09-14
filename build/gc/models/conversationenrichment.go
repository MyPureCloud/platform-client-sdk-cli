package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationenrichmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationenrichmentDud struct { 
    


    

}

// Conversationenrichment - Metadata enrichments provided by the platform.
type Conversationenrichment struct { 
    // Language - Detected language of this message.
    Language Conversationenrichmentlanguage `json:"language"`


    // SentimentV2 - Detected sentiment of this message.
    SentimentV2 Conversationenrichmentsentimentv2 `json:"sentimentV2"`

}

// String returns a JSON representation of the model
func (o *Conversationenrichment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationenrichment) MarshalJSON() ([]byte, error) {
    type Alias Conversationenrichment

    if ConversationenrichmentMarshalled {
        return []byte("{}"), nil
    }
    ConversationenrichmentMarshalled = true

    return json.Marshal(&struct {
        
        Language Conversationenrichmentlanguage `json:"language"`
        
        SentimentV2 Conversationenrichmentsentimentv2 `json:"sentimentV2"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

