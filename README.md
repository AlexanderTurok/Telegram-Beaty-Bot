# beaty-telegram-bot
Telegram Bot for holding a Beaty Contest.

## Tools used for creation
-[Go Telegram SDK](https://github.com/go-telegram-bot-api/telegram-bot-api) for using Telegram API.<br />
-Postgres for storing users data.<br />
-Redis for Caching.<br />

## To start
1. Create .env file
```
  API_KEY = Telegram API key 
  DB_PASSWORD = Password to your SQL database 
  REDIS_PASSWORD = Redis Password 
 ```
2. Edit ```Config.yml``` 
3. Execute command: ```go run cmd/main.go```
