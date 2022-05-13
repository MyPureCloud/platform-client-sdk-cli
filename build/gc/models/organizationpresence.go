package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpresenceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Organizationpresence
type Organizationpresence struct { 
    


    // Name
    Name string `json:"name"`


    // LanguageLabels - The label used for the system presence in each specified language
    LanguageLabels map[string]string `json:"languageLabels"`


    // SystemPresence
    SystemPresence string `json:"systemPresence"`


    // Deactivated
    Deactivated bool `json:"deactivated"`


    // Primary
    Primary bool `json:"primary"`


    // CreatedBy
    CreatedBy User `json:"createdBy"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedBy
    ModifiedBy User `json:"modifiedBy"`


    // ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    

}

// String returns a JSON representation of the model
func (o *Organizationpresence) String() string {
    
     o.LanguageLabels = map[string]string{"": ""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpresence) MarshalJSON() ([]byte, error) {
    type Alias Organizationpresence

    if OrganizationpresenceMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpresenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        LanguageLabels map[string]string `json:"languageLabels"`
        
        SystemPresence string `json:"systemPresence"`
        
        Deactivated bool `json:"deactivated"`
        
        Primary bool `json:"primary"`
        
        CreatedBy User `json:"createdBy"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedBy User `json:"modifiedBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        
        LanguageLabels: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

