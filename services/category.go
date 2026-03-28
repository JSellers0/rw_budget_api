package services

import "database/sql"

type Category struct {
	ID   int    `json:"categoryid" form:"categoryid"`
	Name string `json:"category_name" form:"category_name" binding:"required"`
}

type CategoryService interface {
	CreateCategory(Category) (*int64, error)
	ReadAllCategories() ([]*Category, error)
	ReadCategoryByID(string) (*Category, error)
	ReadCategoriesByName(string) ([]*Category, error)
	UpdateCategory(Category) error
	DeleteCategory(string) error
}

type categoryService struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) CategoryService {
	return &categoryService{db: db}
}

func (s *categoryService) CreateCategory(new_category Category) (*int64, error) {
	query := "INSERT INTO category (category_name) VALUES ('?');"
	res, err := s.db.Exec(query, new_category.Name)
	if err != nil {
		return nil, err
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func (s *categoryService) ReadCategoryByID(id string) (record *Category, err error) {
	query := buildGetCategoryQuery("WHERE categoryid = ?\n;")
	var category Category
	if err := s.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *categoryService) ReadCategoriesByName(name string) (record []*Category, err error) {
	query := buildGetCategoryQuery("WHERE category_name = '?'\n;")
	var category Category
	if err := s.db.QueryRow(query, name).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return nil, err
	}
	return []*Category{&category}, nil
}

func (s *categoryService) ReadAllCategories() (record []*Category, err error) {
	categories := []*Category{}
	results, err := s.db.Query(buildGetCategoryQuery(""))
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var category Category
		err = results.Scan(
			&category.ID, &category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (s *categoryService) UpdateCategory(mod_category Category) error {
	//ToDo: UpdateCategory
	_, err := s.ReadCategoriesByName(mod_category.Name)
	if err == nil {
		// ToDo: Custom Error?
		return nil
	}
	// ToDo: Set Date and User for update
	query := "UPDATE category SET category_name = '?', update_date='?', update_by='?'"
	_, up_err := s.db.Exec(query, mod_category.Name, "", "")
	if up_err != nil {
		return up_err
	}
	return nil
}

func (s *categoryService) DeleteCategory(id string) error {
	_, get_err := s.ReadCategoryByID(id)
	if get_err != nil {
		return get_err

	}
	query := "DELETE FROM category WHERE categoryid = ?;"
	_, del_err := s.db.Exec(query, id)
	if del_err != nil {
		return del_err
	}

	return nil
}

func buildGetCategoryQuery(query_add string) string {
	return "SELECT categoryid, category_name\nFROM category\n" + query_add
}
