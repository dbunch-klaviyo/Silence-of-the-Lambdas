package handlers

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	dynamodb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
)

type Handler struct {
	s silence.Client
}

func NewHandler(s silence.Client) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) GetKiller(req events.APIGatewayProxyRequest, tableName string, db *dynamodb.SilenceOfTheLambsDB) *events.APIGatewayProxyResponse {
	fullName, ok := req.PathParameters["fullName"]
	if !ok {
		response := BadRequest400(events.APIGatewayProxyResponse{}, "Killer", "fullName")
		return &response
	}

	if strings.Contains(fullName, " ") {
		response := BadRequest400(events.APIGatewayProxyResponse{}, "Killer", "fullName")
		return &response
	}

	killer, err := h.s.ReturnKillerByFullName(fullName, tableName, db)
	if err != nil {
		log.Printf("Error getting killer: %v", err)
		response := InternalError500(events.APIGatewayProxyResponse{}, "Killer", err)
		return &response
	}

	if killer == nil {
		response := NotFound404(events.APIGatewayProxyResponse{}, "Killer")
		return &response
	}

	response := OK200(events.APIGatewayProxyResponse{}, killer)
	return &response
}
