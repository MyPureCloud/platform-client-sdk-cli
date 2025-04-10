package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistdivisionviewDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    ImportStatus Importstatus `json:"importStatus"`


    Size int `json:"size"`


    SelfUri string `json:"selfUri"`

}

// Contactlistdivisionview
type Contactlistdivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // ColumnNames - The names of the contact data columns.
    ColumnNames []string `json:"columnNames"`


    // PhoneColumns - Indicates which columns are phone numbers.
    PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`


    // EmailColumns - Indicates which columns are email addresses.
    EmailColumns []Emailcolumn `json:"emailColumns"`


    // WhatsAppColumns - Indicates which columns are whatsApp contacts.
    WhatsAppColumns []Whatsappcolumn `json:"whatsAppColumns"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Contactlistdivisionview) String() string {
    
    
     o.ColumnNames = []string{""} 
     o.PhoneColumns = []Contactphonenumbercolumn{{}} 
     o.EmailColumns = []Emailcolumn{{}} 
     o.WhatsAppColumns = []Whatsappcolumn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Contactlistdivisionview

    if ContactlistdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    ContactlistdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        ColumnNames []string `json:"columnNames"`
        
        PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`
        
        EmailColumns []Emailcolumn `json:"emailColumns"`
        
        WhatsAppColumns []Whatsappcolumn `json:"whatsAppColumns"`
        *Alias
    }{

        


        


        


        
        ColumnNames: []string{""},
        


        
        PhoneColumns: []Contactphonenumbercolumn{{}},
        


        
        EmailColumns: []Emailcolumn{{}},
        


        
        WhatsAppColumns: []Whatsappcolumn{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

