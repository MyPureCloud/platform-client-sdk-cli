package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryDud struct { 
    

}

// Category - List of available Action categories.
type Category struct { 
    // Name - Category name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Category) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Category) MarshalJSON() ([]byte, error) {
    type Alias Category

    if CategoryMarshalled {
        return []byte("{}"), nil
    }
    CategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

