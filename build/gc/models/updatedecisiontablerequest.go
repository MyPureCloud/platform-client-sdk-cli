package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedecisiontablerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedecisiontablerequestDud struct { 
    


    


    

}

// Updatedecisiontablerequest
type Updatedecisiontablerequest struct { 
    // Name - The decision table name.
    Name string `json:"name"`


    // Description - The decision table description.
    Description string `json:"description"`


    // Columns - The column definitions for this decision table.
    Columns Updatedecisiontablecolumnsrequest `json:"columns"`

}

// String returns a JSON representation of the model
func (o *Updatedecisiontablerequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedecisiontablerequest) MarshalJSON() ([]byte, error) {
    type Alias Updatedecisiontablerequest

    if UpdatedecisiontablerequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatedecisiontablerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Columns Updatedecisiontablecolumnsrequest `json:"columns"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

