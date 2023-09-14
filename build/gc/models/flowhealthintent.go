package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthintentDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowhealthintent
type Flowhealthintent struct { 
    


    // Name
    Name string `json:"name"`


    // FlowVersionInfo - Info about given flow version.
    FlowVersionInfo Flowhealthintentversioninfo `json:"flowVersionInfo"`


    // Language - Language provided for this intent's health.
    Language string `json:"language"`


    // Health - Health computation details for given language.
    Health Healthinfo `json:"health"`


    

}

// String returns a JSON representation of the model
func (o *Flowhealthintent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthintent) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthintent

    if FlowhealthintentMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthintentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        FlowVersionInfo Flowhealthintentversioninfo `json:"flowVersionInfo"`
        
        Language string `json:"language"`
        
        Health Healthinfo `json:"health"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

