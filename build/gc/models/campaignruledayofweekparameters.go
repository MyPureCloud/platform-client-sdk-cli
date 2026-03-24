package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofweekparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofweekparametersDud struct { 
    


    

}

// Campaignruledayofweekparameters
type Campaignruledayofweekparameters struct { 
    // InSet - The operand for the \"in\" operator, each value in 1-7 (Monday-Sunday) format
    InSet []int `json:"inSet"`


    // Interval - The operand for the \"between\" operator
    Interval Campaignruledayofweekinterval `json:"interval"`

}

// String returns a JSON representation of the model
func (o *Campaignruledayofweekparameters) String() string {
     o.InSet = []int{0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofweekparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofweekparameters

    if CampaignruledayofweekparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofweekparametersMarshalled = true

    return json.Marshal(&struct {
        
        InSet []int `json:"inSet"`
        
        Interval Campaignruledayofweekinterval `json:"interval"`
        *Alias
    }{

        
        InSet: []int{0},
        


        

        Alias: (*Alias)(u),
    })
}

