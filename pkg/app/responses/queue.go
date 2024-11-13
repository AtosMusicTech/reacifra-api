package responses

import "cifra-api/pkg/domain/entity"

type QueueResponse struct {
	Id    uint             `json:"id"`
	Cifra []*CifraResponse `json:"cifras"`
}

func ConvertToQueueResponse(queue *entity.Queue) *QueueResponse {
	return &QueueResponse{
		Id:    queue.Id,
		Cifra: ConvertToCifrasResponse(queue.Cifras),
	}
}
