package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthscopelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthscopelistingDud struct { 
    


    


    

}

// Oauthscopelisting
type Oauthscopelisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Oauthscope `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Oauthscopelisting) String() string {
    
    
    
    
    
    
     o.Entities = []Oauthscope{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthscopelisting) MarshalJSON() ([]byte, error) {
    type Alias Oauthscopelisting

    if OauthscopelistingMarshalled {
        return []byte("{}"), nil
    }
    OauthscopelistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Oauthscope `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Oauthscope{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

