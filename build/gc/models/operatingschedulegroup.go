package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperatingschedulegroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperatingschedulegroupDud struct { 
    


    


    

}

// Operatingschedulegroup
type Operatingschedulegroup struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Operatingschedulegroup) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operatingschedulegroup) MarshalJSON() ([]byte, error) {
    type Alias Operatingschedulegroup

    if OperatingschedulegroupMarshalled {
        return []byte("{}"), nil
    }
    OperatingschedulegroupMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

