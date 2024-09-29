package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleexecutionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleexecutionsettingsDud struct { }

// Campaignruleexecutionsettings
type Campaignruleexecutionsettings struct { }

// String returns a JSON representation of the model
func (o *Campaignruleexecutionsettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleexecutionsettings) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleexecutionsettings

    if CampaignruleexecutionsettingsMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleexecutionsettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

