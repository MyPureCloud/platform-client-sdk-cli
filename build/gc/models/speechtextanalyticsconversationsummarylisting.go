package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SpeechtextanalyticsconversationsummarylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SpeechtextanalyticsconversationsummarylistingDud struct { 
    

}

// Speechtextanalyticsconversationsummarylisting
type Speechtextanalyticsconversationsummarylisting struct { 
    // Entities
    Entities []Speechtextanalyticsconversationsummary `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Speechtextanalyticsconversationsummarylisting) String() string {
     o.Entities = []Speechtextanalyticsconversationsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Speechtextanalyticsconversationsummarylisting) MarshalJSON() ([]byte, error) {
    type Alias Speechtextanalyticsconversationsummarylisting

    if SpeechtextanalyticsconversationsummarylistingMarshalled {
        return []byte("{}"), nil
    }
    SpeechtextanalyticsconversationsummarylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Speechtextanalyticsconversationsummary `json:"entities"`
        *Alias
    }{

        
        Entities: []Speechtextanalyticsconversationsummary{{}},
        

        Alias: (*Alias)(u),
    })
}

