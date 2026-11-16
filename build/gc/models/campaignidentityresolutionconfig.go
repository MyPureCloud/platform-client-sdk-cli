package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignidentityresolutionconfigDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Campaignidentityresolutionconfig
type Campaignidentityresolutionconfig struct { 
    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Campaignidentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Campaignidentityresolutionconfig

    if CampaignidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    CampaignidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

