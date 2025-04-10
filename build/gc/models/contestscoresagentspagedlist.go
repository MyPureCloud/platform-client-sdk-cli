package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestscoresagentspagedlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestscoresagentspagedlistDud struct { 
    


    


    


    


    


    

}

// Contestscoresagentspagedlist
type Contestscoresagentspagedlist struct { 
    // Entities
    Entities []Contestscoresagents `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // DateWorkday - Workday of the contest scores leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateWorkday time.Time `json:"dateWorkday"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Contestscoresagentspagedlist) String() string {
     o.Entities = []Contestscoresagents{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestscoresagentspagedlist) MarshalJSON() ([]byte, error) {
    type Alias Contestscoresagentspagedlist

    if ContestscoresagentspagedlistMarshalled {
        return []byte("{}"), nil
    }
    ContestscoresagentspagedlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contestscoresagents `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        DateWorkday time.Time `json:"dateWorkday"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Contestscoresagents{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

