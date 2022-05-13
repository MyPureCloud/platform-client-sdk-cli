package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchshifttraderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchshifttraderesponseDud struct { 
    


    


    

}

// Searchshifttraderesponse
type Searchshifttraderesponse struct { 
    // Trade - A trade which matches search criteria
    Trade Shifttraderesponse `json:"trade"`


    // MatchingReceivingShiftIds - IDs of shifts which match the search criteria
    MatchingReceivingShiftIds []string `json:"matchingReceivingShiftIds"`


    // Preview - A preview of what the shift trade would look like if matched
    Preview Shifttradepreviewresponse `json:"preview"`

}

// String returns a JSON representation of the model
func (o *Searchshifttraderesponse) String() string {
    
     o.MatchingReceivingShiftIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchshifttraderesponse) MarshalJSON() ([]byte, error) {
    type Alias Searchshifttraderesponse

    if SearchshifttraderesponseMarshalled {
        return []byte("{}"), nil
    }
    SearchshifttraderesponseMarshalled = true

    return json.Marshal(&struct {
        
        Trade Shifttraderesponse `json:"trade"`
        
        MatchingReceivingShiftIds []string `json:"matchingReceivingShiftIds"`
        
        Preview Shifttradepreviewresponse `json:"preview"`
        *Alias
    }{

        


        
        MatchingReceivingShiftIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

