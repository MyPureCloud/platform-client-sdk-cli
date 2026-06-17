package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateschedulebidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateschedulebidDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Updateschedulebid
type Updateschedulebid struct { 
    // Name - The name of the schedule bid
    Name string `json:"name"`


    // Forecast - The selected forecast used for schedule set generation for this bid
    Forecast Bushorttermforecastweekreference `json:"forecast"`


    // BidWindowStartDate - The bid start date where agents start participating in schedule bidding relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The bid end date relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when schedule sets would be effective for schedule generation relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // WeeksToSchedule - The number of weeks to generate schedule set through this bid
    WeeksToSchedule int `json:"weeksToSchedule"`


    // EndOverridesAndRotations - If true, all existing overrides, work plan rotations will be ended one day before effective date of this bid
    EndOverridesAndRotations bool `json:"endOverridesAndRotations"`


    // AgentRankingType - The type of agent ranking selected for this bid
    AgentRankingType string `json:"agentRankingType"`


    // RankingTiebreakerType - Ranking tiebreaker to be used
    RankingTiebreakerType string `json:"rankingTiebreakerType"`


    // Status - The state of the bid
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updateschedulebid) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateschedulebid) MarshalJSON() ([]byte, error) {
    type Alias Updateschedulebid

    if UpdateschedulebidMarshalled {
        return []byte("{}"), nil
    }
    UpdateschedulebidMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Forecast Bushorttermforecastweekreference `json:"forecast"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        WeeksToSchedule int `json:"weeksToSchedule"`
        
        EndOverridesAndRotations bool `json:"endOverridesAndRotations"`
        
        AgentRankingType string `json:"agentRankingType"`
        
        RankingTiebreakerType string `json:"rankingTiebreakerType"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

