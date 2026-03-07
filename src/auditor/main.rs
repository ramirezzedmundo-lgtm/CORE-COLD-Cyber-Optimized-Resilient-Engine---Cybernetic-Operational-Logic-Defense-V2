use sqlx::postgres::PgPoolOptions;
use std::env;

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    // 1. Conexión a la base de datos (usando la URL de Docker)
    let database_url = env::var("DATABASE_URL")
        .unwrap_or_else(|_| "postgres://tella_admin:password123@db_postgres:5432/core_cold_db".to_string());

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await?;

    println!("✅ Auditor (Fase 1): Conectado a PostgreSQL con éxito.");

    // 2. Simulación de detección de anomalía (Aquí irá tu lógica de red/ataque)
    let source_ip = "192.168.1.50";
    let alert_msg = "Intento de fuerza bruta detectado (Simulado)";
    let risk_level = 3; // Nivel CRÍTICO

    // 3. Registrar en la tabla de auditoría
    sqlx::query!(
        "INSERT INTO access_logs (source_ip, action_detected, risk_level) VALUES ($1, $2, $3)",
        source_ip,
        alert_msg,
        risk_level
    )
    .execute(&pool)
    .await?;

    println!("⚠️ Alerta registrada: {} desde {}", alert_msg, source_ip);

    Ok(())
}
