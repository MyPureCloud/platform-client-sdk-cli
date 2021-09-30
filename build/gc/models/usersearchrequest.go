package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersearchrequestDud struct { 
    


    


    


    


    


    


    


    


    

}

// Usersearchrequest
type Usersearchrequest struct { 
    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // PageSize - The number of results per page
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // Sort - Multi-value sort order, list of multiple sort values
    Sort []Searchsort `json:"sort"`


    // Expand - Provides more details about a specified resource
    Expand []string `json:"expand"`


    // Query
    Query []Usersearchcriteria `json:"query"`


    // IntegrationPresenceSource - Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an \"expand\". When using this parameter the maximum number of users that can be returned is 100.
    IntegrationPresenceSource string `json:"integrationPresenceSource"`


    // EnforcePermissions - This property only applies to api/v2/user/search; when set to true add additional search criteria to filter users by: directory:user:view
    EnforcePermissions bool `json:"enforcePermissions"`

}

// String returns a JSON representation of the model
func (o *Usersearchrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Sort = []Searchsort{{}} 
    
    
    
     o.Expand = []string{""} 
    
    
    
     o.Query = []Usersearchcriteria{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Usersearchrequest

    if UsersearchrequestMarshalled {
        return []byte("{}"), nil
    }
    UsersearchrequestMarshalled = true

    return json.Marshal(&struct { 
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Expand []string `json:"expand"`
        
        Query []Usersearchcriteria `json:"query"`
        
        IntegrationPresenceSource string `json:"integrationPresenceSource"`
        
        EnforcePermissions bool `json:"enforcePermissions"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Sort: []Searchsort{{}},
        

        

        
        Expand: []string{""},
        

        

        
        Query: []Usersearchcriteria{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

