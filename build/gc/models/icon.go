package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IconMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IconDud struct { 
    

}

// Icon - The icon for the launcher button
type Icon struct { 
    // Url - Icon URL for the launcher button
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Icon) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Icon) MarshalJSON() ([]byte, error) {
    type Alias Icon

    if IconMarshalled {
        return []byte("{}"), nil
    }
    IconMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

