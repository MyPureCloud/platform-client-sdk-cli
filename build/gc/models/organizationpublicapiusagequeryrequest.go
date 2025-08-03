package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationpublicapiusagequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationpublicapiusagequeryrequestDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Organizationpublicapiusagequeryrequest
type Organizationpublicapiusagequeryrequest struct { 
    // Interval - Specify the interval to query on. Start and end are inclusive. Start date cannot be more than a year ago. End date cannot be more than 90 days after the start. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Granularity - Specify the granularity to aggregate the data to.
    Granularity string `json:"granularity"`


    // SortBy - Specify how to sort the returned data.
    SortBy []Usagequerysortby `json:"sortBy"`


    // Metrics - Specify which metrics you want returned (all will be returned by default).
    Metrics []string `json:"metrics"`


    // TemplateUris - Specify if you only want a subset of templateUris represented in the query.
    TemplateUris []string `json:"templateUris"`


    // HttpMethods - Specify if you only want a subset of httpMethods represented in the query.
    HttpMethods []string `json:"httpMethods"`


    // Platforms - Specify if you only want a subset of platforms represented in the query.
    Platforms []string `json:"platforms"`


    // GroupBy - Specify how to aggregate the data (by default the data is not aggregated).
    GroupBy []string `json:"groupBy"`


    // UserIds - Specify if you only want a subset of userIds represented in the query.
    UserIds []string `json:"userIds"`


    // OauthClientIds - Specify if you only want a subset of oauthClientIds represented in the query.
    OauthClientIds []string `json:"oauthClientIds"`

}

// String returns a JSON representation of the model
func (o *Organizationpublicapiusagequeryrequest) String() string {
    
    
     o.SortBy = []Usagequerysortby{{}} 
     o.Metrics = []string{""} 
     o.TemplateUris = []string{""} 
     o.HttpMethods = []string{""} 
     o.Platforms = []string{""} 
     o.GroupBy = []string{""} 
     o.UserIds = []string{""} 
     o.OauthClientIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationpublicapiusagequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Organizationpublicapiusagequeryrequest

    if OrganizationpublicapiusagequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    OrganizationpublicapiusagequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Granularity string `json:"granularity"`
        
        SortBy []Usagequerysortby `json:"sortBy"`
        
        Metrics []string `json:"metrics"`
        
        TemplateUris []string `json:"templateUris"`
        
        HttpMethods []string `json:"httpMethods"`
        
        Platforms []string `json:"platforms"`
        
        GroupBy []string `json:"groupBy"`
        
        UserIds []string `json:"userIds"`
        
        OauthClientIds []string `json:"oauthClientIds"`
        *Alias
    }{

        


        


        
        SortBy: []Usagequerysortby{{}},
        


        
        Metrics: []string{""},
        


        
        TemplateUris: []string{""},
        


        
        HttpMethods: []string{""},
        


        
        Platforms: []string{""},
        


        
        GroupBy: []string{""},
        


        
        UserIds: []string{""},
        


        
        OauthClientIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

