package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewuserDud struct { 
    Id string `json:"id"`


    EmailAddress string `json:"emailAddress"`


    SelfUri string `json:"selfUri"`

}

// Journeyviewuser
type Journeyviewuser struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewuser) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewuser) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewuser

    if JourneyviewuserMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewuserMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

