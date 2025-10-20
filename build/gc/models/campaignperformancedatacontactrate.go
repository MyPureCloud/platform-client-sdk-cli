package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignperformancedatacontactrateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignperformancedatacontactrateDud struct { 
    

}

// Campaignperformancedatacontactrate
type Campaignperformancedatacontactrate struct { 
    // ConnectRatio - Ratio of connects to attempts
    ConnectRatio float32 `json:"connectRatio"`

}

// String returns a JSON representation of the model
func (o *Campaignperformancedatacontactrate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignperformancedatacontactrate) MarshalJSON() ([]byte, error) {
    type Alias Campaignperformancedatacontactrate

    if CampaignperformancedatacontactrateMarshalled {
        return []byte("{}"), nil
    }
    CampaignperformancedatacontactrateMarshalled = true

    return json.Marshal(&struct {
        
        ConnectRatio float32 `json:"connectRatio"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

