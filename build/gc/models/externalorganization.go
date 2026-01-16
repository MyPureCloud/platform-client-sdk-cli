package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalorganizationDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    ExternalDataSources []Externaldatasource `json:"externalDataSources"`


    SelfUri string `json:"selfUri"`

}

// Externalorganization
type Externalorganization struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name - The name of the company. Max: 1000 characters. Leading and trailing whitespace stripped.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // CompanyType
    CompanyType string `json:"companyType"`


    // Industry
    Industry string `json:"industry"`


    // Address
    Address Contactaddress `json:"address"`


    // PhoneNumber
    PhoneNumber Phonenumber `json:"phoneNumber"`


    // FaxNumber
    FaxNumber Phonenumber `json:"faxNumber"`


    // EmployeeCount
    EmployeeCount int `json:"employeeCount"`


    // Revenue
    Revenue int `json:"revenue"`


    // Tags
    Tags []string `json:"tags"`


    // Websites
    Websites []string `json:"websites"`


    // Tickers
    Tickers []Ticker `json:"tickers"`


    // TwitterId
    TwitterId Twitterid `json:"twitterId"`


    // ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the organization. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
    ExternalSystemUrl string `json:"externalSystemUrl"`


    // ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifyDate time.Time `json:"modifyDate"`


    // CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreateDate time.Time `json:"createDate"`


    // Trustor
    Trustor Trustor `json:"trustor"`


    // Schema - The schema defining custom fields for this contact
    Schema Dataschema `json:"schema"`


    // CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
    CustomFields map[string]interface{} `json:"customFields"`


    // Identifiers - Identifiers claimed by this external org
    Identifiers []Externalorganizationidentifier `json:"identifiers"`


    // ExternalIds - A list of external identifiers that identify this External Organization in an external system. Max 10 items.
    ExternalIds []Externalid `json:"externalIds"`


    


    

}

// String returns a JSON representation of the model
func (o *Externalorganization) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Tags = []string{""} 
     o.Websites = []string{""} 
     o.Tickers = []Ticker{{}} 
    
    
    
    
    
    
     o.CustomFields = map[string]interface{}{"": Interface{}} 
     o.Identifiers = []Externalorganizationidentifier{{}} 
     o.ExternalIds = []Externalid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalorganization) MarshalJSON() ([]byte, error) {
    type Alias Externalorganization

    if ExternalorganizationMarshalled {
        return []byte("{}"), nil
    }
    ExternalorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writablestarrabledivision `json:"division"`
        
        CompanyType string `json:"companyType"`
        
        Industry string `json:"industry"`
        
        Address Contactaddress `json:"address"`
        
        PhoneNumber Phonenumber `json:"phoneNumber"`
        
        FaxNumber Phonenumber `json:"faxNumber"`
        
        EmployeeCount int `json:"employeeCount"`
        
        Revenue int `json:"revenue"`
        
        Tags []string `json:"tags"`
        
        Websites []string `json:"websites"`
        
        Tickers []Ticker `json:"tickers"`
        
        TwitterId Twitterid `json:"twitterId"`
        
        ExternalSystemUrl string `json:"externalSystemUrl"`
        
        ModifyDate time.Time `json:"modifyDate"`
        
        CreateDate time.Time `json:"createDate"`
        
        Trustor Trustor `json:"trustor"`
        
        Schema Dataschema `json:"schema"`
        
        CustomFields map[string]interface{} `json:"customFields"`
        
        Identifiers []Externalorganizationidentifier `json:"identifiers"`
        
        ExternalIds []Externalid `json:"externalIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        Tags: []string{""},
        


        
        Websites: []string{""},
        


        
        Tickers: []Ticker{{}},
        


        


        


        


        


        


        


        
        CustomFields: map[string]interface{}{"": Interface{}},
        


        
        Identifiers: []Externalorganizationidentifier{{}},
        


        
        ExternalIds: []Externalid{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

