package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterdetailedcategorymoduletemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterdetailedcategorymoduletemplateDud struct { 
    


    

}

// Supportcenterdetailedcategorymoduletemplate
type Supportcenterdetailedcategorymoduletemplate struct { 
    // Active - Whether this template is active or not
    Active bool `json:"active"`


    // Sidebar - Sidebar settings for the template
    Sidebar Supportcenterdetailedcategorymodulesidebar `json:"sidebar"`

}

// String returns a JSON representation of the model
func (o *Supportcenterdetailedcategorymoduletemplate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterdetailedcategorymoduletemplate) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterdetailedcategorymoduletemplate

    if SupportcenterdetailedcategorymoduletemplateMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterdetailedcategorymoduletemplateMarshalled = true

    return json.Marshal(&struct {
        
        Active bool `json:"active"`
        
        Sidebar Supportcenterdetailedcategorymodulesidebar `json:"sidebar"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

