package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MarkcontactuncontactableactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MarkcontactuncontactableactionsettingsDud struct { 
    

}

// Markcontactuncontactableactionsettings
type Markcontactuncontactableactionsettings struct { 
    // MediaTypes - A list of media types to evaluate.
    MediaTypes []string `json:"mediaTypes"`

}

// String returns a JSON representation of the model
func (o *Markcontactuncontactableactionsettings) String() string {
     o.MediaTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Markcontactuncontactableactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Markcontactuncontactableactionsettings

    if MarkcontactuncontactableactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    MarkcontactuncontactableactionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MediaTypes []string `json:"mediaTypes"`
        *Alias
    }{

        
        MediaTypes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

