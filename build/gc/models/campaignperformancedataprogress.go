package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignperformancedataprogressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignperformancedataprogressDud struct { 
    

}

// Campaignperformancedataprogress
type Campaignperformancedataprogress struct { 
    // Percentage - Percentage of contacts processed during the campaign
    Percentage float32 `json:"percentage"`

}

// String returns a JSON representation of the model
func (o *Campaignperformancedataprogress) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignperformancedataprogress) MarshalJSON() ([]byte, error) {
    type Alias Campaignperformancedataprogress

    if CampaignperformancedataprogressMarshalled {
        return []byte("{}"), nil
    }
    CampaignperformancedataprogressMarshalled = true

    return json.Marshal(&struct {
        
        Percentage float32 `json:"percentage"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

