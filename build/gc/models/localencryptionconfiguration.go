package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocalencryptionconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocalencryptionconfigurationDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Localencryptionconfiguration
type Localencryptionconfiguration struct { 
    


    // Name
    Name string `json:"name"`


    // Url - The url for decryption. This must specify the path to where Purecloud can requests decryption
    Url string `json:"url"`


    // ApiId - The api id for Hawk Authentication.
    ApiId string `json:"apiId"`


    // ApiKey - The api shared symmetric key used for hawk authentication
    ApiKey string `json:"apiKey"`


    

}

// String returns a JSON representation of the model
func (o *Localencryptionconfiguration) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localencryptionconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Localencryptionconfiguration

    if LocalencryptionconfigurationMarshalled {
        return []byte("{}"), nil
    }
    LocalencryptionconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        
        ApiId string `json:"apiId"`
        
        ApiKey string `json:"apiKey"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

