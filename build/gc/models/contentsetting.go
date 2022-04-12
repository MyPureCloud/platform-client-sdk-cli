package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentsettingDud struct { 
    

}

// Contentsetting
type Contentsetting struct { 
    // Story - Settings relating to facebook and instagram stories feature
    Story Storysetting `json:"story"`

}

// String returns a JSON representation of the model
func (o *Contentsetting) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentsetting) MarshalJSON() ([]byte, error) {
    type Alias Contentsetting

    if ContentsettingMarshalled {
        return []byte("{}"), nil
    }
    ContentsettingMarshalled = true

    return json.Marshal(&struct { 
        Story Storysetting `json:"story"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

