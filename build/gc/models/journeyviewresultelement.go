package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewresultelementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewresultelementDud struct { 
    Id string `json:"id"`


    Metrics Journeyviewresultmetrics `json:"metrics"`


    FollowedBy []Journeyviewresultlink `json:"followedBy"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewresultelement - An element within a journey view result
type Journeyviewresultelement struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewresultelement) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewresultelement) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewresultelement

    if JourneyviewresultelementMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewresultelementMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

