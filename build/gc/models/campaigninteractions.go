package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigninteractionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigninteractionsDud struct { 
    


    


    


    


    


    

}

// Campaigninteractions
type Campaigninteractions struct { 
    // Campaign
    Campaign Domainentityref `json:"campaign"`


    // PendingInteractions
    PendingInteractions []Campaigninteraction `json:"pendingInteractions"`


    // ProceedingInteractions
    ProceedingInteractions []Campaigninteraction `json:"proceedingInteractions"`


    // PreviewingInteractions
    PreviewingInteractions []Campaigninteraction `json:"previewingInteractions"`


    // InteractingInteractions
    InteractingInteractions []Campaigninteraction `json:"interactingInteractions"`


    // ScheduledInteractions
    ScheduledInteractions []Campaigninteraction `json:"scheduledInteractions"`

}

// String returns a JSON representation of the model
func (o *Campaigninteractions) String() string {
    
    
    
    
    
    
     o.PendingInteractions = []Campaigninteraction{{}} 
    
    
    
     o.ProceedingInteractions = []Campaigninteraction{{}} 
    
    
    
     o.PreviewingInteractions = []Campaigninteraction{{}} 
    
    
    
     o.InteractingInteractions = []Campaigninteraction{{}} 
    
    
    
     o.ScheduledInteractions = []Campaigninteraction{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigninteractions) MarshalJSON() ([]byte, error) {
    type Alias Campaigninteractions

    if CampaigninteractionsMarshalled {
        return []byte("{}"), nil
    }
    CampaigninteractionsMarshalled = true

    return json.Marshal(&struct { 
        Campaign Domainentityref `json:"campaign"`
        
        PendingInteractions []Campaigninteraction `json:"pendingInteractions"`
        
        ProceedingInteractions []Campaigninteraction `json:"proceedingInteractions"`
        
        PreviewingInteractions []Campaigninteraction `json:"previewingInteractions"`
        
        InteractingInteractions []Campaigninteraction `json:"interactingInteractions"`
        
        ScheduledInteractions []Campaigninteraction `json:"scheduledInteractions"`
        
        *Alias
    }{
        

        

        

        
        PendingInteractions: []Campaigninteraction{{}},
        

        

        
        ProceedingInteractions: []Campaigninteraction{{}},
        

        

        
        PreviewingInteractions: []Campaigninteraction{{}},
        

        

        
        InteractingInteractions: []Campaigninteraction{{}},
        

        

        
        ScheduledInteractions: []Campaigninteraction{{}},
        

        
        Alias: (*Alias)(u),
    })
}

