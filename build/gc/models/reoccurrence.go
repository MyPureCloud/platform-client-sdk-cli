package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReoccurrenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReoccurrenceDud struct { 
    


    


    


    


    


    


    

}

// Reoccurrence
type Reoccurrence struct { 
    // Id
    Id string `json:"id"`


    // Start - The  start date time of the initial occurrence as an ISO-8601 string in UTC time, e.g: 2023-11-21T16:30:25.000Z
    Start string `json:"start"`


    // End - The end date time of the initial occurrence as an ISO-8601 string in UTC time, e.g: 2023-12-21T16:30:25.000Z
    End string `json:"end"`


    // TimeZone - The time zone of the schedule e.g.:  America/New_York
    TimeZone string `json:"timeZone"`


    // Pattern - The schedule pattern e.g.: Daily/Weekly
    Pattern Pattern `json:"pattern"`


    // VarRange - The schedule range e.g.: EndDate/NoEnd/Numbered
    VarRange Range `json:"range"`


    // Alterations - Modifications to the original recurrence schedule (Exclusions/Inclusions)
    Alterations []Alteration `json:"alterations"`

}

// String returns a JSON representation of the model
func (o *Reoccurrence) String() string {
    
    
    
    
    
    
     o.Alterations = []Alteration{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reoccurrence) MarshalJSON() ([]byte, error) {
    type Alias Reoccurrence

    if ReoccurrenceMarshalled {
        return []byte("{}"), nil
    }
    ReoccurrenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Start string `json:"start"`
        
        End string `json:"end"`
        
        TimeZone string `json:"timeZone"`
        
        Pattern Pattern `json:"pattern"`
        
        VarRange Range `json:"range"`
        
        Alterations []Alteration `json:"alterations"`
        *Alias
    }{

        


        


        


        


        


        


        
        Alterations: []Alteration{{}},
        

        Alias: (*Alias)(u),
    })
}

