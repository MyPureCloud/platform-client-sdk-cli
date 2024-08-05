package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementfilterrangedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementfilterrangedataDud struct { 
    


    

}

// Journeyviewelementfilterrangedata
type Journeyviewelementfilterrangedata struct { 
    // Duration - an ISO 8601 time duration.Only one of number or duration must be specified. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Duration string `json:"duration"`


    // Number - an Integer value.Only one of number or duration must be specified.
    Number int `json:"number"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementfilterrangedata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementfilterrangedata) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementfilterrangedata

    if JourneyviewelementfilterrangedataMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementfilterrangedataMarshalled = true

    return json.Marshal(&struct {
        
        Duration string `json:"duration"`
        
        Number int `json:"number"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

