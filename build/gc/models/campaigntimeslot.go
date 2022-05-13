package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaigntimeslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaigntimeslotDud struct { 
    


    


    

}

// Campaigntimeslot
type Campaigntimeslot struct { 
    // StartTime - The start time of the interval as an ISO-8601 string, i.e. HH:mm:ss
    StartTime string `json:"startTime"`


    // StopTime - The end time of the interval as an ISO-8601 string, i.e. HH:mm:ss
    StopTime string `json:"stopTime"`


    // Day - The day of the interval. Valid values: [1-7], representing Monday through Sunday
    Day int `json:"day"`

}

// String returns a JSON representation of the model
func (o *Campaigntimeslot) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaigntimeslot) MarshalJSON() ([]byte, error) {
    type Alias Campaigntimeslot

    if CampaigntimeslotMarshalled {
        return []byte("{}"), nil
    }
    CampaigntimeslotMarshalled = true

    return json.Marshal(&struct {
        
        StartTime string `json:"startTime"`
        
        StopTime string `json:"stopTime"`
        
        Day int `json:"day"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

