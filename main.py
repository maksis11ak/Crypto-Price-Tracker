import requests
import time

def get_price(crypto_id='bitcoin', currency='usd'):
    url = f"https://api.coingecko.com/api/v3/simple/price?ids={crypto_id}&vs_currencies={currency}"
    response = requests.get(url)
    data = response.json()
    return data[crypto_id][currency]

if __name__ == "__main__":
    crypto = input("Введите ID криптовалюты (bitcoin, ethereum, ...): ").strip().lower()
    while True:
        price = get_price(crypto)
        print(f"{crypto.upper()} price: ${price}")
        time.sleep(10)
