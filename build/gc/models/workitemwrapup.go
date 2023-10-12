package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemwrapupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemwrapupDud struct { 
    


    


    


    


    

}

// Workitemwrapup
type Workitemwrapup struct { 
    // Workitem - Workitem that the wrapup code has been added to.
    Workitem Workitemreference `json:"workitem"`


    // WrapupCode - The wrapup code used in the workitem.
    WrapupCode Wrapupidreference `json:"wrapupCode"`


    // ModifiedBy - The user who added the wrapup code to the workitem.
    ModifiedBy Userreference `json:"modifiedBy"`


    // User - The user for whom wrapup code was added. This may be the same as modifiedBy.
    User Userreference `json:"user"`


    // DateModified - The modified date of the Workitem when the wrapup code was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Workitemwrapup) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemwrapup) MarshalJSON() ([]byte, error) {
    type Alias Workitemwrapup

    if WorkitemwrapupMarshalled {
        return []byte("{}"), nil
    }
    WorkitemwrapupMarshalled = true

    return json.Marshal(&struct {
        
        Workitem Workitemreference `json:"workitem"`
        
        WrapupCode Wrapupidreference `json:"wrapupCode"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        User Userreference `json:"user"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

