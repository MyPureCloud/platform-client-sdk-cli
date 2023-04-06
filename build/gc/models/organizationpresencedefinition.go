package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpresencedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpresencedefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Organizationpresencedefinition
type Organizationpresencedefinition struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - The type of definition
    VarType string `json:"type"`


    // LanguageLabels - The label used for the definition in each specified language
    LanguageLabels map[string]string `json:"languageLabels"`


    // SystemPresence
    SystemPresence string `json:"systemPresence"`


    // DivisionId
    DivisionId string `json:"divisionId"`


    // Deactivated
    Deactivated bool `json:"deactivated"`


    

}

// String returns a JSON representation of the model
func (o *Organizationpresencedefinition) String() string {
    
    
     o.LanguageLabels = map[string]string{"": ""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpresencedefinition) MarshalJSON() ([]byte, error) {
    type Alias Organizationpresencedefinition

    if OrganizationpresencedefinitionMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpresencedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        LanguageLabels map[string]string `json:"languageLabels"`
        
        SystemPresence string `json:"systemPresence"`
        
        DivisionId string `json:"divisionId"`
        
        Deactivated bool `json:"deactivated"`
        *Alias
    }{

        


        


        


        
        LanguageLabels: map[string]string{"": ""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

