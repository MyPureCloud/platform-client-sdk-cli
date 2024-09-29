package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InternalmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InternalmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Internalmessage
type Internalmessage struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // InitialState - The initial connection state of this communication.
    InitialState string `json:"initialState"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Segments - The time line of the participant's internal message, divided into activity segments.
    Segments []Segment `json:"segments"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // Provider - The source provider for the message.
    Provider string `json:"provider"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // TargetUserId - The user ID for the participant on receiving side of the internal message conversation.
    TargetUserId string `json:"targetUserId"`


    // SourceUserId - The user ID for the participant on sending side of the internal message conversation.
    SourceUserId string `json:"sourceUserId"`


    // ToAddress - Address for the participant on receiving side of the internal message communication.
    ToAddress Address `json:"toAddress"`


    // FromAddress - Address for the participant on the sending side of the internal message communication.
    FromAddress Address `json:"fromAddress"`


    // Messages - The messages sent on this communication channel.
    Messages []Internalmessagedetails `json:"messages"`

}

// String returns a JSON representation of the model
func (o *Internalmessage) String() string {
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
    
    
     o.Messages = []Internalmessagedetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Internalmessage) MarshalJSON() ([]byte, error) {
    type Alias Internalmessage

    if InternalmessageMarshalled {
        return []byte("{}"), nil
    }
    InternalmessageMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        InitialState string `json:"initialState"`
        
        Id string `json:"id"`
        
        Segments []Segment `json:"segments"`
        
        DisconnectType string `json:"disconnectType"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        Provider string `json:"provider"`
        
        PeerId string `json:"peerId"`
        
        TargetUserId string `json:"targetUserId"`
        
        SourceUserId string `json:"sourceUserId"`
        
        ToAddress Address `json:"toAddress"`
        
        FromAddress Address `json:"fromAddress"`
        
        Messages []Internalmessagedetails `json:"messages"`
        *Alias
    }{

        


        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        


        


        


        
        Messages: []Internalmessagedetails{{}},
        

        Alias: (*Alias)(u),
    })
}

