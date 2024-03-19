package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewresultmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewresultmetricsDud struct { 
    ParticipantCount int `json:"participantCount"`


    ActiveCount int `json:"activeCount"`


    CompletedCount int `json:"completedCount"`


    DropoutCount int `json:"dropoutCount"`


    FlowCount int `json:"flowCount"`

}

// Journeyviewresultmetrics - The metrics of an element or a link in journey 
type Journeyviewresultmetrics struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewresultmetrics) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewresultmetrics) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewresultmetrics

    if JourneyviewresultmetricsMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewresultmetricsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

