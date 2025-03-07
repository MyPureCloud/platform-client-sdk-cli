package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentschedulehistorychangemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentschedulehistorychangemetadataDud struct { 
    


    

}

// Buagentschedulehistorychangemetadata
type Buagentschedulehistorychangemetadata struct { 
    // DateModified - The timestamp of the schedule change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The user that made the schedule change. The id may be 'System' if it was an automated process
    ModifiedBy Userreference `json:"modifiedBy"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychangemetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentschedulehistorychangemetadata) MarshalJSON() ([]byte, error) {
    type Alias Buagentschedulehistorychangemetadata

    if BuagentschedulehistorychangemetadataMarshalled {
        return []byte("{}"), nil
    }
    BuagentschedulehistorychangemetadataMarshalled = true

    return json.Marshal(&struct {
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

