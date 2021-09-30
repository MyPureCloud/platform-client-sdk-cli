package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresponsestatsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresponsestatsDud struct { 
    

}

// Queryresponsestats
type Queryresponsestats struct { 
    // Count - The count for this metric
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Queryresponsestats) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresponsestats) MarshalJSON() ([]byte, error) {
    type Alias Queryresponsestats

    if QueryresponsestatsMarshalled {
        return []byte("{}"), nil
    }
    QueryresponsestatsMarshalled = true

    return json.Marshal(&struct { 
        Count int `json:"count"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

