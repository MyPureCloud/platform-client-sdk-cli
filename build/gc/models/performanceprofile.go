package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PerformanceprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PerformanceprofileDud struct { 
    Id string `json:"id"`


    


    Division Division `json:"division"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    


    Active bool `json:"active"`


    MemberCount int `json:"memberCount"`


    


    SelfUri string `json:"selfUri"`

}

// Performanceprofile
type Performanceprofile struct { 
    


    // Name - A name for this performance profile
    Name string `json:"name"`


    


    // Description - A description about this performance profile
    Description string `json:"description"`


    // MetricOrders - Order of the associated metrics. The list should contain valid ids for metrics
    MetricOrders []string `json:"metricOrders"`


    


    // ReportingIntervals - The reporting interval periods for this performance profile
    ReportingIntervals []Reportinginterval `json:"reportingIntervals"`


    


    


    // MaxLeaderboardRankSize - The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
    MaxLeaderboardRankSize int `json:"maxLeaderboardRankSize"`


    

}

// String returns a JSON representation of the model
func (o *Performanceprofile) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MetricOrders = []string{""} 
    
    
    
    
    
     o.ReportingIntervals = []Reportinginterval{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Performanceprofile) MarshalJSON() ([]byte, error) {
    type Alias Performanceprofile

    if PerformanceprofileMarshalled {
        return []byte("{}"), nil
    }
    PerformanceprofileMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        Description string `json:"description"`
        
        MetricOrders []string `json:"metricOrders"`
        
        
        
        ReportingIntervals []Reportinginterval `json:"reportingIntervals"`
        
        
        
        
        
        MaxLeaderboardRankSize int `json:"maxLeaderboardRankSize"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        MetricOrders: []string{""},
        

        

        

        

        
        ReportingIntervals: []Reportinginterval{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

