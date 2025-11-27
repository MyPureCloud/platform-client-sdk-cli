package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewresultDud struct { 
    Elements []Journeyviewresultelement `json:"elements"`


    Charts []Journeyviewchartresult `json:"charts"`

}

// Journeyviewresult - A journey view result
type Journeyviewresult struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewresult) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewresult

    if JourneyviewresultMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

