package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonesrebootMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonesrebootDud struct { 
    


    

}

// Phonesreboot
type Phonesreboot struct { 
    // PhoneIds - The list of phone Ids to reboot.
    PhoneIds []string `json:"phoneIds"`


    // SiteId - ID of the site for which to reboot all phones at that site. no.active.edge and phone.cannot.resolve errors are ignored.
    SiteId string `json:"siteId"`

}

// String returns a JSON representation of the model
func (o *Phonesreboot) String() string {
    
    
     o.PhoneIds = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonesreboot) MarshalJSON() ([]byte, error) {
    type Alias Phonesreboot

    if PhonesrebootMarshalled {
        return []byte("{}"), nil
    }
    PhonesrebootMarshalled = true

    return json.Marshal(&struct { 
        PhoneIds []string `json:"phoneIds"`
        
        SiteId string `json:"siteId"`
        
        *Alias
    }{
        

        
        PhoneIds: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

