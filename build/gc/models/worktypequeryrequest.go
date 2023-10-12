package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypequeryrequestDud struct { 
    


    


    


    


    


    

}

// Worktypequeryrequest
type Worktypequeryrequest struct { 
    // PageSize - Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an `after` key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200.
    PageSize int `json:"pageSize"`


    // VarSelect - Specify the value 'Count' for this parameter in order to return only the record count.
    VarSelect string `json:"select"`


    // Filters - List of filter objects to be used in the search.
    Filters []Workitemfilter `json:"filters"`


    // Attributes - List of entity attributes to be retrieved in the result.
    Attributes []string `json:"attributes"`


    // After - The cursor that points to the end of the set of entities that has been returned.
    After string `json:"after"`


    // Sort - Sort
    Sort Worktypequerysort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Worktypequeryrequest) String() string {
    
    
     o.Filters = []Workitemfilter{{}} 
     o.Attributes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Worktypequeryrequest

    if WorktypequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    WorktypequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        VarSelect string `json:"select"`
        
        Filters []Workitemfilter `json:"filters"`
        
        Attributes []string `json:"attributes"`
        
        After string `json:"after"`
        
        Sort Worktypequerysort `json:"sort"`
        *Alias
    }{

        


        


        
        Filters: []Workitemfilter{{}},
        


        
        Attributes: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

