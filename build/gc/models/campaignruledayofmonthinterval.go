package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofmonthintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofmonthintervalDud struct { }

// Campaignruledayofmonthinterval
type Campaignruledayofmonthinterval struct { }

// String returns a JSON representation of the model
func (o *Campaignruledayofmonthinterval) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofmonthinterval) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofmonthinterval

    if CampaignruledayofmonthintervalMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofmonthintervalMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

