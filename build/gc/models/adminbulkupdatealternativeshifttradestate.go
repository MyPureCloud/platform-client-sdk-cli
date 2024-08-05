package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdminbulkupdatealternativeshifttradestateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdminbulkupdatealternativeshifttradestateDud struct { 
    


    


    

}

// Adminbulkupdatealternativeshifttradestate
type Adminbulkupdatealternativeshifttradestate struct { 
    // TradeId - The ID of the trade for this alternative shift trade
    TradeId string `json:"tradeId"`


    // State - The new alternative shift trade state
    State string `json:"state"`


    // Metadata - Version metadata for this alternative shift trade
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Adminbulkupdatealternativeshifttradestate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adminbulkupdatealternativeshifttradestate) MarshalJSON() ([]byte, error) {
    type Alias Adminbulkupdatealternativeshifttradestate

    if AdminbulkupdatealternativeshifttradestateMarshalled {
        return []byte("{}"), nil
    }
    AdminbulkupdatealternativeshifttradestateMarshalled = true

    return json.Marshal(&struct {
        
        TradeId string `json:"tradeId"`
        
        State string `json:"state"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

