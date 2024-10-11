package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruleexecutionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruleexecutionsettingsDud struct { 
    


    

}

// Campaignruleexecutionsettings
type Campaignruleexecutionsettings struct { 
    // Frequency - Execution control frequency
    Frequency string `json:"frequency"`


    // TimeZoneId - The time zone for the execution control frequency=\"oncePerDay\"; for example, Africa/Abidjan. This property is ignored when frequency is not \"oncePerDay\"
    TimeZoneId string `json:"timeZoneId"`

}

// String returns a JSON representation of the model
func (o *Campaignruleexecutionsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruleexecutionsettings) MarshalJSON() ([]byte, error) {
    type Alias Campaignruleexecutionsettings

    if CampaignruleexecutionsettingsMarshalled {
        return []byte("{}"), nil
    }
    CampaignruleexecutionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Frequency string `json:"frequency"`
        
        TimeZoneId string `json:"timeZoneId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

