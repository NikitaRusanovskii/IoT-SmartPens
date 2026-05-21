package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"smartPens/internal/domain"
	"smartPens/internal/repository"
)

type Sender struct {
	URL    string
	db     *repository.LessonRepository
	client *http.Client
}

func (s *Sender) Serialize(wr *domain.WorksRequest) ([]byte, error) {
	mwr, err := json.Marshal(wr)
	if err != nil {
		return nil, err
	}
	return mwr, nil
}

func (s *Sender) CreateRequest(jsonData []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (s *Sender) SendWorksTo(ctx context.Context, lessonID int) error {
	wr, err := s.db.CollectWorks(ctx, lessonID)
	if err != nil {
		return err
	}
	mwr, err := s.Serialize(wr)
	if err != nil {
		return err
	}
	req, err := s.CreateRequest(mwr)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		fmt.Println("Запрос успешно отправлен!")
	} else {
		fmt.Printf("Сервер вернул ошибку, статус: %d\n", resp.StatusCode)
	}

}
