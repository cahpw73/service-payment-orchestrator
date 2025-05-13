package data

import (
	"database/sql"
	"fmt"
	"service-payment-orchestrator/config"
	"time"

	_ "github.com/godror/godror"
)

type OracleService struct {
	DB *sql.DB
}

func NewOracleService() (*OracleService, error) {
	db, err := sql.Open("godror", config.DatabaseConexion)
	if err != nil {
		return nil, fmt.Errorf("error creando conexión a Oracle: %w", err)
	}

	// Configuración del pool de conexiones
	db.SetMaxOpenConns(20)                  // Número máximo de conexiones abiertas
	db.SetMaxIdleConns(10)                  // Número máximo de conexiones inactivas
	db.SetConnMaxLifetime(30 * time.Minute) // Duración máxima de vida de una conexión

	return &OracleService{DB: db}, nil
}

func (db *OracleService) GetNroPersonaByAccount(account string) (int, error) {
	var codError, desError string
	var nroPersona int

	_, err := db.DB.Exec("BEGIN SP_GET_NRO_PERSONA_BY_ACCOUNT(:1, :2, :3, :4); END;",
		account,                    // PSTRACCOUNT (entrada)
		sql.Out{Dest: &codError},   // PSTRCODERROR (salida)
		sql.Out{Dest: &desError},   // PSTRDESERROR (salida)
		sql.Out{Dest: &nroPersona}, // PINTNROPERSONA (salida)
	)
	if err != nil {
		return 0, fmt.Errorf("error ejecutando SP_GET_NRO_PERSONA_BY_ACCOUNT: %w", err)
	}

	if codError != "" {
		return 0, fmt.Errorf("error ejecutando SP_GET_NRO_PERSONA_BY_ACCOUNT: %v", desError)
	}

	return nroPersona, nil
}
