package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewjobDud struct { 
    Id string `json:"id"`


    DateCreated time.Time `json:"dateCreated"`


    DateCompleted time.Time `json:"dateCompleted"`


    


    


    DateCompletionEstimated time.Time `json:"dateCompletionEstimated"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewjob
type Journeyviewjob struct { 
    


    


    


    // Status - The status of the job
    Status string `json:"status"`


    // JourneyView - The journey view for which the job is executed
    JourneyView Journeyview `json:"journeyView"`


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewjob) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewjob) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewjob

    if JourneyviewjobMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewjobMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        JourneyView Journeyview `json:"journeyView"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

