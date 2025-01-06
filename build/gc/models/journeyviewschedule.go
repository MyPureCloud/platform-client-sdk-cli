package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewscheduleDud struct { 
    Id string `json:"id"`


    


    DateModified time.Time `json:"dateModified"`


    User Addressableentityref `json:"user"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewschedule
type Journeyviewschedule struct { 
    


    // Frequency - Frequency of execution
    Frequency string `json:"frequency"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewschedule) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewschedule) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewschedule

    if JourneyviewscheduleMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Frequency string `json:"frequency"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

