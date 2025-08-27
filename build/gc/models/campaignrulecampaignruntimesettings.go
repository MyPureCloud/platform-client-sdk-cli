package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulecampaignruntimesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulecampaignruntimesettingsDud struct { }

// Campaignrulecampaignruntimesettings
type Campaignrulecampaignruntimesettings struct { }

// String returns a JSON representation of the model
func (o *Campaignrulecampaignruntimesettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulecampaignruntimesettings) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulecampaignruntimesettings

    if CampaignrulecampaignruntimesettingsMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulecampaignruntimesettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

