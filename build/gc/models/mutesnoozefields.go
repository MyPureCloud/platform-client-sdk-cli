package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MutesnoozefieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MutesnoozefieldsDud struct { 
    


    

}

// Mutesnoozefields
type Mutesnoozefields struct { 
    // DateStart - The start date of the mute/snooze period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end date of the mute/snooze period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`

}

// String returns a JSON representation of the model
func (o *Mutesnoozefields) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mutesnoozefields) MarshalJSON() ([]byte, error) {
    type Alias Mutesnoozefields

    if MutesnoozefieldsMarshalled {
        return []byte("{}"), nil
    }
    MutesnoozefieldsMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

