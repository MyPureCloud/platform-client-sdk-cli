package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotnludomainversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotnludomainversionDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Copilotnludomainversion
type Copilotnludomainversion struct { 
    // Id - Id of the NLU v3 domain version.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Copilotnludomainversion) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotnludomainversion) MarshalJSON() ([]byte, error) {
    type Alias Copilotnludomainversion

    if CopilotnludomainversionMarshalled {
        return []byte("{}"), nil
    }
    CopilotnludomainversionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

