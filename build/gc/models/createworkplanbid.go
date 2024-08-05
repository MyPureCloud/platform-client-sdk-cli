package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateworkplanbidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateworkplanbidDud struct { 
    


    


    


    


    


    


    


    

}

// Createworkplanbid - Create work plan bid model
type Createworkplanbid struct { 
    // Name - The name of the work plan bid
    Name string `json:"name"`


    // Forecast - The selected forecast in this work plan bid
    Forecast Bushorttermforecastweekreference `json:"forecast"`


    // BidWindowStartDate - The bid start date where agents start participate in work plan bidding in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The bid end date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when agents will be assigned to the new work plan in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // AgentRankingType - The type of agent ranking selected for this bid
    AgentRankingType string `json:"agentRankingType"`


    // RankingTiebreakerType - Ranking tiebreaker to be used
    RankingTiebreakerType string `json:"rankingTiebreakerType"`


    // WorkPlanFieldsVisibleToAgents - The work plan fields visible to agents whenever work plan preferences are made
    WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`

}

// String returns a JSON representation of the model
func (o *Createworkplanbid) String() string {
    
    
    
    
    
    
    
     o.WorkPlanFieldsVisibleToAgents = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createworkplanbid) MarshalJSON() ([]byte, error) {
    type Alias Createworkplanbid

    if CreateworkplanbidMarshalled {
        return []byte("{}"), nil
    }
    CreateworkplanbidMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Forecast Bushorttermforecastweekreference `json:"forecast"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        AgentRankingType string `json:"agentRankingType"`
        
        RankingTiebreakerType string `json:"rankingTiebreakerType"`
        
        WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        WorkPlanFieldsVisibleToAgents: []string{""},
        

        Alias: (*Alias)(u),
    })
}

