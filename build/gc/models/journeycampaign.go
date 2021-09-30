package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneycampaignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneycampaignDud struct { 
    


    


    


    


    


    


    

}

// Journeycampaign
type Journeycampaign struct { 
    // Content - Differentiate ads or links that point to the same URL (e.g. textlink).
    Content string `json:"content"`


    // Medium - Identify a medium such as email or cost-per-click (e.g. CPC).
    Medium string `json:"medium"`


    // Name - Identify a specific product promotion or strategic campaign (e.g. 320banner).
    Name string `json:"name"`


    // Source - Identify a search engine, newsletter name, or other source (e.g. Google).
    Source string `json:"source"`


    // Term - Note the keywords for this ad (e.g. running+shoes).
    Term string `json:"term"`


    // ClickId - The click ID (unique number that is generated when a potential customer clicks on an affiliate link).
    ClickId string `json:"clickId"`


    // Network - The ad network to which the click ID belongs.
    Network string `json:"network"`

}

// String returns a JSON representation of the model
func (o *Journeycampaign) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeycampaign) MarshalJSON() ([]byte, error) {
    type Alias Journeycampaign

    if JourneycampaignMarshalled {
        return []byte("{}"), nil
    }
    JourneycampaignMarshalled = true

    return json.Marshal(&struct { 
        Content string `json:"content"`
        
        Medium string `json:"medium"`
        
        Name string `json:"name"`
        
        Source string `json:"source"`
        
        Term string `json:"term"`
        
        ClickId string `json:"clickId"`
        
        Network string `json:"network"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

