package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaiceselectedpairMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaiceselectedpairDud struct { 
    


    


    

}

// Mediaiceselectedpair
type Mediaiceselectedpair struct { 
    // Client - The remote candidate that was chosen
    Client Mediaiceselectedcandidate `json:"client"`


    // Server - The local candidate that was chosen
    Server Mediaiceselectedcandidate `json:"server"`


    // CandidatePairSelectedMilliseconds - Relative milliseconds since creation of endpoint when this ICE candidate pair has been selected
    CandidatePairSelectedMilliseconds int `json:"candidatePairSelectedMilliseconds"`

}

// String returns a JSON representation of the model
func (o *Mediaiceselectedpair) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaiceselectedpair) MarshalJSON() ([]byte, error) {
    type Alias Mediaiceselectedpair

    if MediaiceselectedpairMarshalled {
        return []byte("{}"), nil
    }
    MediaiceselectedpairMarshalled = true

    return json.Marshal(&struct {
        
        Client Mediaiceselectedcandidate `json:"client"`
        
        Server Mediaiceselectedcandidate `json:"server"`
        
        CandidatePairSelectedMilliseconds int `json:"candidatePairSelectedMilliseconds"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

