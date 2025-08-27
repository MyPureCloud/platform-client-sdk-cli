package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleweekdayofmonthparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleweekdayofmonthparametersDud struct { }

// Campaignruleweekdayofmonthparameters
type Campaignruleweekdayofmonthparameters struct { }

// String returns a JSON representation of the model
func (o *Campaignruleweekdayofmonthparameters) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleweekdayofmonthparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleweekdayofmonthparameters

    if CampaignruleweekdayofmonthparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleweekdayofmonthparametersMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

