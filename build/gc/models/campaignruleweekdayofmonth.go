package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleweekdayofmonthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleweekdayofmonthDud struct { }

// Campaignruleweekdayofmonth
type Campaignruleweekdayofmonth struct { }

// String returns a JSON representation of the model
func (o *Campaignruleweekdayofmonth) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleweekdayofmonth) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleweekdayofmonth

    if CampaignruleweekdayofmonthMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleweekdayofmonthMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

