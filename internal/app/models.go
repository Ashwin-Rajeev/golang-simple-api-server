package app

// Status hold the status information
type Status struct {
	DBStatus string `json:"db_status"`
}

// Product hold the product data.
type Product struct {
	ID         int        `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Price      int64      `json:"price,omitempty"`
	Categories []int      `json:"category_ids,omitempty"`
	Category   []Category `json:"category,omitempty"`
}

// Category hold the category data.
type Category struct {
	ID              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	ChildCategories []Category `json:"child_categories,omitempty"`
	ParentCategory  int        `json:"parent_category,omitempty"`
}
