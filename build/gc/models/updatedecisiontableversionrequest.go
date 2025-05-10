package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedecisiontableversionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedecisiontableversionrequestDud struct { 
    

}

// Updatedecisiontableversionrequest
type Updatedecisiontableversionrequest struct { 
    // RowIndexUpdate - An update to a decision table version row index, which moves the row to a new position in the table. Execution output is returned based on the first matching row.
    RowIndexUpdate Updaterowindexrequest `json:"rowIndexUpdate"`

}

// String returns a JSON representation of the model
func (o *Updatedecisiontableversionrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedecisiontableversionrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatedecisiontableversionrequest

    if UpdatedecisiontableversionrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatedecisiontableversionrequestMarshalled = true

    return json.Marshal(&struct {
        
        RowIndexUpdate Updaterowindexrequest `json:"rowIndexUpdate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

