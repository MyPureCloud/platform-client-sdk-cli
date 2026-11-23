package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestbulkevaluateresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestbulkevaluateresultDud struct { 
    


    

}

// Timeoffrequestbulkevaluateresult
type Timeoffrequestbulkevaluateresult struct { 
    // Entities
    Entities []Timeoffrequestbulkevaluatetemplateitem `json:"entities"`


    // ManagementUnit - The management unit to which the time-off requests belong
    ManagementUnit Managementunitreference `json:"managementUnit"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestbulkevaluateresult) String() string {
     o.Entities = []Timeoffrequestbulkevaluatetemplateitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestbulkevaluateresult) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestbulkevaluateresult

    if TimeoffrequestbulkevaluateresultMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestbulkevaluateresultMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffrequestbulkevaluatetemplateitem `json:"entities"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        *Alias
    }{

        
        Entities: []Timeoffrequestbulkevaluatetemplateitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

