package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdaterowindexrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdaterowindexrequestDud struct { 
    


    

}

// Updaterowindexrequest
type Updaterowindexrequest struct { 
    // RowId - The row UUID.
    RowId string `json:"rowId"`


    // RowIndex - The updated row index. Must be an integer value greater than or equal to 1. Must be less than or equal to x, where x is the number of rows in the decision table version.
    RowIndex int `json:"rowIndex"`

}

// String returns a JSON representation of the model
func (o *Updaterowindexrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updaterowindexrequest) MarshalJSON() ([]byte, error) {
    type Alias Updaterowindexrequest

    if UpdaterowindexrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdaterowindexrequestMarshalled = true

    return json.Marshal(&struct {
        
        RowId string `json:"rowId"`
        
        RowIndex int `json:"rowIndex"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

