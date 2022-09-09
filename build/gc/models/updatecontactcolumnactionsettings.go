package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatecontactcolumnactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatecontactcolumnactionsettingsDud struct { 
    


    

}

// Updatecontactcolumnactionsettings
type Updatecontactcolumnactionsettings struct { 
    // Properties - A mapping of contact columns to their new values.
    Properties map[string]string `json:"properties"`


    // UpdateOption - The type of update to make to the specified contact column(s).
    UpdateOption string `json:"updateOption"`

}

// String returns a JSON representation of the model
func (o *Updatecontactcolumnactionsettings) String() string {
     o.Properties = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatecontactcolumnactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Updatecontactcolumnactionsettings

    if UpdatecontactcolumnactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    UpdatecontactcolumnactionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Properties map[string]string `json:"properties"`
        
        UpdateOption string `json:"updateOption"`
        *Alias
    }{

        
        Properties: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

