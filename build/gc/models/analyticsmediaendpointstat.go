package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsmediaendpointstatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsmediaendpointstatDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsmediaendpointstat
type Analyticsmediaendpointstat struct { 
    // Codecs - The MIME type(s) of the audio encodings used by the audio streams belonging to this endpoint
    Codecs []string `json:"codecs"`


    // DiscardedPackets - The total number of packets received too late or too early, jitter queue overrun or underrun, for all audio streams belonging to this endpoint
    DiscardedPackets int `json:"discardedPackets"`


    // DuplicatePackets - The total number of packets received with the same sequence number as another one recently received (window of 64 packets), for all audio streams belonging to this endpoint
    DuplicatePackets int `json:"duplicatePackets"`


    // EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // InvalidPackets - The total number of malformed or not RTP packets, unknown payload type, or discarded probation packets for all audio streams belonging to this endpoint
    InvalidPackets int `json:"invalidPackets"`


    // MaxLatencyMs - The maximum latency experienced by any audio stream belonging to this endpoint, in milliseconds
    MaxLatencyMs int `json:"maxLatencyMs"`


    // MinMos - The lowest estimated average MOS among all the audio streams belonging to this endpoint
    MinMos float64 `json:"minMos"`


    // MinRFactor - The lowest R-factor value among all of the audio streams belonging to this endpoint
    MinRFactor float64 `json:"minRFactor"`


    // OverrunPackets - The total number of packets for which there was no room in the jitter queue when it was received, for all audio streams belonging to this endpoint (also counted in discarded)
    OverrunPackets int `json:"overrunPackets"`


    // ReceivedPackets - The total number of packets received for all audio streams belonging to this endpoint (includes invalid, duplicate, and discarded packets)
    ReceivedPackets int `json:"receivedPackets"`


    // UnderrunPackets - The total number of packets received after their timestamp/seqnum has been played out, for all audio streams belonging to this endpoint (also counted in discarded)
    UnderrunPackets int `json:"underrunPackets"`

}

// String returns a JSON representation of the model
func (o *Analyticsmediaendpointstat) String() string {
    
    
     o.Codecs = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsmediaendpointstat) MarshalJSON() ([]byte, error) {
    type Alias Analyticsmediaendpointstat

    if AnalyticsmediaendpointstatMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsmediaendpointstatMarshalled = true

    return json.Marshal(&struct { 
        Codecs []string `json:"codecs"`
        
        DiscardedPackets int `json:"discardedPackets"`
        
        DuplicatePackets int `json:"duplicatePackets"`
        
        EventTime time.Time `json:"eventTime"`
        
        InvalidPackets int `json:"invalidPackets"`
        
        MaxLatencyMs int `json:"maxLatencyMs"`
        
        MinMos float64 `json:"minMos"`
        
        MinRFactor float64 `json:"minRFactor"`
        
        OverrunPackets int `json:"overrunPackets"`
        
        ReceivedPackets int `json:"receivedPackets"`
        
        UnderrunPackets int `json:"underrunPackets"`
        
        *Alias
    }{
        

        
        Codecs: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

