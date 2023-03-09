package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppendtodncactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppendtodncactionsettingsDud struct { 
    


    


    

}

// Appendtodncactionsettings
type Appendtodncactionsettings struct { 
    // Expire - Whether to expire the record appended to the DNC list.
    Expire bool `json:"expire"`


    // ExpirationDuration - If 'expire' is set to true, how long to keep the record.
    ExpirationDuration string `json:"expirationDuration"`


    // ListType - The Dnc List Type to append entries to
    ListType string `json:"listType"`

}

// String returns a JSON representation of the model
func (o *Appendtodncactionsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appendtodncactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Appendtodncactionsettings

    if AppendtodncactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    AppendtodncactionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Expire bool `json:"expire"`
        
        ExpirationDuration string `json:"expirationDuration"`
        
        ListType string `json:"listType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

