package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutostatustransitiondetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutostatustransitiondetailDud struct { 
    


    


    

}

// Autostatustransitiondetail
type Autostatustransitiondetail struct { 
    // NextStatus - Next status of auto status transition.
    NextStatus Workitemstatusreference `json:"nextStatus"`


    // DateOfTransition - Date at which auto status transition occurs. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateOfTransition time.Time `json:"dateOfTransition"`


    // ErrorDetails - This property will be set if auto status transition is failed.
    ErrorDetails Taskmanagementerrordetails `json:"errorDetails"`

}

// String returns a JSON representation of the model
func (o *Autostatustransitiondetail) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Autostatustransitiondetail) MarshalJSON() ([]byte, error) {
    type Alias Autostatustransitiondetail

    if AutostatustransitiondetailMarshalled {
        return []byte("{}"), nil
    }
    AutostatustransitiondetailMarshalled = true

    return json.Marshal(&struct {
        
        NextStatus Workitemstatusreference `json:"nextStatus"`
        
        DateOfTransition time.Time `json:"dateOfTransition"`
        
        ErrorDetails Taskmanagementerrordetails `json:"errorDetails"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

