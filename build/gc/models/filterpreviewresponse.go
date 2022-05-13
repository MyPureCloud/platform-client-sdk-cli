package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FilterpreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FilterpreviewresponseDud struct { 
    


    


    

}

// Filterpreviewresponse
type Filterpreviewresponse struct { 
    // FilteredContacts
    FilteredContacts int `json:"filteredContacts"`


    // TotalContacts
    TotalContacts int `json:"totalContacts"`


    // Preview
    Preview []Dialercontact `json:"preview"`

}

// String returns a JSON representation of the model
func (o *Filterpreviewresponse) String() string {
    
    
     o.Preview = []Dialercontact{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Filterpreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Filterpreviewresponse

    if FilterpreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    FilterpreviewresponseMarshalled = true

    return json.Marshal(&struct {
        
        FilteredContacts int `json:"filteredContacts"`
        
        TotalContacts int `json:"totalContacts"`
        
        Preview []Dialercontact `json:"preview"`
        *Alias
    }{

        


        


        
        Preview: []Dialercontact{{}},
        

        Alias: (*Alias)(u),
    })
}

