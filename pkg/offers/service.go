// реализация работы с Postgres
package offers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// пул соеденений к базе данных
type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

// сущность Offer, то какие данные хранятся внутри сервиса
type Offer struct {
	ID      int64  `json:"id"`
	Company string `json:"company"`
	Percent string `json:"percent"`
	Comment string `json:"comment"`
}

func (s *Service) All(ctx context.Context) ([]*Offer, error) {
	// пустой слайс предложений
	items := make([]*Offer, 0)

	rows, err := s.pool.Query(ctx, "SELECT id, company, percent, comment FROM offers")
	if err != nil {
		if err == pgx.ErrNoRows {
			return items, nil
		}

		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := &Offer{}
		err = rows.Scan(&item.ID, &item.Company, &item.Percent, &item.Comment)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return items, nil
}

func (s *Service) ByID(ctx context.Context, id int64) (*Offer, error) {
	item := &Offer{ID: id}
	err := s.pool.QueryRow(
		ctx,
		`Select company, percent, comment FROM offers WHERE id = $1`,
		id,
	).Scan(&item.Company, &item.Percent, &item.Comment)

	if err != nil {
		log.Println(err)
	}

	return item, nil
}

func (s *Service) Save(ctx context.Context, itemToSave *Offer) (*Offer, error) {
	if itemToSave.ID == 0 {
		err := s.pool.QueryRow(
			ctx,
			`INSERT INTO offers (company, percent, comment) VALUES($1, $2, $3) RETURNING id`,
			itemToSave.Company, itemToSave.Percent, itemToSave.Comment,
		).Scan(&itemToSave)
		if err != nil {
			return nil, err
		}
		return itemToSave, nil
	}

	tag, err := s.pool.Exec(
		ctx,
		`UPDATE offers SET company=$2, percent = $3, comment = $4, WHERE id = $1`,
		itemToSave.ID, itemToSave.Company, itemToSave.Percent, itemToSave.Comment,
	)
	if err != nil {
		return nil, err
	}

	if tag.RowsAffected() != 1 {
		return nil, errors.New("No rows updated")
	}

	return itemToSave, nil
}

func (s Service) Delete(ctx context.Context, id int64) (*Offer, error) {
	//var offer Offer
	offer := &Offer{ID: id}
	err := s.pool.QueryRow(
		ctx,
		`DELETE FROM offers 
			WHERE id = $1 
			RETURNING id, company, percent, comment`,
		offer.ID,
	).Scan(&offer.Company, &offer.Percent, &offer.Comment)
	if err != nil {
		return nil, err
	}

	return offer, nil
}
