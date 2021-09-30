package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatguestmediarequestentitylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatguestmediarequestentitylistDud struct { 
    

}

// Webchatguestmediarequestentitylist
type Webchatguestmediarequestentitylist struct { 
    // Entities
    Entities []Webchatguestmediarequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequestentitylist) String() string {
    
    
     o.Entities = []Webchatguestmediarequest{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatguestmediarequestentitylist) MarshalJSON() ([]byte, error) {
    type Alias Webchatguestmediarequestentitylist

    if WebchatguestmediarequestentitylistMarshalled {
        return []byte("{}"), nil
    }
    WebchatguestmediarequestentitylistMarshalled = true

    return json.Marshal(&struct { 
        Entities []Webchatguestmediarequest `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Webchatguestmediarequest{{}},
        

        
        Alias: (*Alias)(u),
    })
}

