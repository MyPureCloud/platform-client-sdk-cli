package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediartpstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediartpstatisticsDud struct { 
    


    


    


    


    

}

// Mediartpstatistics
type Mediartpstatistics struct { 
    // PacketsReceived - Number of packets received, including all invalid, duplicate, and discarded packets
    PacketsReceived int `json:"packetsReceived"`


    // PacketsSent - Number of packets sent
    PacketsSent int `json:"packetsSent"`


    // RtpEventsReceived - Number of RFC#2833 packets received
    RtpEventsReceived int `json:"rtpEventsReceived"`


    // RtpEventsSent - Number of RFC#2833 packets sent
    RtpEventsSent int `json:"rtpEventsSent"`


    // EstimatedAverageMos - The estimated average MOS score
    EstimatedAverageMos float64 `json:"estimatedAverageMos"`

}

// String returns a JSON representation of the model
func (o *Mediartpstatistics) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediartpstatistics) MarshalJSON() ([]byte, error) {
    type Alias Mediartpstatistics

    if MediartpstatisticsMarshalled {
        return []byte("{}"), nil
    }
    MediartpstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        PacketsReceived int `json:"packetsReceived"`
        
        PacketsSent int `json:"packetsSent"`
        
        RtpEventsReceived int `json:"rtpEventsReceived"`
        
        RtpEventsSent int `json:"rtpEventsSent"`
        
        EstimatedAverageMos float64 `json:"estimatedAverageMos"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

