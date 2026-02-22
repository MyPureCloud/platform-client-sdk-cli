package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogleauthtokenrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogleauthtokenrequestDud struct { 
    


    

}

// Googleauthtokenrequest
type Googleauthtokenrequest struct { 
    // Code - One-time authorization code from Google. See docs: https://developers.google.com/identity/protocols/oauth2
    Code string `json:"code"`


    // RedirectUri - The origin of the page from which the auth flow was called. Used as redirectUri when exchanging the one-time code for auth token
    RedirectUri string `json:"redirectUri"`

}

// String returns a JSON representation of the model
func (o *Googleauthtokenrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googleauthtokenrequest) MarshalJSON() ([]byte, error) {
    type Alias Googleauthtokenrequest

    if GoogleauthtokenrequestMarshalled {
        return []byte("{}"), nil
    }
    GoogleauthtokenrequestMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        RedirectUri string `json:"redirectUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

