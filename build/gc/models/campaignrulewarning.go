package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignrulewarningMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignrulewarningDud struct { 
    


    


    

}

// Campaignrulewarning
type Campaignrulewarning struct { 
    // Code - Warning code for this warning.
    Code string `json:"code"`


    // Message - Warning message for this warning.
    Message string `json:"message"`


    // Params - Additional warning information
    Params Campaignrulewarningparameters `json:"params"`

}

// String returns a JSON representation of the model
func (o *Campaignrulewarning) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignrulewarning) MarshalJSON() ([]byte, error) {
    type Alias Campaignrulewarning

    if CampaignrulewarningMarshalled {
        return []byte("{}"), nil
    }
    CampaignrulewarningMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Params Campaignrulewarningparameters `json:"params"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

