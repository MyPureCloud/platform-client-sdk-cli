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


    


    


    

}

// String returns a JSON representation of the model
func (o *Contactlistdivisionview) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.ColumnNames = []string{""} 
    
    
    
     o.PhoneColumns = []Contactphonenumbercolumn{{}} 
    
    
    
    
    
    
    
    

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
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        ColumnNames: []string{""},
        

        

        
        PhoneColumns: []Contactphonenumbercolumn{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

