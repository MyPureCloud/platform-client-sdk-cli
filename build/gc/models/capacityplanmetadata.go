package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanmetadataDud struct { 
    ModifiedBy Userreference `json:"modifiedBy"`


    ModifiedDate time.Time `json:"modifiedDate"`


    CreatedDate time.Time `json:"createdDate"`


    CreatedBy Userreference `json:"createdBy"`


    

}

// Capacityplanmetadata
type Capacityplanmetadata struct { 
    


    


    


    


    // Version - The version of the capacity plan
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Capacityplanmetadata) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanmetadata) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanmetadata

    if CapacityplanmetadataMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanmetadataMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

