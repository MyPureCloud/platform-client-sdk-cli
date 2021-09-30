package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentationresultDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Documentationresult
type Documentationresult struct { 
    // Id - The globally unique identifier for the object.
    Id int `json:"id"`


    // Categories - The category of the documentation entity. Will be returned in responses for certain entities.
    Categories []int `json:"categories"`


    // Description - The description of the documentation entity. Will be returned in responses for certain entities.
    Description string `json:"description"`


    // Content - The text or html content for the documentation entity. Will be returned in responses for certain entities.
    Content string `json:"content"`


    // Excerpt - The excerpt of the documentation entity. Will be returned in responses for certain entities.
    Excerpt string `json:"excerpt"`


    // Link - URL link for the documentation entity. Will be returned in responses for certain entities.
    Link string `json:"link"`


    // Modified - The modified date for the documentation entity. Will be returned in responses for certain entities. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Modified time.Time `json:"modified"`


    // Name - The name of the documentation entity. Will be returned in responses for certain entities.
    Name string `json:"name"`


    // Service - The service of the documentation entity. Will be returned in responses for certain entities.
    Service []int `json:"service"`


    // Slug - The slug of the documentation entity. Will be returned in responses for certain entities.
    Slug string `json:"slug"`


    // Title - The title of the documentation entity. Will be returned in responses for certain entities.
    Title string `json:"title"`


    // GetType - The search type. Will be returned in responses for certain entities.
    GetType string `json:"get_type"`


    // FacetFeature - The facet feature of the documentation entity. Will be returned in responses for certain entities.
    FacetFeature []int `json:"facet_feature"`


    // FacetRole - The facet role of the documentation entity. Will be returned in responses for certain entities.
    FacetRole []int `json:"facet_role"`


    // FacetService - The facet service of the documentation entity. Will be returned in responses for certain entities.
    FacetService []int `json:"facet_service"`


    // FaqCategories - The faq categories of the documentation entity. Will be returned in responses for certain entities.
    FaqCategories []int `json:"faq_categories"`


    // ReleasenoteCategory - The releasenote category of the documentation entity. Will be returned in responses for certain entities.
    ReleasenoteCategory []int `json:"releasenote_category"`


    // ReleasenoteTag - The releasenote tag of the documentation entity. Will be returned in responses for certain entities.
    ReleasenoteTag []int `json:"releasenote_tag"`


    // ServiceArea - The service area of the documentation entity. Will be returned in responses for certain entities.
    ServiceArea []int `json:"service-area"`


    // VideoCategories - The video categories of the documentation entity. Will be returned in responses for certain entities.
    VideoCategories []int `json:"video_categories"`

}

// String returns a JSON representation of the model
func (o *Documentationresult) String() string {
    
    
    
    
    
    
     o.Categories = []int{0} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Service = []int{0} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FacetFeature = []int{0} 
    
    
    
     o.FacetRole = []int{0} 
    
    
    
     o.FacetService = []int{0} 
    
    
    
     o.FaqCategories = []int{0} 
    
    
    
     o.ReleasenoteCategory = []int{0} 
    
    
    
     o.ReleasenoteTag = []int{0} 
    
    
    
     o.ServiceArea = []int{0} 
    
    
    
     o.VideoCategories = []int{0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentationresult) MarshalJSON() ([]byte, error) {
    type Alias Documentationresult

    if DocumentationresultMarshalled {
        return []byte("{}"), nil
    }
    DocumentationresultMarshalled = true

    return json.Marshal(&struct { 
        Id int `json:"id"`
        
        Categories []int `json:"categories"`
        
        Description string `json:"description"`
        
        Content string `json:"content"`
        
        Excerpt string `json:"excerpt"`
        
        Link string `json:"link"`
        
        Modified time.Time `json:"modified"`
        
        Name string `json:"name"`
        
        Service []int `json:"service"`
        
        Slug string `json:"slug"`
        
        Title string `json:"title"`
        
        GetType string `json:"get_type"`
        
        FacetFeature []int `json:"facet_feature"`
        
        FacetRole []int `json:"facet_role"`
        
        FacetService []int `json:"facet_service"`
        
        FaqCategories []int `json:"faq_categories"`
        
        ReleasenoteCategory []int `json:"releasenote_category"`
        
        ReleasenoteTag []int `json:"releasenote_tag"`
        
        ServiceArea []int `json:"service-area"`
        
        VideoCategories []int `json:"video_categories"`
        
        *Alias
    }{
        

        

        

        
        Categories: []int{0},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Service: []int{0},
        

        

        

        

        

        

        

        

        
        FacetFeature: []int{0},
        

        

        
        FacetRole: []int{0},
        

        

        
        FacetService: []int{0},
        

        

        
        FaqCategories: []int{0},
        

        

        
        ReleasenoteCategory: []int{0},
        

        

        
        ReleasenoteTag: []int{0},
        

        

        
        ServiceArea: []int{0},
        

        

        
        VideoCategories: []int{0},
        

        
        Alias: (*Alias)(u),
    })
}

