package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthscopeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthscopeDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Oauthscope
type Oauthscope struct { 
    


    // Description
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Oauthscope) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthscope) MarshalJSON() ([]byte, error) {
    type Alias Oauthscope

    if OauthscopeMarshalled {
        return []byte("{}"), nil
    }
    OauthscopeMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

