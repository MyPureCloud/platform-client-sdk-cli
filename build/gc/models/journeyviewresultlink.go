package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewresultlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewresultlinkDud struct { 
    Id string `json:"id"`


    ConnectionCount int `json:"connectionCount"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewresultlink - Represents a link between 2 elements in a journey view result
type Journeyviewresultlink struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewresultlink) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewresultlink) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewresultlink

    if JourneyviewresultlinkMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewresultlinkMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

