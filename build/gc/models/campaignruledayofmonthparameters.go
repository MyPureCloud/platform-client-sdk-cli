package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofmonthparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofmonthparametersDud struct { }

// Campaignruledayofmonthparameters
type Campaignruledayofmonthparameters struct { }

// String returns a JSON representation of the model
func (o *Campaignruledayofmonthparameters) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofmonthparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofmonthparameters

    if CampaignruledayofmonthparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofmonthparametersMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

