package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartresultDud struct { 
    Id string `json:"id"`


    Version int `json:"version"`


    Metrics []Journeyviewchartmetricresult `json:"metrics"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewchartresult
type Journeyviewchartresult struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewchartresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartresult) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartresult

    if JourneyviewchartresultMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

