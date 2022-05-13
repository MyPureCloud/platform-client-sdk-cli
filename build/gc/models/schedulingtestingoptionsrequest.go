package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingtestingoptionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingtestingoptionsrequestDud struct { 
    


    


    


    


    

}

// Schedulingtestingoptionsrequest
type Schedulingtestingoptionsrequest struct { 
    // FastScheduling - Whether to enable fast scheduling
    FastScheduling bool `json:"fastScheduling"`


    // DelayScheduling - Whether to force delayed scheduling
    DelayScheduling bool `json:"delayScheduling"`


    // FailScheduling - Whether to force scheduling to fail
    FailScheduling bool `json:"failScheduling"`


    // PopulateWarnings - Whether to populate warnings in the generated schedule
    PopulateWarnings bool `json:"populateWarnings"`


    // PopulateDeprecatedWarnings - Whether to populate deprecated warnings in the generated schedule
    PopulateDeprecatedWarnings bool `json:"populateDeprecatedWarnings"`

}

// String returns a JSON representation of the model
func (o *Schedulingtestingoptionsrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingtestingoptionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulingtestingoptionsrequest

    if SchedulingtestingoptionsrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulingtestingoptionsrequestMarshalled = true

    return json.Marshal(&struct {
        
        FastScheduling bool `json:"fastScheduling"`
        
        DelayScheduling bool `json:"delayScheduling"`
        
        FailScheduling bool `json:"failScheduling"`
        
        PopulateWarnings bool `json:"populateWarnings"`
        
        PopulateDeprecatedWarnings bool `json:"populateDeprecatedWarnings"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

