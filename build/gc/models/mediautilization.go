package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediautilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediautilizationDud struct { 
    


    


    

}

// Mediautilization
type Mediautilization struct { 
    // MaximumCapacity - Defines the maximum number of conversations of this type that an agent can handle at one time.
    MaximumCapacity int `json:"maximumCapacity"`


    // InterruptableMediaTypes - Defines the list of other media types that can interrupt a conversation of this media type.  Values include call, chat, email, callback, and message.
    InterruptableMediaTypes []string `json:"interruptableMediaTypes"`


    // IncludeNonAcd - If true, then track non-ACD conversations against utilization
    IncludeNonAcd bool `json:"includeNonAcd"`

}

// String returns a JSON representation of the model
func (o *Mediautilization) String() string {
    
    
    
    
    
    
     o.InterruptableMediaTypes = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediautilization) MarshalJSON() ([]byte, error) {
    type Alias Mediautilization

    if MediautilizationMarshalled {
        return []byte("{}"), nil
    }
    MediautilizationMarshalled = true

    return json.Marshal(&struct { 
        MaximumCapacity int `json:"maximumCapacity"`
        
        InterruptableMediaTypes []string `json:"interruptableMediaTypes"`
        
        IncludeNonAcd bool `json:"includeNonAcd"`
        
        *Alias
    }{
        

        

        

        
        InterruptableMediaTypes: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

