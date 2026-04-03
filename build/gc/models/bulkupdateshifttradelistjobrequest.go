package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateshifttradelistjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateshifttradelistjobrequestDud struct { 
    


    


    

}

// Bulkupdateshifttradelistjobrequest
type Bulkupdateshifttradelistjobrequest struct { 
    // ManagementUnitIds - The IDs of the management units from which to update shift trades
    ManagementUnitIds []string `json:"managementUnitIds"`


    // WeekDates - The start week dates in which the shift trades being updated occur in the business unit time zone (yyyy-MM-dd format)
    WeekDates []time.Time `json:"weekDates"`


    // Entities - The shift trades that are being updated
    Entities []Bulkupdateshifttradestaterequestitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradelistjobrequest) String() string {
     o.ManagementUnitIds = []string{""} 
     o.WeekDates = []time.Time{{}} 
     o.Entities = []Bulkupdateshifttradestaterequestitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateshifttradelistjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateshifttradelistjobrequest

    if BulkupdateshifttradelistjobrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateshifttradelistjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        WeekDates []time.Time `json:"weekDates"`
        
        Entities []Bulkupdateshifttradestaterequestitem `json:"entities"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        WeekDates: []time.Time{{}},
        


        
        Entities: []Bulkupdateshifttradestaterequestitem{{}},
        

        Alias: (*Alias)(u),
    })
}

