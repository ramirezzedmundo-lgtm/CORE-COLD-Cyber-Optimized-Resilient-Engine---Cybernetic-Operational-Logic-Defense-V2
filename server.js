const axios = require('axios'); // Instala con: npm install axios

async function reportToMonitor(ip) {
    try {
        await axios.get(`http://localhost:8080/report?ip=${ip}`);
    } catch (error) {
        console.error("Error conectando con el monitor de Go");
    }
}

// En tu ruta de login, si la clave es incorrecta:
// reportToMonitor(req.ip);
