package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminbulkupdatealternativeshifttradestaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminbulkupdatealternativeshifttradestaterequestDud struct { 
    


    

}

// Adminbulkupdatealternativeshifttradestaterequest
type Adminbulkupdatealternativeshifttradestaterequest struct { 
    // Entities
    Entities []Adminbulkupdatealternativeshifttradestate `json:"entities"`


    // ManagementUnitId - The ID of the management unit for this alternative shift bulk trade update
    ManagementUnitId string `json:"managementUnitId"`

}

// String returns a JSON representation of the model
func (o *Adminbulkupdatealternativeshifttradestaterequest) String() string {
     o.Entities = []Adminbulkupdatealternativeshifttradestate{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminbulkupdatealternativeshifttradestaterequest) MarshalJSON() ([]byte, error) {
    type Alias Adminbulkupdatealternativeshifttradestaterequest

    if AdminbulkupdatealternativeshifttradestaterequestMarshalled {
        return []byte("{}"), nil
    }
    AdminbulkupdatealternativeshifttradestaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adminbulkupdatealternativeshifttradestate `json:"entities"`
        
        ManagementUnitId string `json:"managementUnitId"`
        *Alias
    }{

        
        Entities: []Adminbulkupdatealternativeshifttradestate{{}},
        


        

        Alias: (*Alias)(u),
    })
}

