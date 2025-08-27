package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofweekparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofweekparametersDud struct { }

// Campaignruledayofweekparameters
type Campaignruledayofweekparameters struct { }

// String returns a JSON representation of the model
func (o *Campaignruledayofweekparameters) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofweekparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofweekparameters

    if CampaignruledayofweekparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofweekparametersMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

