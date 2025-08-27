package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalcontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalcontactDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    ExternalDataSources []Externaldatasource `json:"externalDataSources"`


    VarType string `json:"type"`


    CanonicalContact Contactaddressableentityref `json:"canonicalContact"`


    MergeSet []Contactaddressableentityref `json:"mergeSet"`


    MergedFrom []Contactaddressableentityref `json:"mergedFrom"`


    MergedTo Contactaddressableentityref `json:"mergedTo"`


    MergeOperation Mergeoperation `json:"mergeOperation"`


    SelfUri string `json:"selfUri"`

}

// Externalcontact
type Externalcontact struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // FirstName - The first name of the contact.
    FirstName string `json:"firstName"`


    // MiddleName
    MiddleName string `json:"middleName"`


    // LastName - The last name of the contact.
    LastName string `json:"lastName"`


    // Salutation
    Salutation string `json:"salutation"`


    // Title
    Title string `json:"title"`


    // WorkPhone
    WorkPhone Phonenumber `json:"workPhone"`


    // CellPhone
    CellPhone Phonenumber `json:"cellPhone"`


    // HomePhone
    HomePhone Phonenumber `json:"homePhone"`


    // OtherPhone
    OtherPhone Phonenumber `json:"otherPhone"`


    // WorkEmail
    WorkEmail string `json:"workEmail"`


    // PersonalEmail
    PersonalEmail string `json:"personalEmail"`


    // OtherEmail
    OtherEmail string `json:"otherEmail"`


    // Address
    Address Contactaddress `json:"address"`


    // TwitterId
    TwitterId Twitterid `json:"twitterId"`


    // LineId
    LineId Lineid `json:"lineId"`


    // WhatsAppId
    WhatsAppId Whatsappid `json:"whatsAppId"`


    // FacebookId
    FacebookId Facebookid `json:"facebookId"`


    // InstagramId - User information for an Instagram account
    InstagramId Instagramid `json:"instagramId"`


    // AppleOpaqueIds - User information for an Apple account
    AppleOpaqueIds []Appleopaqueid `json:"appleOpaqueIds"`


    // ExternalIds - A list of external identifiers that identify this contact in an external system
    ExternalIds []Externalid `json:"externalIds"`


    // Identifiers - Identifiers claimed by this contact
    Identifiers []Contactidentifier `json:"identifiers"`


    // ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifyDate time.Time `json:"modifyDate"`


    // CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreateDate time.Time `json:"createDate"`


    // ExternalOrganization
    ExternalOrganization Externalorganization `json:"externalOrganization"`


    // SurveyOptOut
    SurveyOptOut bool `json:"surveyOptOut"`


    // ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the contact. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
    ExternalSystemUrl string `json:"externalSystemUrl"`


    // Schema - The schema defining custom fields for this contact
    Schema Dataschema `json:"schema"`


    // CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
    CustomFields map[string]interface{} `json:"customFields"`


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Externalcontact) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AppleOpaqueIds = []Appleopaqueid{{}} 
     o.ExternalIds = []Externalid{{}} 
     o.Identifiers = []Contactidentifier{{}} 
    
    
    
    
    
    
     o.CustomFields = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalcontact) MarshalJSON() ([]byte, error) {
    type Alias Externalcontact

    if ExternalcontactMarshalled {
        return []byte("{}"), nil
    }
    ExternalcontactMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Division Writablestarrabledivision `json:"division"`
        
        FirstName string `json:"firstName"`
        
        MiddleName string `json:"middleName"`
        
        LastName string `json:"lastName"`
        
        Salutation string `json:"salutation"`
        
        Title string `json:"title"`
        
        WorkPhone Phonenumber `json:"workPhone"`
        
        CellPhone Phonenumber `json:"cellPhone"`
        
        HomePhone Phonenumber `json:"homePhone"`
        
        OtherPhone Phonenumber `json:"otherPhone"`
        
        WorkEmail string `json:"workEmail"`
        
        PersonalEmail string `json:"personalEmail"`
        
        OtherEmail string `json:"otherEmail"`
        
        Address Contactaddress `json:"address"`
        
        TwitterId Twitterid `json:"twitterId"`
        
        LineId Lineid `json:"lineId"`
        
        WhatsAppId Whatsappid `json:"whatsAppId"`
        
        FacebookId Facebookid `json:"facebookId"`
        
        InstagramId Instagramid `json:"instagramId"`
        
        AppleOpaqueIds []Appleopaqueid `json:"appleOpaqueIds"`
        
        ExternalIds []Externalid `json:"externalIds"`
        
        Identifiers []Contactidentifier `json:"identifiers"`
        
        ModifyDate time.Time `json:"modifyDate"`
        
        CreateDate time.Time `json:"createDate"`
        
        ExternalOrganization Externalorganization `json:"externalOrganization"`
        
        SurveyOptOut bool `json:"surveyOptOut"`
        
        ExternalSystemUrl string `json:"externalSystemUrl"`
        
        Schema Dataschema `json:"schema"`
        
        CustomFields map[string]interface{} `json:"customFields"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        AppleOpaqueIds: []Appleopaqueid{{}},
        


        
        ExternalIds: []Externalid{{}},
        


        
        Identifiers: []Contactidentifier{{}},
        


        


        


        


        


        


        


        
        CustomFields: map[string]interface{}{"": Interface{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

