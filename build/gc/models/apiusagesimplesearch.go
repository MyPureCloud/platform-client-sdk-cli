package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApiusagesimplesearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApiusagesimplesearchDud struct { 
    


    


    


    


    

}

// Apiusagesimplesearch
type Apiusagesimplesearch struct { 
    // Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Metrics - Behaves like a SQL SELECT clause. Enables retrieving only named metrics. If omitted, all metrics that are available will be returned (like SELECT *).
    Metrics []string `json:"metrics"`


    // OauthClientNames - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of OAuth client names to be queried.
    OauthClientNames []string `json:"oauthClientNames"`


    // HttpMethods - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of HTTP methods to be queried.
    HttpMethods []string `json:"httpMethods"`


    // TemplateUris - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of Template Uris to be queried.
    TemplateUris []string `json:"templateUris"`

}

// String returns a JSON representation of the model
func (o *Apiusagesimplesearch) String() string {
    
     o.Metrics = []string{""} 
     o.OauthClientNames = []string{""} 
     o.HttpMethods = []string{""} 
     o.TemplateUris = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Apiusagesimplesearch) MarshalJSON() ([]byte, error) {
    type Alias Apiusagesimplesearch

    if ApiusagesimplesearchMarshalled {
        return []byte("{}"), nil
    }
    ApiusagesimplesearchMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Metrics []string `json:"metrics"`
        
        OauthClientNames []string `json:"oauthClientNames"`
        
        HttpMethods []string `json:"httpMethods"`
        
        TemplateUris []string `json:"templateUris"`
        *Alias
    }{

        


        
        Metrics: []string{""},
        


        
        OauthClientNames: []string{""},
        


        
        HttpMethods: []string{""},
        


        
        TemplateUris: []string{""},
        

        Alias: (*Alias)(u),
    })
}

