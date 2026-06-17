package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminagentschedulesetpreferencesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminagentschedulesetpreferencesDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adminagentschedulesetpreferences
type Adminagentschedulesetpreferences struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Bid - The schedule bid
    Bid Schedulebidreference `json:"bid"`


    // BidGroup - The schedule bid group
    BidGroup Schedulebidgroupreference `json:"bidGroup"`


    // AgentsScheduleBidPreferences - The agents' schedule bidding preferences
    AgentsScheduleBidPreferences []Adminagentschedulebidbiddingpreference `json:"agentsScheduleBidPreferences"`


    

}

// String returns a JSON representation of the model
func (o *Adminagentschedulesetpreferences) String() string {
    
    
    
     o.AgentsScheduleBidPreferences = []Adminagentschedulebidbiddingpreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminagentschedulesetpreferences) MarshalJSON() ([]byte, error) {
    type Alias Adminagentschedulesetpreferences

    if AdminagentschedulesetpreferencesMarshalled {
        return []byte("{}"), nil
    }
    AdminagentschedulesetpreferencesMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Bid Schedulebidreference `json:"bid"`
        
        BidGroup Schedulebidgroupreference `json:"bidGroup"`
        
        AgentsScheduleBidPreferences []Adminagentschedulebidbiddingpreference `json:"agentsScheduleBidPreferences"`
        *Alias
    }{

        


        


        


        
        AgentsScheduleBidPreferences: []Adminagentschedulebidbiddingpreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

