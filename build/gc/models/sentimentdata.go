package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SentimentdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SentimentdataDud struct { 
    

}

// Sentimentdata
type Sentimentdata struct { 
    // Insights - The sentiment insights extracted from the conversation
    Insights Sentimentinsights `json:"insights"`

}

// String returns a JSON representation of the model
func (o *Sentimentdata) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sentimentdata) MarshalJSON() ([]byte, error) {
    type Alias Sentimentdata

    if SentimentdataMarshalled {
        return []byte("{}"), nil
    }
    SentimentdataMarshalled = true

    return json.Marshal(&struct {
        
        Insights Sentimentinsights `json:"insights"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

