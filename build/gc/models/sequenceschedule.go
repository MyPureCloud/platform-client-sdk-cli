package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SequencescheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SequencescheduleDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Sequenceschedule
type Sequenceschedule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Intervals - A list of intervals during which to run the associated CampaignSequence.
    Intervals []Scheduleinterval `json:"intervals"`


    // TimeZone - The time zone for this SequenceSchedule. For example, Africa/Abidjan.
    TimeZone string `json:"timeZone"`


    // Sequence - The CampaignSequence that this SequenceSchedule is for.
    Sequence Domainentityref `json:"sequence"`


    

}

// String returns a JSON representation of the model
func (o *Sequenceschedule) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Intervals = []Scheduleinterval{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sequenceschedule) MarshalJSON() ([]byte, error) {
    type Alias Sequenceschedule

    if SequencescheduleMarshalled {
        return []byte("{}"), nil
    }
    SequencescheduleMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        Version int `json:"version"`
        
        Intervals []Scheduleinterval `json:"intervals"`
        
        TimeZone string `json:"timeZone"`
        
        Sequence Domainentityref `json:"sequence"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Intervals: []Scheduleinterval{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

