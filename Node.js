const axios = require('axios');

async function getPrice(cryptoId = 'bitcoin', currency = 'usd') {
    const url = `https://api.coingecko.com/api/v3/simple/price?ids=${cryptoId}&vs_currencies=${currency}`;
    const response = await axios.get(url);
    return response.data[cryptoId][currency];
}

(async () => {
    const crypto = process.argv[2] || 'bitcoin';
    setInterval(async () => {
        const price = await getPrice(crypto);
        console.log(`${crypto.toUpperCase()} price: $${price}`);
    }, 10000);
})();
