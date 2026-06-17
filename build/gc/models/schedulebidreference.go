package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Schedulebidreference
type Schedulebidreference struct { 
    // Id - The ID of the schedule bid
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Schedulebidreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidreference) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidreference

    if SchedulebidreferenceMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

