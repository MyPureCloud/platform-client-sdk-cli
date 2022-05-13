package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserlicensesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserlicensesDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Userlicenses
type Userlicenses struct { 
    


    // Licenses
    Licenses []string `json:"licenses"`


    

}

// String returns a JSON representation of the model
func (o *Userlicenses) String() string {
     o.Licenses = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userlicenses) MarshalJSON() ([]byte, error) {
    type Alias Userlicenses

    if UserlicensesMarshalled {
        return []byte("{}"), nil
    }
    UserlicensesMarshalled = true

    return json.Marshal(&struct {
        
        Licenses []string `json:"licenses"`
        *Alias
    }{

        


        
        Licenses: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

