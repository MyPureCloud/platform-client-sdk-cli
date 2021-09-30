package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenttimeoffrequestpatchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenttimeoffrequestpatchDud struct { 
    


    


    

}

// Agenttimeoffrequestpatch
type Agenttimeoffrequestpatch struct { 
    // MarkedAsRead - Whether this request has been read by the agent
    MarkedAsRead bool `json:"markedAsRead"`


    // Status - The status of this time off request. Can only be canceled if the requested date has not already passed
    Status string `json:"status"`


    // Notes - Notes about the time off request. Can only be edited while the request is still pending
    Notes string `json:"notes"`

}

// String returns a JSON representation of the model
func (o *Agenttimeoffrequestpatch) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenttimeoffrequestpatch) MarshalJSON() ([]byte, error) {
    type Alias Agenttimeoffrequestpatch

    if AgenttimeoffrequestpatchMarshalled {
        return []byte("{}"), nil
    }
    AgenttimeoffrequestpatchMarshalled = true

    return json.Marshal(&struct { 
        MarkedAsRead bool `json:"markedAsRead"`
        
        Status string `json:"status"`
        
        Notes string `json:"notes"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

