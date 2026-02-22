package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogleoauthsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogleoauthsettingsresponseDud struct { 
    ClientId string `json:"clientId"`


    Scopes []string `json:"scopes"`

}

// Googleoauthsettingsresponse
type Googleoauthsettingsresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Googleoauthsettingsresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googleoauthsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Googleoauthsettingsresponse

    if GoogleoauthsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    GoogleoauthsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

