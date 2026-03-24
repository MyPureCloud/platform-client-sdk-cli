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
type CampaignruledayofmonthintervalDud struct { 
    


    

}

// Campaignruledayofmonthinterval
type Campaignruledayofmonthinterval struct { 
    // Min - The minimal day of month (exact day: 1-31) for the \"between\" operator
    Min string `json:"min"`


    // Max - The maximum value of month (exact day: 1-31 or \"LAST_DAY\") for the \"between\" operator
    Max string `json:"max"`

}

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
        
        Min string `json:"min"`
        
        Max string `json:"max"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

