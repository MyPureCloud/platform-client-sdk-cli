package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotnludomainMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotnludomainDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Copilotnludomain
type Copilotnludomain struct { 
    // Id - Id of the NLU domain.
    Id string `json:"id"`


    // UseLatestVersion - Use the latest version of the NLU domain. If false, version is required.
    UseLatestVersion bool `json:"useLatestVersion"`


    // Version - NLU domain version.
    Version Copilotnludomainversion `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Copilotnludomain) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotnludomain) MarshalJSON() ([]byte, error) {
    type Alias Copilotnludomain

    if CopilotnludomainMarshalled {
        return []byte("{}"), nil
    }
    CopilotnludomainMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        UseLatestVersion bool `json:"useLatestVersion"`
        
        Version Copilotnludomainversion `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

