package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MergecontactsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MergecontactsrequestDud struct { 
    


    

}

// Mergecontactsrequest
type Mergecontactsrequest struct { 
    // ContactIds - The IDs of all contacts involved in the merge operation (must be between 2 and 25).
    ContactIds []string `json:"contactIds"`


    // ValueOverride - Override data to set for specific Contact fields after a merge. Any null fields in `valueOverride` will not replace existing data.
    ValueOverride Externalcontact `json:"valueOverride"`

}

// String returns a JSON representation of the model
func (o *Mergecontactsrequest) String() string {
     o.ContactIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mergecontactsrequest) MarshalJSON() ([]byte, error) {
    type Alias Mergecontactsrequest

    if MergecontactsrequestMarshalled {
        return []byte("{}"), nil
    }
    MergecontactsrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContactIds []string `json:"contactIds"`
        
        ValueOverride Externalcontact `json:"valueOverride"`
        *Alias
    }{

        
        ContactIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

