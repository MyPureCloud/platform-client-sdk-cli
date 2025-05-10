package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillinginvoiceitemlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillinginvoiceitemlistingDud struct { 
    


    


    


    

}

// Billinginvoiceitemlisting
type Billinginvoiceitemlisting struct { 
    // Entities
    Entities []Billinginvoiceitem `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Billinginvoiceitemlisting) String() string {
     o.Entities = []Billinginvoiceitem{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billinginvoiceitemlisting) MarshalJSON() ([]byte, error) {
    type Alias Billinginvoiceitemlisting

    if BillinginvoiceitemlistingMarshalled {
        return []byte("{}"), nil
    }
    BillinginvoiceitemlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Billinginvoiceitem `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Billinginvoiceitem{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

