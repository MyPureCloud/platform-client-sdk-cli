package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchshifttradesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchshifttradesrequestDud struct { 
    


    

}

// Searchshifttradesrequest
type Searchshifttradesrequest struct { 
    // ReceivingScheduleId - The ID of the schedule for which to search for available shift trades
    ReceivingScheduleId string `json:"receivingScheduleId"`


    // ReceivingShiftIds - The IDs of shifts that the receiving user would potentially be willing to trade. If empty, only returns one sided trades (pick up a shift)
    ReceivingShiftIds []string `json:"receivingShiftIds"`

}

// String returns a JSON representation of the model
func (o *Searchshifttradesrequest) String() string {
    
     o.ReceivingShiftIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchshifttradesrequest) MarshalJSON() ([]byte, error) {
    type Alias Searchshifttradesrequest

    if SearchshifttradesrequestMarshalled {
        return []byte("{}"), nil
    }
    SearchshifttradesrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReceivingScheduleId string `json:"receivingScheduleId"`
        
        ReceivingShiftIds []string `json:"receivingShiftIds"`
        *Alias
    }{

        


        
        ReceivingShiftIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

