package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Conversationenrichmentsentimentv2Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Conversationenrichmentsentimentv2Dud struct { 
    

}

// Conversationenrichmentsentimentv2 - Sentiment analysis of this message.
type Conversationenrichmentsentimentv2 struct { 
    // Tag - Detected Sentiment tag
    Tag string `json:"tag"`

}

// String returns a JSON representation of the model
func (o *Conversationenrichmentsentimentv2) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationenrichmentsentimentv2) MarshalJSON() ([]byte, error) {
    type Alias Conversationenrichmentsentimentv2

    if Conversationenrichmentsentimentv2Marshalled {
        return []byte("{}"), nil
    }
    Conversationenrichmentsentimentv2Marshalled = true

    return json.Marshal(&struct {
        
        Tag string `json:"tag"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

