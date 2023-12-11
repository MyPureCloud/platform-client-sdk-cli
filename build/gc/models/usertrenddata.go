package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsertrenddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsertrenddataDud struct { 
    


    


    


    


    


    

}

// Usertrenddata
type Usertrenddata struct { 
    // DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    // DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEndWorkday time.Time `json:"dateEndWorkday"`


    // PercentOfGoal - Percent of goal
    PercentOfGoal float64 `json:"percentOfGoal"`


    // AverageValue - Average metric value
    AverageValue float64 `json:"averageValue"`


    // RankTotalPoints - Rank, ordered by total points
    RankTotalPoints int `json:"rankTotalPoints"`


    // RankPercentagePoints - Rank, ordered by percentage of points
    RankPercentagePoints int `json:"rankPercentagePoints"`

}

// String returns a JSON representation of the model
func (o *Usertrenddata) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usertrenddata) MarshalJSON() ([]byte, error) {
    type Alias Usertrenddata

    if UsertrenddataMarshalled {
        return []byte("{}"), nil
    }
    UsertrenddataMarshalled = true

    return json.Marshal(&struct {
        
        DateStartWorkday time.Time `json:"dateStartWorkday"`
        
        DateEndWorkday time.Time `json:"dateEndWorkday"`
        
        PercentOfGoal float64 `json:"percentOfGoal"`
        
        AverageValue float64 `json:"averageValue"`
        
        RankTotalPoints int `json:"rankTotalPoints"`
        
        RankPercentagePoints int `json:"rankPercentagePoints"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

