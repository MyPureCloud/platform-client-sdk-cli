package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticationpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticationpropertiesDud struct { 
    

}

// Authenticationproperties
type Authenticationproperties struct { 
    // Sharepoint
    Sharepoint Sharepointauthenticationproperties `json:"sharepoint"`

}

// String returns a JSON representation of the model
func (o *Authenticationproperties) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticationproperties) MarshalJSON() ([]byte, error) {
    type Alias Authenticationproperties

    if AuthenticationpropertiesMarshalled {
        return []byte("{}"), nil
    }
    AuthenticationpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Sharepoint Sharepointauthenticationproperties `json:"sharepoint"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

