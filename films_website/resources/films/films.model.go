package films

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId             int       `gorm:"column:film_id;primaryKey"`
	Title              string    `gorm:"type:varchar(128);not null"`
	Description        string    `gorm:"type:text"`
	ReleaseYear        int       `gorm:"type:year"`
	LanguageId         int       `gorm:"type:tinyint;not null"`
	OriginalLanguageId int       `gorm:"type:tinyint;default:null"`
	RentalDuration     float64   `gorm:"type:tinyint;not null;default:4.99"`
	RentalRate         float64   `gorm:"type:decimal(4,2);"`
	Length             int       `gorm:"type:smallint"`
	ReplacementCost    float64   `gorm:"type:decimal(5,2);noy null;default:19.99"`
	Rating             string    `gorm:"type:enum('G','PG','PG-13','R','NC-17');default:'G'"`
	SpecialFeatures    string    `gorm:"type:set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes')"`
	LastUpdate         time.Time `gorm:"autoCreateTime"`
	// Categories         []Category `gorm:"many2many:film_category;foreignKey:FilmId;joinForeignKey:FilmId;References:CategoryId;joinReferences:CategoryId;constraint:onDelete:CASCADE;onUpdate:CASCADE;"`
	Categories []Category `gorm:"many2many:film_category;foreignKey:FilmId;joinForeignKey:FilmId;References:CategoryId;joinReferences:CategoryId;"`
	//Categories []Category `gorm:"many2many:film_category;foreignKey:FilmId;joinForeignKey:FilmId;References:CategoryId;joinReferences:CategoryId;"`

	// constraint:onDelete:CASCADE;onUpdate:CASCADE
	// Categories         []Category `gorm:"foreignKey:CategoryId"`
	//Categories         []Category `gorm:"foreignKey:CategoryId;constraint:onDelete:CASCADE;"`
	//Categories []Category `gorm:"many2many:film_category;"`
	// Categories []*Category `gorm:"many2many:film_category;"`
}

// define the table know
func (Film) TableName() string {
	return "film"
}

type Category struct {
	CategoryId int       `gorm:"column:category_id;primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(25)"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
	// Films      []*Film   `gorm:"many2many:film_category"`
	//  `gorm:"many2many:film_category;constraint:onDelete:CASCADE;"`
}

func (Category) TableName() string {
	return "category"
}

type FilmCategory struct {
	FilmId             int       `gorm:"column:film_id;primaryKey"`
	CategoryId         int       `gorm:"column:categoryid;primaryKey"`
	LastUpdate         time.Time `gorm:"type:timestamp"`
	FilmFilmId         int       `gorm:"type:tinyint"`
	CategoryCategoryId int       `gorm:"type:tinyint"`
	FilmReferId        int       `gorm:"type:tinyint"`
	CategoryRefer      int       `gorm:"type:tinyint"`
}

func (FilmCategory) TableName() string {
	return "film_category"
}

type FilmRequest struct {
	*Film
}

func (f *FilmRequest) Bind(r *http.Request) error {
	// check existing fields of structure
	if f.Film == nil {
		return errors.New("missing required Film fields")
	}

	// cast title to upper case
	f.Film.Title = strings.ToUpper(f.Film.Title)
	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
