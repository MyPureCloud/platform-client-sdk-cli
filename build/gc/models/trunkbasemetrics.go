package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkbasemetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkbasemetricsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Trunkbasemetrics
type Trunkbasemetrics struct { 
    


    // Name
    Name string `json:"name"`


    // TrunkBase - TrunkBase resource identifier
    TrunkBase Addressableentityref `json:"trunkBase"`


    // TrunkType - The type of the TrunkBase
    TrunkType string `json:"trunkType"`


    // InboundCallCount - The number of inbound calls
    InboundCallCount int `json:"inboundCallCount"`


    // OutboundCallCount - The number of outbound calls
    OutboundCallCount int `json:"outboundCallCount"`


    // EventTime - The timestamp of the data publication. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // QosMismatchCount - The number of QoS mismatch events
    QosMismatchCount int `json:"qosMismatchCount"`


    

}

// String returns a JSON representation of the model
func (o *Trunkbasemetrics) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkbasemetrics) MarshalJSON() ([]byte, error) {
    type Alias Trunkbasemetrics

    if TrunkbasemetricsMarshalled {
        return []byte("{}"), nil
    }
    TrunkbasemetricsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        TrunkBase Addressableentityref `json:"trunkBase"`
        
        TrunkType string `json:"trunkType"`
        
        InboundCallCount int `json:"inboundCallCount"`
        
        OutboundCallCount int `json:"outboundCallCount"`
        
        EventTime time.Time `json:"eventTime"`
        
        QosMismatchCount int `json:"qosMismatchCount"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

