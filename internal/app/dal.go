package app

// DBGetStatus get status from DB.
func (db *DB) DBGetStatus() (Status, error) {
	err := db.Ping()
	if err != nil {
		return Status{DBStatus: "nok"}, err
	}
	return Status{DBStatus: "ok"}, nil
}

const addCategoryQuery = `
		INSERT INTO category(
			name,
			parent_category
		)
		VALUES(
			$1,
			$2
		)
		RETURNING id; 
`

// DBAddCategory add category into DB.
func (db *DB) DBAddCategory(c Category) (int, error) {
	var id int
	err := db.QueryRow(addCategoryQuery, c.Name, c.ParentCategory).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const addProductQuery = `
		INSERT INTO product(
			name,
			price
		)
		VALUES(
			$1,
			$2
		)
		RETURNING id; 
`

const addProductCategoryMappingQuery = `
		INSERT INTO product_category_mapping(
			product_id,
			category_id
		)
		VALUES(
			$1,
			$2
		); 
`

// DBAddProduct add product into DB.
func (db *DB) DBAddProduct(p Product) (int, error) {
	var id int
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	err = tx.QueryRow(addProductQuery, p.Name, p.Price).Scan(&id)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	for _, v := range p.Categories {
		_, err := tx.Exec(addProductCategoryMappingQuery, id, v)
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				return 0, rbErr
			}
			return 0, err
		}
	}
	return id, tx.Commit()
}

const updateProductQuery = `
		UPDATE product
		SET 
			name = $1,
			price = $2
		WHERE 
			id = $3
		RETURNING id; 
`

// DBUpdateProduct update product in DB.
func (db *DB) DBUpdateProduct(p Product) (int, error) {
	var id int
	err := db.QueryRow(updateProductQuery, p.Name, p.Price, p.ID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}

const getProductByCategoryQuery = `
		SELECT 
			p.id,
			p.name,
			p.price
		FROM   product AS p
			JOIN product_category_mapping AS pcm
				ON p.id = pcm.product_id
		WHERE  pcm.category_id = $1; 
`

// DBGetProductByCategory get product by category in DB.
func (db *DB) DBGetProductByCategory(categoryID string) ([]*Product, error) {
	var products []*Product
	rows, err := db.Query(getProductByCategoryQuery, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

const getAllCategoryQuery = `
		SELECT 
			id,
			name,
			parent_category
		FROM   category; 
`

// DBGetAllCategory get all category in DB.
func (db *DB) DBGetAllCategory() ([]*Category, error) {
	var category []*Category
	rows, err := db.Query(getAllCategoryQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name, &c.ParentCategory)
		if err != nil {
			return nil, err
		}
		category = append(category, &c)
	}
	for _, v := range category {
		for _, vv := range category {
			if v.ID == vv.ParentCategory {
				v.ChildCategories = append(v.ChildCategories, *vv)
			}
		}
	}
	return category, nil
}
