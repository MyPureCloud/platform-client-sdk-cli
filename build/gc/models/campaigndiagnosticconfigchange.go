package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigndiagnosticconfigchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigndiagnosticconfigchangeDud struct { 
    


    


    


    

}

// Campaigndiagnosticconfigchange
type Campaigndiagnosticconfigchange struct { 
    // Date - Timestamp when the configuration change occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Date time.Time `json:"date"`


    // Field - Field name that was changed
    Field string `json:"field"`


    // Value - New value assigned to the configuration field
    Value string `json:"value"`


    // Action - Type of operation applied
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Campaigndiagnosticconfigchange) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigndiagnosticconfigchange) MarshalJSON() ([]byte, error) {
    type Alias Campaigndiagnosticconfigchange

    if CampaigndiagnosticconfigchangeMarshalled {
        return []byte("{}"), nil
    }
    CampaigndiagnosticconfigchangeMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        Field string `json:"field"`
        
        Value string `json:"value"`
        
        Action string `json:"action"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

