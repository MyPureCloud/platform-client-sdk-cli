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
type CampaignrulespecificdateparametersDud struct { 
    


    


    

}

// Campaignrulespecificdateparameters
type Campaignrulespecificdateparameters struct { 
    // IncludeYear - If true, includes year in date comparison for specificDate type. When false, only month and day are compared. Default is true. 
    IncludeYear bool `json:"includeYear"`


    // ThresholdValue - The operand for the \"equals\", \"after\" and \"before\" operators in yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false) format.
    ThresholdValue string `json:"thresholdValue"`


    // Interval - The operand for the \"between\" operator
    Interval Campaignrulespecificdateinterval `json:"interval"`

}

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
        
        IncludeYear bool `json:"includeYear"`
        
        ThresholdValue string `json:"thresholdValue"`
        
        Interval Campaignrulespecificdateinterval `json:"interval"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

