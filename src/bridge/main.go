package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8" // Para manejar los tokens rotativos
)

var ctx = context.Background()

func main() {
	fmt.Println("🚀 CORE-FRÍO: Bridge (Fase 2) Operativo. Monitoreando Redis...")

	// 1. Conexión a Redis (El buzón de mensajes rápidos)
	rdb := redis.NewClient(&redis.Options{
		Addr: "cache_redis:6379", // Nombre del servicio en tu Docker
	})

	// 2. Bucle de Reacción (Goroutine)
	for {
		// El Bridge "escucha" si el Auditor (Rust) reporta una IP peligrosa
		val, err := rdb.Get(ctx, "last_alert_risk").Result()
		if err == nil && val == "3" {
			fmt.Println("⚠️ RIESGO CRÍTICO DETECTADO: Iniciando Rotación de Tokens...")
			
			// Lógica de Rotación: Cambiamos la "llave" maestra en Redis
			newToken := fmt.Sprintf("CORED-COLD-%d", time.Now().Unix())
			rdb.Set(ctx, "active_session_token", newToken, 5*time.Minute)
			
			fmt.Printf("🔐 Nuevo Token Rotativo Generado: %s\n", newToken)
			
			// Limpiamos la alerta para esperar la siguiente
			rdb.Del(ctx, "last_alert_risk")
		}

		time.Sleep(1 * time.Second) // Pausa para no saturar el CPU
	}
}
