package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowmilestoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowmilestoneDud struct { 
    


    


    

}

// Textbotflowmilestone
type Textbotflowmilestone struct { 
    // Id - The Milestone's ID.
    Id string `json:"id"`


    // DateReached - The timestamp of when the milestone was reached. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateReached time.Time `json:"dateReached"`


    // Sequence - The sequence number of the milestone.
    Sequence int `json:"sequence"`

}

// String returns a JSON representation of the model
func (o *Textbotflowmilestone) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowmilestone) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowmilestone

    if TextbotflowmilestoneMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowmilestoneMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DateReached time.Time `json:"dateReached"`
        
        Sequence int `json:"sequence"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

