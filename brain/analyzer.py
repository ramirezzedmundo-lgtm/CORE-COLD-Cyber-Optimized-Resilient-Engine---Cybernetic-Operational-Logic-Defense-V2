import time

/**
 * CORE-COLD V2 - Layer 5 (The Brain)
 * Sistema de Inteligencia Defensiva contra Agentes Persistentes.
 */

class CoreColdBrain:
    def __init__(self):
        # Memoria de amenazas (lo que aprendimos del ataque previo)
        self.blacklist_signatures = ["fido-v2", "fenix-alpha", "service_worker_payload"]
        self.critical_limit_mb = 205.0
        self.alert_level = "NORMAL"

    def audit_environment(self, process_name, memory_usage_mb):
        print(f"[🧠] BRAIN: Escaneando proceso '{process_name}'...")
        
        # Deteccion de 'Fenix Alpha' por volumen de RAM
        if memory_usage_mb >= self.critical_limit_mb:
            self.alert_level = "CRITICAL"
            return f"[🚨] ALERTA DE IA: Proceso de {memory_usage_mb}MB detectado. Bloqueando via Gatekeeper..."

        # Deteccion por firma de texto
        for threat in self.blacklist_signatures:
            if threat in process_name.lower():
                return f"[🚫] BLOQUEO: Firma de amenaza '{threat}' identificada."

        return "[✅] ESTADO: Entorno seguro. CPU Frio."

# --- EJECUCION DE PRUEBA EN SANDBOX ---
brain = CoreColdBrain()
# Simulamos un intento de re-entrada del agente de 205MB
print(brain.audit_environment("hidden_agent_v3", 205.1))
