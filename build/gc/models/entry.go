package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntryDud struct { 
    


    

}

// Entry
type Entry struct { 
    // Value - A value included in this facet.
    Value string `json:"value"`


    // Count - The number of results with this value.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Entry) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entry) MarshalJSON() ([]byte, error) {
    type Alias Entry

    if EntryMarshalled {
        return []byte("{}"), nil
    }
    EntryMarshalled = true

    return json.Marshal(&struct { 
        Value string `json:"value"`
        
        Count int `json:"count"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

