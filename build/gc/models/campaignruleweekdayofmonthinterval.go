package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleweekdayofmonthintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleweekdayofmonthintervalDud struct { 
    


    

}

// Campaignruleweekdayofmonthinterval
type Campaignruleweekdayofmonthinterval struct { 
    // Min - The minimum value of the interval. Required for the \"between\" operator
    Min Campaignruleweekdayofmonth `json:"min"`


    // Max - The maximum value of the interval. Required for the \"between\" operator
    Max Campaignruleweekdayofmonth `json:"max"`

}

// String returns a JSON representation of the model
func (o *Campaignruleweekdayofmonthinterval) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleweekdayofmonthinterval) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleweekdayofmonthinterval

    if CampaignruleweekdayofmonthintervalMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleweekdayofmonthintervalMarshalled = true

    return json.Marshal(&struct {
        
        Min Campaignruleweekdayofmonth `json:"min"`
        
        Max Campaignruleweekdayofmonth `json:"max"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

