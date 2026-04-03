package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradematchresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradematchresponseitemDud struct { 
    


    

}

// Shifttradematchresponseitem
type Shifttradematchresponseitem struct { 
    // ShiftIds - The IDs of the receiving shift trades which match the initiating shift trade
    ShiftIds []string `json:"shiftIds"`


    // Preview - A preview of what the shift trade would look like if matched
    Preview Shifttradepreviewresponse `json:"preview"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchresponseitem) String() string {
     o.ShiftIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradematchresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradematchresponseitem

    if ShifttradematchresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradematchresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        ShiftIds []string `json:"shiftIds"`
        
        Preview Shifttradepreviewresponse `json:"preview"`
        *Alias
    }{

        
        ShiftIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

