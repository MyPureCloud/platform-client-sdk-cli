package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonebasesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonebasesettingsDud struct { 
    


    


    

}

// Phonebasesettings
type Phonebasesettings struct { 
    // Id - The globally unique identifier for this phone base settings
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Phonebasesettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonebasesettings) MarshalJSON() ([]byte, error) {
    type Alias Phonebasesettings

    if PhonebasesettingsMarshalled {
        return []byte("{}"), nil
    }
    PhonebasesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

