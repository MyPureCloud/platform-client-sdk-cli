package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontablerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontablerequestDud struct { 
    


    


    


    


    

}

// Createdecisiontablerequest
type Createdecisiontablerequest struct { 
    // Name - The decision table name.
    Name string `json:"name"`


    // Description - The decision table description.
    Description string `json:"description"`


    // DivisionId - The ID of the division the decision table belongs to.
    DivisionId string `json:"divisionId"`


    // SchemaId - The ID of the rules schema used by the decision table.
    SchemaId string `json:"schemaId"`


    // Columns - The column definitions for this decision table.
    Columns Createdecisiontablecolumnsrequest `json:"columns"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontablerequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontablerequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontablerequest

    if CreatedecisiontablerequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontablerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        
        SchemaId string `json:"schemaId"`
        
        Columns Createdecisiontablecolumnsrequest `json:"columns"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

