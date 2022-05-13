package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserappMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserappDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    IntegrationType Integrationtype `json:"integrationType"`


    


    SelfUri string `json:"selfUri"`

}

// Userapp - Details for a UserApp
type Userapp struct { 
    


    


    


    // Config
    Config Userappconfigurationinfo `json:"config"`


    

}

// String returns a JSON representation of the model
func (o *Userapp) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userapp) MarshalJSON() ([]byte, error) {
    type Alias Userapp

    if UserappMarshalled {
        return []byte("{}"), nil
    }
    UserappMarshalled = true

    return json.Marshal(&struct {
        
        Config Userappconfigurationinfo `json:"config"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

