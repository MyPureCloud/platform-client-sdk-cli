package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulespecificdateparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulespecificdateparametersDud struct { }

// Campaignrulespecificdateparameters
type Campaignrulespecificdateparameters struct { }

// String returns a JSON representation of the model
func (o *Campaignrulespecificdateparameters) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulespecificdateparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulespecificdateparameters

    if CampaignrulespecificdateparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulespecificdateparametersMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

