package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchunmatchedshifttraderesponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchunmatchedshifttraderesponseitemDud struct { 
    


    

}

// Searchunmatchedshifttraderesponseitem
type Searchunmatchedshifttraderesponseitem struct { 
    // Trade - A trade which matches search criteria
    Trade Shifttraderesponseitem `json:"trade"`


    // ReceivingShiftMatches - The shifts that match the search criteria
    ReceivingShiftMatches []Shifttradematchresponseitem `json:"receivingShiftMatches"`

}

// String returns a JSON representation of the model
func (o *Searchunmatchedshifttraderesponseitem) String() string {
    
     o.ReceivingShiftMatches = []Shifttradematchresponseitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchunmatchedshifttraderesponseitem) MarshalJSON() ([]byte, error) {
    type Alias Searchunmatchedshifttraderesponseitem

    if SearchunmatchedshifttraderesponseitemMarshalled {
        return []byte("{}"), nil
    }
    SearchunmatchedshifttraderesponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Trade Shifttraderesponseitem `json:"trade"`
        
        ReceivingShiftMatches []Shifttradematchresponseitem `json:"receivingShiftMatches"`
        *Alias
    }{

        


        
        ReceivingShiftMatches: []Shifttradematchresponseitem{{}},
        

        Alias: (*Alias)(u),
    })
}

