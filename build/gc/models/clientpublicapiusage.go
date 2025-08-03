package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientpublicapiusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientpublicapiusageDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Clientpublicapiusage
type Clientpublicapiusage struct { 
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


    // OrganizationId - The organizationId that made the request.
    OrganizationId string `json:"organizationId"`

}

// String returns a JSON representation of the model
func (o *Clientpublicapiusage) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientpublicapiusage) MarshalJSON() ([]byte, error) {
    type Alias Clientpublicapiusage

    if ClientpublicapiusageMarshalled {
        return []byte("{}"), nil
    }
    ClientpublicapiusageMarshalled = true

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
        
        OrganizationId string `json:"organizationId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

