package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpublicapiusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpublicapiusageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Organizationpublicapiusage
type Organizationpublicapiusage struct { 
    // Date - The date of the usage. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // Platform - The platform the request(s) is/were made on.
    Platform string `json:"platform"`


    // HttpMethod - The http method of the request(s)
    HttpMethod string `json:"httpMethod"`


    // TemplateUri - The templateUri of the request(s).
    TemplateUri string `json:"templateUri"`


    // RequestCount - The total number of requests.
    RequestCount int `json:"requestCount"`


    // Status200 - The number of requests resulting in a 2xx HTTP status code.
    Status200 int `json:"status200"`


    // Status300 - The number of requests resulting in a 3xx HTTP status code.
    Status300 int `json:"status300"`


    // Status400 - The number of requests resulting in a 4xx HTTP status code.
    Status400 int `json:"status400"`


    // Status429 - The number of requests resulting in a 429 HTTP status code.
    Status429 int `json:"status429"`


    // Status500 - The number of requests resulting in a 5xx HTTP status code.
    Status500 int `json:"status500"`


    // OauthClient - The id of the oauthClient that made the request(s).
    OauthClient Domainentityref `json:"oauthClient"`


    // User - The id of the user who made the request(s).
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Organizationpublicapiusage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpublicapiusage) MarshalJSON() ([]byte, error) {
    type Alias Organizationpublicapiusage

    if OrganizationpublicapiusageMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpublicapiusageMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        Platform string `json:"platform"`
        
        HttpMethod string `json:"httpMethod"`
        
        TemplateUri string `json:"templateUri"`
        
        RequestCount int `json:"requestCount"`
        
        Status200 int `json:"status200"`
        
        Status300 int `json:"status300"`
        
        Status400 int `json:"status400"`
        
        Status429 int `json:"status429"`
        
        Status500 int `json:"status500"`
        
        OauthClient Domainentityref `json:"oauthClient"`
        
        User Userreference `json:"user"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

