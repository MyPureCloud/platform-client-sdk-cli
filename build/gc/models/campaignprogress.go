package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignprogressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignprogressDud struct { 
    


    


    NumberOfContactsCalled int `json:"numberOfContactsCalled"`


    NumberOfContactsMessaged int `json:"numberOfContactsMessaged"`


    TotalNumberOfContacts int `json:"totalNumberOfContacts"`


    Percentage int `json:"percentage"`

}

// Campaignprogress
type Campaignprogress struct { 
    // Campaign - Identifier of the campaign
    Campaign Domainentityref `json:"campaign"`


    // ContactList - Identifier of the contact list
    ContactList Domainentityref `json:"contactList"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignprogress) MarshalJSON() ([]byte, error) {
    type Alias Campaignprogress

    if CampaignprogressMarshalled {
        return []byte("{}"), nil
    }
    CampaignprogressMarshalled = true

    return json.Marshal(&struct { 
        Campaign Domainentityref `json:"campaign"`
        
        ContactList Domainentityref `json:"contactList"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

