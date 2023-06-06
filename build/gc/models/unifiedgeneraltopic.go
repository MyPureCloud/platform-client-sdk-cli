package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnifiedgeneraltopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnifiedgeneraltopicDud struct { 
    


    

}

// Unifiedgeneraltopic
type Unifiedgeneraltopic struct { 
    // Name
    Name string `json:"name"`


    // Status
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Unifiedgeneraltopic) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unifiedgeneraltopic) MarshalJSON() ([]byte, error) {
    type Alias Unifiedgeneraltopic

    if UnifiedgeneraltopicMarshalled {
        return []byte("{}"), nil
    }
    UnifiedgeneraltopicMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

