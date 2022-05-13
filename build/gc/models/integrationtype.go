package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationtypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationtypeDud struct { 
    


    


    Description string `json:"description"`


    Provider string `json:"provider"`


    Category string `json:"category"`


    Images []Userimage `json:"images"`


    ConfigPropertiesSchemaUri string `json:"configPropertiesSchemaUri"`


    ConfigAdvancedSchemaUri string `json:"configAdvancedSchemaUri"`


    HelpUri string `json:"helpUri"`


    TermsOfServiceUri string `json:"termsOfServiceUri"`


    VendorName string `json:"vendorName"`


    VendorWebsiteUri string `json:"vendorWebsiteUri"`


    MarketplaceUri string `json:"marketplaceUri"`


    FaqUri string `json:"faqUri"`


    PrivacyPolicyUri string `json:"privacyPolicyUri"`


    SupportContactUri string `json:"supportContactUri"`


    SalesContactUri string `json:"salesContactUri"`


    HelpLinks []Helplink `json:"helpLinks"`


    Credentials map[string]Credentialspecification `json:"credentials"`


    NonInstallable bool `json:"nonInstallable"`


    MaxInstances int `json:"maxInstances"`


    UserPermissions []string `json:"userPermissions"`


    VendorOAuthClientIds []string `json:"vendorOAuthClientIds"`


    SelfUri string `json:"selfUri"`

}

// Integrationtype - Descriptor for a type of Integration.
type Integrationtype struct { 
    // Id - The ID of the integration type.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Integrationtype) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationtype) MarshalJSON() ([]byte, error) {
    type Alias Integrationtype

    if IntegrationtypeMarshalled {
        return []byte("{}"), nil
    }
    IntegrationtypeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

