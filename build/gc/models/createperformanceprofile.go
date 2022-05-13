package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateperformanceprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateperformanceprofileDud struct { 
    Id string `json:"id"`


    


    


    


    MetricOrders []string `json:"metricOrders"`


    DateCreated time.Time `json:"dateCreated"`


    


    


    MemberCount int `json:"memberCount"`


    


    SelfUri string `json:"selfUri"`

}

// Createperformanceprofile
type Createperformanceprofile struct { 
    


    // Name - A name for this performance profile
    Name string `json:"name"`


    // Division - The associated division for this Performance Profile
    Division Writabledivision `json:"division"`


    // Description - A description about this performance profile
    Description string `json:"description"`


    


    


    // ReportingIntervals - The reporting interval periods for this performance profile
    ReportingIntervals []Reportinginterval `json:"reportingIntervals"`


    // Active - The flag for active profiles
    Active bool `json:"active"`


    


    // MaxLeaderboardRankSize - The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
    MaxLeaderboardRankSize int `json:"maxLeaderboardRankSize"`


    

}

// String returns a JSON representation of the model
func (o *Createperformanceprofile) String() string {
    
    
    
     o.ReportingIntervals = []Reportinginterval{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createperformanceprofile) MarshalJSON() ([]byte, error) {
    type Alias Createperformanceprofile

    if CreateperformanceprofileMarshalled {
        return []byte("{}"), nil
    }
    CreateperformanceprofileMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        ReportingIntervals []Reportinginterval `json:"reportingIntervals"`
        
        Active bool `json:"active"`
        
        MaxLeaderboardRankSize int `json:"maxLeaderboardRankSize"`
        *Alias
    }{

        


        


        


        


        


        


        
        ReportingIntervals: []Reportinginterval{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

