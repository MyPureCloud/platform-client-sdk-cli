package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2userMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2userDud struct { 
    Id string `json:"id"`


    Schemas []string `json:"schemas"`


    


    


    


    


    


    


    


    


    Groups []Scimv2groupreference `json:"groups"`


    


    


    


    Meta Scimmetadata `json:"meta"`

}

// Scimv2user - Defines a SCIM user.
type Scimv2user struct { 
    


    


    // Active - Indicates whether the user's administrative status is active.
    Active bool `json:"active"`


    // UserName - The user's Genesys Cloud email address. Must be unique.
    UserName string `json:"userName"`


    // DisplayName - The display name of the user.
    DisplayName string `json:"displayName"`


    // Password - The new password for the Genesys Cloud user. Does not return an existing password. When creating a user, if a password is not supplied, then a password will be randomly generated that is 40 characters in length and contains five characters from each of the password policy groups.
    Password string `json:"password"`


    // Title - The user's title.
    Title string `json:"title"`


    // PhoneNumbers - The list of the user's phone numbers.
    PhoneNumbers []Scimphonenumber `json:"phoneNumbers"`


    // Emails - The list of the user's email addresses.
    Emails []Scimemail `json:"emails"`


    // ExternalId - The external ID of the user. Set by the provisioning client. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readWrite\".
    ExternalId string `json:"externalId"`


    


    // Roles - The list of roles assigned to the user.
    Roles []Scimuserrole `json:"roles"`


    // UrnIetfParamsScimSchemasExtensionEnterprise20User - The URI of the schema for the enterprise user.
    UrnIetfParamsScimSchemasExtensionEnterprise20User Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`


    // UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User - The URI of the schema for the Genesys Cloud user.
    UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User"`


    

}

// String returns a JSON representation of the model
func (o *Scimv2user) String() string {
    
    
    
    
    
     o.PhoneNumbers = []Scimphonenumber{{}} 
     o.Emails = []Scimemail{{}} 
    
     o.Roles = []Scimuserrole{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2user) MarshalJSON() ([]byte, error) {
    type Alias Scimv2user

    if Scimv2userMarshalled {
        return []byte("{}"), nil
    }
    Scimv2userMarshalled = true

    return json.Marshal(&struct {
        
        Active bool `json:"active"`
        
        UserName string `json:"userName"`
        
        DisplayName string `json:"displayName"`
        
        Password string `json:"password"`
        
        Title string `json:"title"`
        
        PhoneNumbers []Scimphonenumber `json:"phoneNumbers"`
        
        Emails []Scimemail `json:"emails"`
        
        ExternalId string `json:"externalId"`
        
        Roles []Scimuserrole `json:"roles"`
        
        UrnIetfParamsScimSchemasExtensionEnterprise20User Scimv2enterpriseuser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
        
        UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User Scimuserextensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        PhoneNumbers: []Scimphonenumber{{}},
        


        
        Emails: []Scimemail{{}},
        


        


        


        
        Roles: []Scimuserrole{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

