package freshservice

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type CustomObjectService[T any] struct {
	client *Client
}

func (svc *CustomObjectService[T]) CreateCustomObjectRecord(ctx context.Context, customObjectID int, request CreateCustomObjectRecordRequest[T]) (CreateCustomObjectRecordResponse[T], error) {
	var response CreateCustomObjectRecordResponse[T]
	_, err := svc.client.Post(fmt.Sprintf("objects/%d/records", customObjectID), request, &response)
	if err != nil {
		return CreateCustomObjectRecordResponse[T]{}, err
	}
	return response, nil
}

func (svc *CustomObjectService[T]) ListCustomObjectRecords(ctx context.Context, customObjectID int, pageSize int) (ListCustomObjectRecordsResponse[T], error) {
	u := fmt.Sprintf("objects/%d/records", customObjectID)
	if pageSize > 0 {
		u = fmt.Sprintf("objects/%d/records?page_size=%d", customObjectID, pageSize)
	}
	var response ListCustomObjectRecordsResponse[T]
	_, err := svc.client.Get(u, &response)
	if err != nil {
		return ListCustomObjectRecordsResponse[T]{}, err
	}
	return response, nil
}

func (svc *CustomObjectService[T]) UpdateCustomObjectRecord(ctx context.Context, customObjectID int, customObjectRecordID int, request UpdateCustomObjectRecordRequest[T]) (UpdateCustomObjectRecordResponse[T], error) {
	var response UpdateCustomObjectRecordResponse[T]
	_, err := svc.client.Put(fmt.Sprintf("objects/%d/records/%d", customObjectID, customObjectRecordID), request, &response)
	if err != nil {
		return UpdateCustomObjectRecordResponse[T]{}, err
	}
	return response, nil
}

func (svc *CustomObjectService[T]) DeleteCustomObjectRecord(ctx context.Context, customObjectID int, customObjectRecordID int) error {
	success, _, err := svc.client.Delete(fmt.Sprintf("objects/%d/records/%d", customObjectID, customObjectRecordID))
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("failed to delete custom object record: %d in custom object: %d", customObjectRecordID, customObjectID)
	}
	return nil
}

type customObjectRecordRequest[T any] struct {
	Data T `json:"data"`
}

type customObjectRecordResponse[T any] struct {
	CustomObject CustomObjectRecord[T] `json:"custom_object"`
}

type CreateCustomObjectRecordRequest[T any] struct {
	customObjectRecordRequest[T]
}

type CreateCustomObjectRecordResponse[T any] struct {
	customObjectRecordResponse[T]
}

type ListCustomObjectRecordsResponse[T any] struct {
	Records []CustomObjectRecord[T] `json:"records"`
}

type UpdateCustomObjectRecordRequest[T any] struct {
	customObjectRecordRequest[T]
}

type UpdateCustomObjectRecordResponse[T any] struct {
	customObjectRecordResponse[T]
}

type CustomObjectRecord[T any] struct {
	Data     T        `json:"data"`
	NextPage *url.URL `json:"next_page_link"`
}

type CustomObjectRecordMetadata struct {
	CreatedAt time.Time `json:"bo_created_at,omitempty"`
	DisplayId int       `json:"bo_display_id,omitempty"`
	UpdatedAt time.Time `json:"bo_updated_at,omitempty"`
}
