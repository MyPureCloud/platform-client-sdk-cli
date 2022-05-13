package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricscallsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricscallsDud struct { 
    


    

}

// Trunkmetricscalls
type Trunkmetricscalls struct { 
    // InboundCallCount
    InboundCallCount int `json:"inboundCallCount"`


    // OutboundCallCount
    OutboundCallCount int `json:"outboundCallCount"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricscalls) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetricscalls) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetricscalls

    if TrunkmetricscallsMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricscallsMarshalled = true

    return json.Marshal(&struct {
        
        InboundCallCount int `json:"inboundCallCount"`
        
        OutboundCallCount int `json:"outboundCallCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

