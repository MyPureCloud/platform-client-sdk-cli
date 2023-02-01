package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CobrowsewebmessagingsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CobrowsewebmessagingsessionDud struct { 
    Id string `json:"id"`


    


    JoinCode string `json:"joinCode"`


    WebsocketUrl string `json:"websocketUrl"`


    DateOfferEnds time.Time `json:"dateOfferEnds"`


    CommunicationType string `json:"communicationType"`


    SelfUri string `json:"selfUri"`

}

// Cobrowsewebmessagingsession
type Cobrowsewebmessagingsession struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Cobrowsewebmessagingsession) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cobrowsewebmessagingsession) MarshalJSON() ([]byte, error) {
    type Alias Cobrowsewebmessagingsession

    if CobrowsewebmessagingsessionMarshalled {
        return []byte("{}"), nil
    }
    CobrowsewebmessagingsessionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

