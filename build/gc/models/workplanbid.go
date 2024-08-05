package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workplanbid - Work plan bid response
type Workplanbid struct { 
    // Id - The ID of the work plan bid
    Id string `json:"id"`


    // Name - The name of the work plan bid
    Name string `json:"name"`


    // Forecast - The selected forecast in this work plan bid
    Forecast Bushorttermforecastweekreference `json:"forecast"`


    // BidWindowStartDate - The bid start date where agents start participate in work plan bidding. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The bid end date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when agents will be assigned to the new work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // Status - The state of the bid
    Status string `json:"status"`


    // AgentRankingType - The type of agent ranking selected for this bid
    AgentRankingType string `json:"agentRankingType"`


    // RankingTiebreakerType - Ranking tiebreaker
    RankingTiebreakerType string `json:"rankingTiebreakerType"`


    // PublishedDate - The date the work plan bid published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PublishedDate time.Time `json:"publishedDate"`


    // WorkPlanFieldsVisibleToAgents - The work plan fields visible to agents whenever work plan preferences are made
    WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`


    // Metadata - The meta data of this bid
    Metadata Workplanbidmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Workplanbid) String() string {
    
    
    
    
    
    
    
    
    
    
     o.WorkPlanFieldsVisibleToAgents = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbid) MarshalJSON() ([]byte, error) {
    type Alias Workplanbid

    if WorkplanbidMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Forecast Bushorttermforecastweekreference `json:"forecast"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        Status string `json:"status"`
        
        AgentRankingType string `json:"agentRankingType"`
        
        RankingTiebreakerType string `json:"rankingTiebreakerType"`
        
        PublishedDate time.Time `json:"publishedDate"`
        
        WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`
        
        Metadata Workplanbidmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        WorkPlanFieldsVisibleToAgents: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

