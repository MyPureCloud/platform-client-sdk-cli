package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestrequesingparticipantdailyinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestrequesingparticipantdailyinfoDud struct { 
    


    

}

// Contestrequesingparticipantdailyinfo
type Contestrequesingparticipantdailyinfo struct { 
    // DateWorkday - Workday of the contest info. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateWorkday time.Time `json:"dateWorkday"`


    // ContestScore - The Contest score
    ContestScore Contestscoreranked `json:"contestScore"`

}

// String returns a JSON representation of the model
func (o *Contestrequesingparticipantdailyinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestrequesingparticipantdailyinfo) MarshalJSON() ([]byte, error) {
    type Alias Contestrequesingparticipantdailyinfo

    if ContestrequesingparticipantdailyinfoMarshalled {
        return []byte("{}"), nil
    }
    ContestrequesingparticipantdailyinfoMarshalled = true

    return json.Marshal(&struct {
        
        DateWorkday time.Time `json:"dateWorkday"`
        
        ContestScore Contestscoreranked `json:"contestScore"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

