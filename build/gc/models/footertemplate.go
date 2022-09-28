package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FootertemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FootertemplateDud struct { 
    


    

}

// Footertemplate - The Footer template identifies the Footer type and its footerUsage
type Footertemplate struct { 
    // VarType - Specifies the type represented by Footer.
    VarType string `json:"type"`


    // ApplicableResources - Specifies the canned response template where the footer can be used.
    ApplicableResources []string `json:"applicableResources"`

}

// String returns a JSON representation of the model
func (o *Footertemplate) String() string {
    
     o.ApplicableResources = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Footertemplate) MarshalJSON() ([]byte, error) {
    type Alias Footertemplate

    if FootertemplateMarshalled {
        return []byte("{}"), nil
    }
    FootertemplateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ApplicableResources []string `json:"applicableResources"`
        *Alias
    }{

        


        
        ApplicableResources: []string{""},
        

        Alias: (*Alias)(u),
    })
}

