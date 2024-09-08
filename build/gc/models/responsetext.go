package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsetextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsetextDud struct { 
    


    


    

}

// Responsetext - Contains information about the text associated with a response.
type Responsetext struct { 
    // Content - Response text content.
    Content string `json:"content"`


    // ContentType - Response text content type.
    ContentType string `json:"contentType"`


    // VarType - Response text type.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Responsetext) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsetext) MarshalJSON() ([]byte, error) {
    type Alias Responsetext

    if ResponsetextMarshalled {
        return []byte("{}"), nil
    }
    ResponsetextMarshalled = true

    return json.Marshal(&struct {
        
        Content string `json:"content"`
        
        ContentType string `json:"contentType"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

