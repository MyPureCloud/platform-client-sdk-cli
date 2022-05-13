package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignsequenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignsequenceDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    CurrentCampaign int `json:"currentCampaign"`


    


    StopMessage string `json:"stopMessage"`


    


    SelfUri string `json:"selfUri"`

}

// Campaignsequence
type Campaignsequence struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Campaigns - The ordered list of Campaigns that this CampaignSequence will run.
    Campaigns []Domainentityref `json:"campaigns"`


    


    // Status - The current status of the CampaignSequence. A CampaignSequence can be turned 'on' or 'off'.
    Status string `json:"status"`


    


    // Repeat - Indicates if a sequence should repeat from the beginning after the last campaign completes. Default is false.
    Repeat bool `json:"repeat"`


    

}

// String returns a JSON representation of the model
func (o *Campaignsequence) String() string {
    
    
     o.Campaigns = []Domainentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignsequence) MarshalJSON() ([]byte, error) {
    type Alias Campaignsequence

    if CampaignsequenceMarshalled {
        return []byte("{}"), nil
    }
    CampaignsequenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Campaigns []Domainentityref `json:"campaigns"`
        
        Status string `json:"status"`
        
        Repeat bool `json:"repeat"`
        *Alias
    }{

        


        


        


        


        


        
        Campaigns: []Domainentityref{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

