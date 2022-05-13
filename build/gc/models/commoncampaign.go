package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommoncampaignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommoncampaignDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Commoncampaign
type Commoncampaign struct { 
    


    // Name - The name of the Campaign.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // MediaType - The media type used for this campaign.
    MediaType string `json:"mediaType"`


    

}

// String returns a JSON representation of the model
func (o *Commoncampaign) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commoncampaign) MarshalJSON() ([]byte, error) {
    type Alias Commoncampaign

    if CommoncampaignMarshalled {
        return []byte("{}"), nil
    }
    CommoncampaignMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

