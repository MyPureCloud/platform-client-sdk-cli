package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulespecificdateintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulespecificdateintervalDud struct { }

// Campaignrulespecificdateinterval
type Campaignrulespecificdateinterval struct { }

// String returns a JSON representation of the model
func (o *Campaignrulespecificdateinterval) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulespecificdateinterval) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulespecificdateinterval

    if CampaignrulespecificdateintervalMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulespecificdateintervalMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

