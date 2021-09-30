package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrusteeauditqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrusteeauditqueryrequestDud struct { 
    


    


    


    


    


    


    

}

// Trusteeauditqueryrequest
type Trusteeauditqueryrequest struct { 
    // TrusteeOrganizationIds - Limit returned audits to these trustee organizationIds.
    TrusteeOrganizationIds []string `json:"trusteeOrganizationIds"`


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
func (o *Trusteeauditqueryrequest) String() string {
    
    
     o.TrusteeOrganizationIds = []string{""} 
    
    
    
     o.TrusteeUserIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Facets = []Facet{{}} 
    
    
    
     o.Filters = []Filter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trusteeauditqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Trusteeauditqueryrequest

    if TrusteeauditqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    TrusteeauditqueryrequestMarshalled = true

    return json.Marshal(&struct { 
        TrusteeOrganizationIds []string `json:"trusteeOrganizationIds"`
        
        TrusteeUserIds []string `json:"trusteeUserIds"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        QueryPhrase string `json:"queryPhrase"`
        
        Facets []Facet `json:"facets"`
        
        Filters []Filter `json:"filters"`
        
        *Alias
    }{
        

        
        TrusteeOrganizationIds: []string{""},
        

        

        
        TrusteeUserIds: []string{""},
        

        

        

        

        

        

        

        

        
        Facets: []Facet{{}},
        

        

        
        Filters: []Filter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

