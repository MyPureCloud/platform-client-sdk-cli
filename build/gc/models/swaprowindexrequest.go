package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SwaprowindexrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SwaprowindexrequestDud struct { 
    


    


    


    

}

// Swaprowindexrequest
type Swaprowindexrequest struct { 
    // SourceRowId - Unique identifier of the source row to swap
    SourceRowId string `json:"sourceRowId"`


    // SourceRowIndex - The current index position of the source row. Must be positive, starting from 1 and less than or equal to the size of the table
    SourceRowIndex int `json:"sourceRowIndex"`


    // TargetRowId - Unique identifier of the target row to swap
    TargetRowId string `json:"targetRowId"`


    // TargetRowIndex - The current index position of the target row. Must be positive, starting from 1 and less than or equal to the size of the table
    TargetRowIndex int `json:"targetRowIndex"`

}

// String returns a JSON representation of the model
func (o *Swaprowindexrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Swaprowindexrequest) MarshalJSON() ([]byte, error) {
    type Alias Swaprowindexrequest

    if SwaprowindexrequestMarshalled {
        return []byte("{}"), nil
    }
    SwaprowindexrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceRowId string `json:"sourceRowId"`
        
        SourceRowIndex int `json:"sourceRowIndex"`
        
        TargetRowId string `json:"targetRowId"`
        
        TargetRowIndex int `json:"targetRowIndex"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

