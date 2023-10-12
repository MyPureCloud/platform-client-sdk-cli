package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaiceselectedcandidateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaiceselectedcandidateDud struct { 
    


    

}

// Mediaiceselectedcandidate
type Mediaiceselectedcandidate struct { 
    // Address - IP address and port of the candidate
    Address string `json:"address"`


    // VarType - Type of the selected candidate
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Mediaiceselectedcandidate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaiceselectedcandidate) MarshalJSON() ([]byte, error) {
    type Alias Mediaiceselectedcandidate

    if MediaiceselectedcandidateMarshalled {
        return []byte("{}"), nil
    }
    MediaiceselectedcandidateMarshalled = true

    return json.Marshal(&struct {
        
        Address string `json:"address"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

