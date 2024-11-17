package store

import (
	"fmt"
	"strings"

	"github.com/VitalyCone/kuznecov_communities_api/internal/app/model"
)

type PublicationRepository struct {
	store *Store
}

func (r *PublicationRepository) Create(m *model.Publication) error {
	fileIds := make([]string, len(m.FileIds))
    for i, id := range m.FileIds {
        fileIds[i] = fmt.Sprintf("%d", id)
    }
    
    fileIdsArray := "{" + strings.Join(fileIds, ",") + "}"

	if err := r.store.db.QueryRow(
		"INSERT INTO publications (text, file_ids) "+
			"VALUES ($1, $2) RETURNING id, created_at",
		m.Text, fileIdsArray).Scan(&m.ID, &m.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (r *PublicationRepository) Get(id int) (*model.Publication, error) {
	var m model.Publication
	var fileIdsString string
	m.FileIds = make([]int, 0)

	m.ID = id
	if err := r.store.db.QueryRow(
		"SELECT text, created_at, file_ids FROM publications WHERE id = $1",
		id).Scan(&m.Text, &m.CreatedAt, &fileIdsString); err != nil {
		return nil,err
	}

	if fileIdsString != "" {
        fileIdsString = strings.Trim(fileIdsString, "{}") // Убираем фигурные скобки
        ids := strings.Split(fileIdsString, ",")         // Разделяем по запятой
        for _, idStr := range ids {
            var idInt int
            if _, err := fmt.Sscanf(idStr, "%d", &idInt); err == nil {
                m.FileIds = append(m.FileIds, idInt) // Добавляем в срез FileIds
            }
        }
    }
	
	return &m,nil
}

func (r *PublicationRepository) Delete(id int) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM publications WHERE id = $1;",
		id); err != nil {
		return err
	}
	return nil
}