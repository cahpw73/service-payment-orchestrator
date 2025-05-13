package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"service-payment-orchestrator/app/models"
	"service-payment-orchestrator/client/redis"
	"service-payment-orchestrator/config"
)

type MiddlewareServiceInterface interface {
	MakeRequest(method, endpoint string, payload interface{}, responseModel interface{}) (interface{}, error)
	FetchNewToken() (string, error)
}

type MiddlewareService struct {
	redisClient redis.RedisServiceInterface
}

func NewMiddlewareService(redisClient redis.RedisServiceInterface) MiddlewareServiceInterface {
	return &MiddlewareService{redisClient: redisClient}
}

// Helper para construir la URL del middleware
func buildMiddlewareURL(endpoint string) (string, error) {
	middlewareURL := config.MiddlewareUrl
	if middlewareURL == "" {
		return "", fmt.Errorf("la variable de entorno MIDDLEWARE_URL no está configurada")
	}
	return fmt.Sprintf("%s/%s", middlewareURL, endpoint), nil
}

// Obtiene el token desde Redis o lo genera si es necesario
func (m *MiddlewareService) getTokenFromRedis() (string, error) {
	isUpdateTokenMiddleware, err := m.redisClient.GetData("servicePaymentMiddleware:TokenIsUpdate")
	if err != nil {
		return m.FetchNewToken()
	}

	// Convertir el valor de Redis a booleano
	boolValue, err := strconv.ParseBool(isUpdateTokenMiddleware)
	if err != nil {
		boolValue = false
	}

	// Si el token está en proceso de actualización, espera 5 segundos
	if boolValue {
		time.Sleep(5 * time.Second)
	}

	// Intentar obtener el token desde Redis
	token, err := m.redisClient.GetData("servicePaymentMiddleware:Token")
	if err != nil || token == "" {
		return m.FetchNewToken()
	}

	return token, nil
}

// Realiza una solicitud HTTP genérica y maneja la respuesta
func (m *MiddlewareService) MakeRequest(method, endpoint string, payload interface{}, responseModel interface{}) (interface{}, error) {
	url, err := buildMiddlewareURL(endpoint)
	if err != nil {
		return nil, err
	}

	var requestBody []byte
	if payload != nil {
		requestBody, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Obtener el token desde Redis y configurar el encabezado Authorization
	token, err := m.getTokenFromRedis()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("topaz-channel", config.TopazChannel)
	req.Header.Set("application-id", config.ApplicationId)
	req.Header.Set("device-id", config.DeviceId)
	req.Header.Set("device-ip", config.DeviceIp)

	// Definir Content-Type para métodos POST y PUT
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout global para el cliente
	}

	// Ejecutar la solicitud y manejar errores
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {

		fmt.Println("Recibido 401, reintentando una vez...")
		tokenRenew, err := m.FetchNewToken()
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+tokenRenew)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: código de estado %d recibido", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Manejar la estructura de respuesta general
	var tempResponse struct {
		Data    json.RawMessage `json:"data"`
		ErrCode string          `json:"errCode"`
		ErrMsg  string          `json:"errMsg"`
	}

	if err := json.Unmarshal(responseBody, &tempResponse); err != nil {
		return nil, err
	}

	// Deserializar el campo "data" en el modelo de respuesta esperado
	if err := json.Unmarshal(tempResponse.Data, responseModel); err != nil {
		return nil, err
	}

	return responseModel, nil
}

// Realiza una solicitud para obtener un nuevo token de autenticación
func (m *MiddlewareService) FetchNewToken() (string, error) {
	// Marcar que el token se está actualizando en Redis
	m.redisClient.SetData("servicePaymentMiddleware:TokenIsUpdate", "true")

	url, err := buildMiddlewareURL("payment-services-manager/oauth/token?grant_type=client_credentials")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("topaz-channel", config.TopazChannel)
	req.Header.Set("application-id", config.ApplicationId)
	req.Header.Set("secret", config.MiddlewareSecret)

	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout global para el cliente
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: código de estado %d recibido", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseModel models.ResponseTokenMiddleware
	if err := json.Unmarshal(responseBody, &responseModel); err != nil {
		return "", err
	}

	m.redisClient.SetData("servicePaymentMiddleware:Token", responseModel.AccessToken)
	m.redisClient.SetData("servicePaymentMiddleware:TokenIsUpdate", "false")

	return responseModel.AccessToken, nil
}
