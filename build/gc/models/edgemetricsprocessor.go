package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgemetricsprocessorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgemetricsprocessorDud struct { 
    


    


    


    


    

}

// Edgemetricsprocessor
type Edgemetricsprocessor struct { 
    // ActiveTimePct - Percent time processor was active.
    ActiveTimePct float64 `json:"activeTimePct"`


    // CpuId - Machine CPU identifier. 'total' will always be included in the array and is the total of all CPU resources.
    CpuId string `json:"cpuId"`


    // IdleTimePct - Percent time processor was idle.
    IdleTimePct float64 `json:"idleTimePct"`


    // PrivilegedTimePct - Percent time processor spent in privileged mode.
    PrivilegedTimePct float64 `json:"privilegedTimePct"`


    // UserTimePct - Percent time processor spent in user mode.
    UserTimePct float64 `json:"userTimePct"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsprocessor) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgemetricsprocessor) MarshalJSON() ([]byte, error) {
    type Alias Edgemetricsprocessor

    if EdgemetricsprocessorMarshalled {
        return []byte("{}"), nil
    }
    EdgemetricsprocessorMarshalled = true

    return json.Marshal(&struct { 
        ActiveTimePct float64 `json:"activeTimePct"`
        
        CpuId string `json:"cpuId"`
        
        IdleTimePct float64 `json:"idleTimePct"`
        
        PrivilegedTimePct float64 `json:"privilegedTimePct"`
        
        UserTimePct float64 `json:"userTimePct"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

