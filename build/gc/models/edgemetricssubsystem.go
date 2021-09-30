package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricssubsystemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricssubsystemDud struct { 
    


    


    

}

// Edgemetricssubsystem
type Edgemetricssubsystem struct { 
    // DelayMs - Delay in milliseconds.
    DelayMs int `json:"delayMs"`


    // ProcessName - Name of the Edge process.
    ProcessName string `json:"processName"`


    // MediaSubsystem - Subsystem for an Edge device.
    MediaSubsystem *Edgemetricssubsystem `json:"mediaSubsystem"`

}

// String returns a JSON representation of the model
func (o *Edgemetricssubsystem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetricssubsystem) MarshalJSON() ([]byte, error) {
    type Alias Edgemetricssubsystem

    if EdgemetricssubsystemMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricssubsystemMarshalled = true

    return json.Marshal(&struct { 
        DelayMs int `json:"delayMs"`
        
        ProcessName string `json:"processName"`
        
        MediaSubsystem *Edgemetricssubsystem `json:"mediaSubsystem"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

