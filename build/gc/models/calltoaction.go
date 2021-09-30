package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CalltoactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CalltoactionDud struct { 
    


    


    

}

// Calltoaction
type Calltoaction struct { 
    // Text - Text displayed on the call to action button.
    Text string `json:"text"`


    // Url - URL to open when user clicks on the call to action button.
    Url string `json:"url"`


    // Target - Where the URL should be opened when the user clicks on the call to action button.
    Target string `json:"target"`

}

// String returns a JSON representation of the model
func (o *Calltoaction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Calltoaction) MarshalJSON() ([]byte, error) {
    type Alias Calltoaction

    if CalltoactionMarshalled {
        return []byte("{}"), nil
    }
    CalltoactionMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        Url string `json:"url"`
        
        Target string `json:"target"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

