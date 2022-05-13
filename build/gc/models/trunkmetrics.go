package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricsDud struct { 
    


    


    


    


    

}

// Trunkmetrics
type Trunkmetrics struct { 
    // EventTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // LogicalInterface
    LogicalInterface Domainentityref `json:"logicalInterface"`


    // Trunk
    Trunk Domainentityref `json:"trunk"`


    // Calls
    Calls Trunkmetricscalls `json:"calls"`


    // Qos
    Qos Trunkmetricsqos `json:"qos"`

}

// String returns a JSON representation of the model
func (o *Trunkmetrics) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetrics) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetrics

    if TrunkmetricsMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricsMarshalled = true

    return json.Marshal(&struct {
        
        EventTime time.Time `json:"eventTime"`
        
        LogicalInterface Domainentityref `json:"logicalInterface"`
        
        Trunk Domainentityref `json:"trunk"`
        
        Calls Trunkmetricscalls `json:"calls"`
        
        Qos Trunkmetricsqos `json:"qos"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

