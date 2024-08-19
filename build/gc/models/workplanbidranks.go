package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidranksMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidranksDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Workplanbidranks
type Workplanbidranks struct { 
    // Id - The globally unique identifier for the user.
    Id string `json:"id"`


    // WorkPlanBiddingPerformance - Work plan bidding agent performance ranking. The range of values is between 0 and 9999.
    WorkPlanBiddingPerformance int `json:"workPlanBiddingPerformance"`


    // BiddingTieBreaker - Custom agent ranking metric that some customers can use.
    BiddingTieBreaker string `json:"biddingTieBreaker"`


    

}

// String returns a JSON representation of the model
func (o *Workplanbidranks) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidranks) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidranks

    if WorkplanbidranksMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidranksMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WorkPlanBiddingPerformance int `json:"workPlanBiddingPerformance"`
        
        BiddingTieBreaker string `json:"biddingTieBreaker"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

