package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulecampaignwaittimesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulecampaignwaittimesettingsDud struct { }

// Campaignrulecampaignwaittimesettings
type Campaignrulecampaignwaittimesettings struct { }

// String returns a JSON representation of the model
func (o *Campaignrulecampaignwaittimesettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulecampaignwaittimesettings) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulecampaignwaittimesettings

    if CampaignrulecampaignwaittimesettingsMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulecampaignwaittimesettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

