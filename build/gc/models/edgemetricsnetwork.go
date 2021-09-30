package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricsnetworkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricsnetworkDud struct { 
    


    


    


    


    

}

// Edgemetricsnetwork
type Edgemetricsnetwork struct { 
    // Ifname - Identifier for the network adapter.
    Ifname string `json:"ifname"`


    // SentBytesPerSec - Number of byes sent per second.
    SentBytesPerSec int `json:"sentBytesPerSec"`


    // ReceivedBytesPerSec - Number of byes received per second.
    ReceivedBytesPerSec int `json:"receivedBytesPerSec"`


    // BandwidthBitsPerSec - Total bandwidth of the adapter in bits per second.
    BandwidthBitsPerSec float64 `json:"bandwidthBitsPerSec"`


    // UtilizationPct - Percent utilization of the network adapter.
    UtilizationPct float64 `json:"utilizationPct"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsnetwork) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetricsnetwork) MarshalJSON() ([]byte, error) {
    type Alias Edgemetricsnetwork

    if EdgemetricsnetworkMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricsnetworkMarshalled = true

    return json.Marshal(&struct { 
        Ifname string `json:"ifname"`
        
        SentBytesPerSec int `json:"sentBytesPerSec"`
        
        ReceivedBytesPerSec int `json:"receivedBytesPerSec"`
        
        BandwidthBitsPerSec float64 `json:"bandwidthBitsPerSec"`
        
        UtilizationPct float64 `json:"utilizationPct"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

