package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustorDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Orguser `json:"createdBy"`


    Organization Organization `json:"organization"`


    Authorization Trusteeauthorization `json:"authorization"`


    SelfUri string `json:"selfUri"`

}

// Trustor
type Trustor struct { 
    


    // Enabled - If disabled no trustee user will have access, even if they were previously added.
    Enabled bool `json:"enabled"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Trustor) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustor) MarshalJSON() ([]byte, error) {
    type Alias Trustor

    if TrustorMarshalled {
        return []byte("{}"), nil
    }
    TrustorMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

