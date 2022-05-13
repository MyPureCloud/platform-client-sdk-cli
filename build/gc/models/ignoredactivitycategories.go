package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoredactivitycategoriesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoredactivitycategoriesDud struct { 
    

}

// Ignoredactivitycategories
type Ignoredactivitycategories struct { 
    // Values - Activity categories list
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Ignoredactivitycategories) String() string {
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoredactivitycategories) MarshalJSON() ([]byte, error) {
    type Alias Ignoredactivitycategories

    if IgnoredactivitycategoriesMarshalled {
        return []byte("{}"), nil
    }
    IgnoredactivitycategoriesMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        *Alias
    }{

        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

