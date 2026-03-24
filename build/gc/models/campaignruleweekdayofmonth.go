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
type CampaignruleweekdayofmonthDud struct { 
    


    


    

}

// Campaignruleweekdayofmonth
type Campaignruleweekdayofmonth struct { 
    // DayOfWeek - Day of week (1-7)
    DayOfWeek int `json:"dayOfWeek"`


    // Month - Month (1-12)
    Month int `json:"month"`


    // Occurrence - Occurrence 1-4, -1 (last)
    Occurrence int `json:"occurrence"`

}

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
        
        DayOfWeek int `json:"dayOfWeek"`
        
        Month int `json:"month"`
        
        Occurrence int `json:"occurrence"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

