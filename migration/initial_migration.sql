BEGIN;

CREATE TABLE IF NOT EXISTS product (
  id serial PRIMARY key, 
  name VARCHAR, 
  price BIGINT
);

CREATE TABLE IF NOT EXISTS category (
  id serial PRIMARY key, 
  name VARCHAR,
  parent_category INT
);

CREATE TABLE IF NOT EXISTS product_category_mapping (
  id serial PRIMARY key, 
  product_id INT, 
  category_id INT, 
  FOREIGN key(product_id) REFERENCES product(id), 
  FOREIGN key(category_id) REFERENCES category(id)
);

END;
