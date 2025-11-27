package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemquerypostrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemquerypostrequestDud struct { 
    


    


    


    


    


    


    

}

// Workitemquerypostrequest
type Workitemquerypostrequest struct { 
    // PageSize - Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an `after` key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200.
    PageSize int `json:"pageSize"`


    // VarSelect - Specify the value 'Count' for this parameter in order to return only the record count.
    VarSelect string `json:"select"`


    // Filters - List of filter objects to be used in the search. Valid filter names are: 'id', 'name', 'description', 'languageId', 'priority', 'dateCreated', 'dateModified', 'dateDue', 'dateExpires', 'durationInSeconds', 'ttl', 'statusId', 'statusCategory', 'dateClosed', 'externalContactId', 'reporterId', 'queueId', 'externalTag', 'modifiedBy', 'assignmentState', 'divisionId', 'customFields.<custom field name>'
    Filters []Workitemfilter `json:"filters"`


    // Attributes - List of entity attributes to be retrieved in the result.
    Attributes []string `json:"attributes"`


    // After - The cursor that points to the end of the set of entities that has been returned.
    After string `json:"after"`


    // Sort - Sort
    Sort Workitemquerysort `json:"sort"`


    // Expands - List of entity attributes to be expanded in the result.
    Expands []string `json:"expands"`

}

// String returns a JSON representation of the model
func (o *Workitemquerypostrequest) String() string {
    
    
     o.Filters = []Workitemfilter{{}} 
     o.Attributes = []string{""} 
    
    
     o.Expands = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemquerypostrequest) MarshalJSON() ([]byte, error) {
    type Alias Workitemquerypostrequest

    if WorkitemquerypostrequestMarshalled {
        return []byte("{}"), nil
    }
    WorkitemquerypostrequestMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        VarSelect string `json:"select"`
        
        Filters []Workitemfilter `json:"filters"`
        
        Attributes []string `json:"attributes"`
        
        After string `json:"after"`
        
        Sort Workitemquerysort `json:"sort"`
        
        Expands []string `json:"expands"`
        *Alias
    }{

        


        


        
        Filters: []Workitemfilter{{}},
        


        
        Attributes: []string{""},
        


        


        


        
        Expands: []string{""},
        

        Alias: (*Alias)(u),
    })
}

