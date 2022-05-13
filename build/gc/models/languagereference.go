package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LanguagereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LanguagereferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Languagereference
type Languagereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Languagereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Languagereference) MarshalJSON() ([]byte, error) {
    type Alias Languagereference

    if LanguagereferenceMarshalled {
        return []byte("{}"), nil
    }
    LanguagereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

