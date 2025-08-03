package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanlistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanlistitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Capacityplanlistitem
type Capacityplanlistitem struct { 
    


    // Name
    Name string `json:"name"`


    // StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`


    // EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`


    // Metadata - The metadata of this capacity plan
    Metadata Capacityplanmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Capacityplanlistitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanlistitem) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanlistitem

    if CapacityplanlistitemMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanlistitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        StartBusinessUnitDate time.Time `json:"startBusinessUnitDate"`
        
        EndBusinessUnitDate time.Time `json:"endBusinessUnitDate"`
        
        Metadata Capacityplanmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

