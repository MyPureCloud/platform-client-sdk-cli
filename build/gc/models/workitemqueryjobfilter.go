package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobfilterDud struct { 
    


    


    

}

// Workitemqueryjobfilter
type Workitemqueryjobfilter struct { 
    // Name - Attribute name. Valid filter names are: 'workbinId', 'id', 'typeId', 'priority', 'dateCreated', 'dateDue', 'statusId', 'dateClosed', 'externalContactId', 'assigneeId', 'assignmentState', 'queueId', 'externalTag', 'divisionId'
    Name string `json:"name"`


    // Operator - Filter operator.
    Operator string `json:"operator"`


    // Values - List of values to be used in the filter.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobfilter) String() string {
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobfilter) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobfilter

    if WorkitemqueryjobfilterMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobfilterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

