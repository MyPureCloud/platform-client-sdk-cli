package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulebid
type Schedulebid struct { 
    // Id - The ID of the schedule bid
    Id string `json:"id"`


    // Name - The name of the schedule bid
    Name string `json:"name"`


    // BidWindowStartDate - The bid start date when agents can start participating in schedule bidding relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The bid end date relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when schedule sets would be effective for schedule generation relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // Status - The state of the bid
    Status string `json:"status"`


    // BidType - The type of the bid
    BidType string `json:"bidType"`


    // Forecast - The selected forecast used for schedule set generation for this bid
    Forecast Bushorttermforecastweekreference `json:"forecast"`


    // WeeksToSchedule - The number of weeks to generate schedule sets through this bid
    WeeksToSchedule int `json:"weeksToSchedule"`


    // EndOverridesAndRotations - If true, all existing overrides, work plan rotations will be ended one day before effective date of this bid
    EndOverridesAndRotations bool `json:"endOverridesAndRotations"`


    // AgentRankingType - The type of agent ranking selected for this bid
    AgentRankingType string `json:"agentRankingType"`


    // RankingTiebreakerType - Ranking tiebreaker
    RankingTiebreakerType string `json:"rankingTiebreakerType"`


    // PublishedDate - The date the schedule bid is published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PublishedDate time.Time `json:"publishedDate"`


    // EndDate - The end date until which schedule sets can be used for schedule generation. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndDate time.Time `json:"endDate"`


    // Metadata - The metadata of this bid
    Metadata Workplanbidmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Schedulebid) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebid) MarshalJSON() ([]byte, error) {
    type Alias Schedulebid

    if SchedulebidMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        Status string `json:"status"`
        
        BidType string `json:"bidType"`
        
        Forecast Bushorttermforecastweekreference `json:"forecast"`
        
        WeeksToSchedule int `json:"weeksToSchedule"`
        
        EndOverridesAndRotations bool `json:"endOverridesAndRotations"`
        
        AgentRankingType string `json:"agentRankingType"`
        
        RankingTiebreakerType string `json:"rankingTiebreakerType"`
        
        PublishedDate time.Time `json:"publishedDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Metadata Workplanbidmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

