package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticcampaignactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticcampaignactionDud struct { }

// Campaigndiagnosticcampaignaction
type Campaigndiagnosticcampaignaction struct { }

// String returns a JSON representation of the model
func (o *Campaigndiagnosticcampaignaction) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnosticcampaignaction) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnosticcampaignaction

    if CampaigndiagnosticcampaignactionMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticcampaignactionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

