const nodemailer = require('nodemailer');

// 1. Configuración del "Mensajero" (Transporter)
// Puedes usar Gmail o Outlook cambiando el host
const transporter = nodemailer.createTransport({
  service: 'gmail', // O 'outlook'
  auth: {
    user: 'tu-correo@gmail.com', // Aquí irá tu cuenta
    pass: 'tu-contraseña-de-aplicacion' // No es tu clave normal, es una de "App Password"
  }
});

// 2. Función de Auditoría de Emergencia
async function enviarAlertaCibernetica(ip, riesgo, detalle) {
  const mailOptions = {
    from: '"🛡️ CORE-FRÍO V2" <tu-correo@gmail.com>',
    to: 'tu-correo-destino@outlook.com',
    subject: `🚨 ALERTA DE SEGURIDAD - NIVEL ${riesgo}`,
    text: `Reporte de Auditoría Digital:\n\nOrigen: ${ip}\nAcción: ${detalle}\nEstado: Fase 2 (Bridge) ha rotado los tokens.`,
    html: `<b><h1>CORE-FRÍO V2: Alerta de Intrusión</h1></b><p>Se ha detectado una anomalía crítica desde la IP: <b>${ip}</b>.</p>`
  };

  try {
    let info = await transporter.sendMail(mailOptions);
    console.log('✅ Correo de alerta enviado: ' + info.response);
  } catch (error) {
    console.error('❌ Error al enviar alerta: ', error);
  }
}

// Simulación de escucha de eventos (En la Fase 3 esto se activa por la IA)
// enviarAlertaCibernetica('192.168.1.50', 'CRÍTICO', 'Intento de fuerza bruta');
