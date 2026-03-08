package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Mapa para contar intentos (En el futuro usaremos Redis)
var (
	incidentCount = make(map[string]int)
	mutex         sync.Mutex
)

func monitorHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	
	mutex.Lock()
	incidentCount[ip]++
	count := incidentCount[ip]
	mutex.Unlock()

	fmt.Printf("[ALERTA] Intento de acceso desde IP: %s | Total: %d\n", ip, count)

	if count > 5 {
		fmt.Printf("[FASE 2 ACTIVADA] Bloqueando IP %s y generando señuelo...\n", ip)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Amenaza detectada"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Registro de monitoreo actualizado"))
}

func main() {
	http.HandleFunc("/report", monitorHandler)
	fmt.Println("Monitor de Seguridad (Go) escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
