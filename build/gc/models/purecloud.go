package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PurecloudMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PurecloudDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Purecloud
type Purecloud struct { 
    


    // Name
    Name string `json:"name"`


    // Disabled
    Disabled bool `json:"disabled"`


    

}

// String returns a JSON representation of the model
func (o *Purecloud) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Purecloud) MarshalJSON() ([]byte, error) {
    type Alias Purecloud

    if PurecloudMarshalled {
        return []byte("{}"), nil
    }
    PurecloudMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Disabled bool `json:"disabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

