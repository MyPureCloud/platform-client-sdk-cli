package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartmetricresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartmetricresultDud struct { 
    Id string `json:"id"`


    Values []Journeyviewchartmetricresultvalue `json:"values"`

}

// Journeyviewchartmetricresult
type Journeyviewchartmetricresult struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewchartmetricresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartmetricresult) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartmetricresult

    if JourneyviewchartmetricresultMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartmetricresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

