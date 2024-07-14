package product

import (
	"database/sql"
	"tutorial/tiago-complete-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	prod := new(types.Product)

	err := rows.Scan(
		&prod.ID,
		&prod.Name,
		&prod.Description,
		&prod.Image,
		&prod.Price,
		&prod.Quantity,
		&prod.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return prod, nil
}
