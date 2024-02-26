package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewDud struct { 
    Id string `json:"id"`


    


    


    Version int `json:"version"`


    CreatedBy Journeyviewuser `json:"createdBy"`


    ModifiedBy Journeyviewuser `json:"modifiedBy"`


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Journeyview - A journey view
type Journeyview struct { 
    


    // Name
    Name string `json:"name"`


    // Description - A description of the journey view
    Description string `json:"description"`


    


    


    


    // Interval - An absolute timeframe for the journey view, expressed as an ISO 8601 interval. Only one of interval or duration must be specified. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Duration - A relative timeframe for the journey view, expressed as an ISO 8601 duration. Only one of interval or duration must be specified. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Duration string `json:"duration"`


    // Elements - The elements within the journey view
    Elements []Journeyviewelement `json:"elements"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyview) String() string {
    
    
    
    
     o.Elements = []Journeyviewelement{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyview) MarshalJSON() ([]byte, error) {
    type Alias Journeyview

    if JourneyviewMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Interval string `json:"interval"`
        
        Duration string `json:"duration"`
        
        Elements []Journeyviewelement `json:"elements"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Elements: []Journeyviewelement{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

