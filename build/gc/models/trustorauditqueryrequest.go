package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustorauditqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustorauditqueryrequestDud struct { 
    


    


    


    


    


    


    

}

// Trustorauditqueryrequest
type Trustorauditqueryrequest struct { 
    // TrustorOrganizationId - Limit returned audits to this trustor organizationId.
    TrustorOrganizationId string `json:"trustorOrganizationId"`


    // TrusteeUserIds - Limit returned audits to these trustee userIds.
    TrusteeUserIds []string `json:"trusteeUserIds"`


    // StartDate - Starting date/time for the audit search. ISO-8601 formatted date-time, UTC.
    StartDate time.Time `json:"startDate"`


    // EndDate - Ending date/time for the audit search. ISO-8601 formatted date-time, UTC.
    EndDate time.Time `json:"endDate"`


    // QueryPhrase - Word or phrase to look for in audit bodies.
    QueryPhrase string `json:"queryPhrase"`


    // Facets - Facet information to be returned with the query results.
    Facets []Facet `json:"facets"`


    // Filters - Additional custom filters to be applied to the query.
    Filters []Filter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *Trustorauditqueryrequest) String() string {
    
    
    
    
    
    
     o.TrusteeUserIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Facets = []Facet{{}} 
    
    
    
     o.Filters = []Filter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustorauditqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Trustorauditqueryrequest

    if TrustorauditqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    TrustorauditqueryrequestMarshalled = true

    return json.Marshal(&struct { 
        TrustorOrganizationId string `json:"trustorOrganizationId"`
        
        TrusteeUserIds []string `json:"trusteeUserIds"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        QueryPhrase string `json:"queryPhrase"`
        
        Facets []Facet `json:"facets"`
        
        Filters []Filter `json:"filters"`
        
        *Alias
    }{
        

        

        

        
        TrusteeUserIds: []string{""},
        

        

        

        

        

        

        

        

        
        Facets: []Facet{{}},
        

        

        
        Filters: []Filter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

