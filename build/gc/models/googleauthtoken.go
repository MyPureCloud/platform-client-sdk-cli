package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogleauthtokenMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogleauthtokenDud struct { 
    Id string `json:"id"`


    ClientId string `json:"clientId"`


    Scopes []string `json:"scopes"`


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Domainentityref `json:"createdBy"`

}

// Googleauthtoken
type Googleauthtoken struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Googleauthtoken) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googleauthtoken) MarshalJSON() ([]byte, error) {
    type Alias Googleauthtoken

    if GoogleauthtokenMarshalled {
        return []byte("{}"), nil
    }
    GoogleauthtokenMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

