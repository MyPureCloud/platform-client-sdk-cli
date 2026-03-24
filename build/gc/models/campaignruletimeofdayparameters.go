package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruletimeofdayparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruletimeofdayparametersDud struct { 
    


    

}

// Campaignruletimeofdayparameters
type Campaignruletimeofdayparameters struct { 
    // Interval - The operand for the \"between\" operator
    Interval Campaignruletimeofdayinterval `json:"interval"`


    // ThresholdValue - Time is represented as an ISO-8601 string without a timezone. For example: HH:mm:ss.SSS
    ThresholdValue string `json:"thresholdValue"`

}

// String returns a JSON representation of the model
func (o *Campaignruletimeofdayparameters) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruletimeofdayparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruletimeofdayparameters

    if CampaignruletimeofdayparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruletimeofdayparametersMarshalled = true

    return json.Marshal(&struct {
        
        Interval Campaignruletimeofdayinterval `json:"interval"`
        
        ThresholdValue string `json:"thresholdValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

