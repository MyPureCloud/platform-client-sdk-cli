package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricsDud struct { 
    


    


    


    


    


    


    


    

}

// Edgemetrics
type Edgemetrics struct { 
    // Edge
    Edge Domainentityref `json:"edge"`


    // EventTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // UpTimeMsec
    UpTimeMsec int `json:"upTimeMsec"`


    // Processors
    Processors []Edgemetricsprocessor `json:"processors"`


    // Memory
    Memory []Edgemetricsmemory `json:"memory"`


    // Disks
    Disks []Edgemetricsdisk `json:"disks"`


    // Subsystems
    Subsystems []Edgemetricssubsystem `json:"subsystems"`


    // Networks
    Networks []Edgemetricsnetwork `json:"networks"`

}

// String returns a JSON representation of the model
func (o *Edgemetrics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Processors = []Edgemetricsprocessor{{}} 
    
    
    
     o.Memory = []Edgemetricsmemory{{}} 
    
    
    
     o.Disks = []Edgemetricsdisk{{}} 
    
    
    
     o.Subsystems = []Edgemetricssubsystem{{}} 
    
    
    
     o.Networks = []Edgemetricsnetwork{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetrics) MarshalJSON() ([]byte, error) {
    type Alias Edgemetrics

    if EdgemetricsMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricsMarshalled = true

    return json.Marshal(&struct { 
        Edge Domainentityref `json:"edge"`
        
        EventTime time.Time `json:"eventTime"`
        
        UpTimeMsec int `json:"upTimeMsec"`
        
        Processors []Edgemetricsprocessor `json:"processors"`
        
        Memory []Edgemetricsmemory `json:"memory"`
        
        Disks []Edgemetricsdisk `json:"disks"`
        
        Subsystems []Edgemetricssubsystem `json:"subsystems"`
        
        Networks []Edgemetricsnetwork `json:"networks"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Processors: []Edgemetricsprocessor{{}},
        

        

        
        Memory: []Edgemetricsmemory{{}},
        

        

        
        Disks: []Edgemetricsdisk{{}},
        

        

        
        Subsystems: []Edgemetricssubsystem{{}},
        

        

        
        Networks: []Edgemetricsnetwork{{}},
        

        
        Alias: (*Alias)(u),
    })
}

