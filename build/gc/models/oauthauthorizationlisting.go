package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthauthorizationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthauthorizationlistingDud struct { 
    


    


    

}

// Oauthauthorizationlisting
type Oauthauthorizationlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Oauthauthorization `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Oauthauthorizationlisting) String() string {
    
     o.Entities = []Oauthauthorization{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthauthorizationlisting) MarshalJSON() ([]byte, error) {
    type Alias Oauthauthorizationlisting

    if OauthauthorizationlistingMarshalled {
        return []byte("{}"), nil
    }
    OauthauthorizationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Oauthauthorization `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Oauthauthorization{{}},
        


        

        Alias: (*Alias)(u),
    })
}

