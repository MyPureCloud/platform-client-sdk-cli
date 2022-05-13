package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SortitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SortitemDud struct { 
    


    

}

// Sortitem
type Sortitem struct { 
    // Name
    Name string `json:"name"`


    // Ascending
    Ascending bool `json:"ascending"`

}

// String returns a JSON representation of the model
func (o *Sortitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sortitem) MarshalJSON() ([]byte, error) {
    type Alias Sortitem

    if SortitemMarshalled {
        return []byte("{}"), nil
    }
    SortitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Ascending bool `json:"ascending"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

