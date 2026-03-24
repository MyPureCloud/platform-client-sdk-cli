package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofweekintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofweekintervalDud struct { 
    


    

}

// Campaignruledayofweekinterval
type Campaignruledayofweekinterval struct { 
    // Min - The minimum value of the interval in 1-7 (Monday-Sunday). Required for the \"between\" operator
    Min int `json:"min"`


    // Max - The maximum value of the interval in 1-7 (Monday-Sunday). Required for the \"between\" operator
    Max int `json:"max"`

}

// String returns a JSON representation of the model
func (o *Campaignruledayofweekinterval) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofweekinterval) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofweekinterval

    if CampaignruledayofweekintervalMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofweekintervalMarshalled = true

    return json.Marshal(&struct {
        
        Min int `json:"min"`
        
        Max int `json:"max"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

