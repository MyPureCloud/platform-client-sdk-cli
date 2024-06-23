package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignbusinesscategorymetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignbusinesscategorymetricsDud struct { 
    


    


    

}

// Campaignbusinesscategorymetrics
type Campaignbusinesscategorymetrics struct { 
    // SuccessCount - Number of calls categorized as business success
    SuccessCount int `json:"successCount"`


    // NeutralCount - Number of calls categorized as business neutral
    NeutralCount int `json:"neutralCount"`


    // FailureCount - Number of calls categorized as business failure
    FailureCount int `json:"failureCount"`

}

// String returns a JSON representation of the model
func (o *Campaignbusinesscategorymetrics) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignbusinesscategorymetrics) MarshalJSON() ([]byte, error) {
    type Alias Campaignbusinesscategorymetrics

    if CampaignbusinesscategorymetricsMarshalled {
        return []byte("{}"), nil
    }
    CampaignbusinesscategorymetricsMarshalled = true

    return json.Marshal(&struct {
        
        SuccessCount int `json:"successCount"`
        
        NeutralCount int `json:"neutralCount"`
        
        FailureCount int `json:"failureCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

