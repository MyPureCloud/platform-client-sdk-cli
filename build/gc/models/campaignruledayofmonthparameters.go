package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledayofmonthparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledayofmonthparametersDud struct { 
    


    


    

}

// Campaignruledayofmonthparameters
type Campaignruledayofmonthparameters struct { 
    // ThresholdValue - The operand for the \"before\" and \"after\" operators, can be either exact day (1-31) or \"LAST_DAY\"
    ThresholdValue string `json:"thresholdValue"`


    // InSet - The operand for the \"in\" operator, each element can be either exact day (1,31) or \"LAST_DAY\", \"EVEN_DAY\", \"ODD_DAY\"
    InSet []string `json:"inSet"`


    // Interval - The interval operand for the \"between\" operator
    Interval Campaignruledayofmonthinterval `json:"interval"`

}

// String returns a JSON representation of the model
func (o *Campaignruledayofmonthparameters) String() string {
    
     o.InSet = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledayofmonthparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledayofmonthparameters

    if CampaignruledayofmonthparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledayofmonthparametersMarshalled = true

    return json.Marshal(&struct {
        
        ThresholdValue string `json:"thresholdValue"`
        
        InSet []string `json:"inSet"`
        
        Interval Campaignruledayofmonthinterval `json:"interval"`
        *Alias
    }{

        


        
        InSet: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

