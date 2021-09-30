package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApiusagerowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApiusagerowDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Apiusagerow
type Apiusagerow struct { 
    // ClientId - Client Id associated with this query result
    ClientId string `json:"clientId"`


    // ClientName - Client Name associated with this query result
    ClientName string `json:"clientName"`


    // OrganizationId - Organization Id associated with this query result
    OrganizationId string `json:"organizationId"`


    // UserId - User Id associated with this query result
    UserId string `json:"userId"`


    // TemplateUri - Template Uri associated with this query result
    TemplateUri string `json:"templateUri"`


    // HttpMethod - HTTP Method associated with this query result
    HttpMethod string `json:"httpMethod"`


    // Status200 - Number of requests resulting in a 2xx HTTP status code
    Status200 int `json:"status200"`


    // Status300 - Number of requests resulting in a 3xx HTTP status code
    Status300 int `json:"status300"`


    // Status400 - Number of requests resulting in a 4xx HTTP status code
    Status400 int `json:"status400"`


    // Status500 - Number of requests resulting in a 5xx HTTP status code
    Status500 int `json:"status500"`


    // Status429 - Number of requests resulting in a 429 HTTP status code, this is a subset of the count returned with status400
    Status429 int `json:"status429"`


    // Requests - Total number of requests
    Requests int `json:"requests"`


    // Date - Date of requests, based on granularity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Date time.Time `json:"date"`

}

// String returns a JSON representation of the model
func (o *Apiusagerow) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Apiusagerow) MarshalJSON() ([]byte, error) {
    type Alias Apiusagerow

    if ApiusagerowMarshalled {
        return []byte("{}"), nil
    }
    ApiusagerowMarshalled = true

    return json.Marshal(&struct { 
        ClientId string `json:"clientId"`
        
        ClientName string `json:"clientName"`
        
        OrganizationId string `json:"organizationId"`
        
        UserId string `json:"userId"`
        
        TemplateUri string `json:"templateUri"`
        
        HttpMethod string `json:"httpMethod"`
        
        Status200 int `json:"status200"`
        
        Status300 int `json:"status300"`
        
        Status400 int `json:"status400"`
        
        Status500 int `json:"status500"`
        
        Status429 int `json:"status429"`
        
        Requests int `json:"requests"`
        
        Date time.Time `json:"date"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

