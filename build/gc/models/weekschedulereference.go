package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekschedulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekschedulereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    

}

// Weekschedulereference
type Weekschedulereference struct { 
    


    


    // WeekDate - First day of this week schedule in yyyy-MM-dd format
    WeekDate string `json:"weekDate"`

}

// String returns a JSON representation of the model
func (o *Weekschedulereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekschedulereference) MarshalJSON() ([]byte, error) {
    type Alias Weekschedulereference

    if WeekschedulereferenceMarshalled {
        return []byte("{}"), nil
    }
    WeekschedulereferenceMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate string `json:"weekDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

