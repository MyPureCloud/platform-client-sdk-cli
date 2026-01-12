package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Screenmonitoring
type Screenmonitoring struct { 
    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // InitialState - The initial connection state of this communication.
    InitialState string `json:"initialState"`


    // State - The connection state of this communication.
    State string `json:"state"`


    // Segments - The time line of the participant's Screen Monitoring media, divided into activity segments.
    Segments []Segment `json:"segments"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // Provider - The source provider of Screen Monitoring media.
    Provider string `json:"provider"`


    // TargetUser - The user participant who is being screen monitored.
    TargetUser Addressableentityref `json:"targetUser"`


    // ScreenMonitoring - Screen Monitoring identifier unique to the sourceUserId-targetUserId pair.
    ScreenMonitoring Addressableentityref `json:"screenMonitoring"`


    // MonitoringType - The screen monitoring type.
    MonitoringType string `json:"monitoringType"`


    // Count - Number of Screen Monitoring sessions the targetUserId is involved in.
    Count int `json:"count"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`

}

// String returns a JSON representation of the model
func (o *Screenmonitoring) String() string {
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoring) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoring

    if ScreenmonitoringMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        InitialState string `json:"initialState"`
        
        State string `json:"state"`
        
        Segments []Segment `json:"segments"`
        
        DisconnectType string `json:"disconnectType"`
        
        Provider string `json:"provider"`
        
        TargetUser Addressableentityref `json:"targetUser"`
        
        ScreenMonitoring Addressableentityref `json:"screenMonitoring"`
        
        MonitoringType string `json:"monitoringType"`
        
        Count int `json:"count"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        *Alias
    }{

        


        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

