# Hexa Progress

Hexa Progress is a discord bot for calculating the progress of a Maplestory character's hexa matrix.

## Screenshots
![Hexa Progress](/screenshots/6-27-24.png)

## Installation
### Environement Variables
HEXAPROGRESS_DISCORD_TOKEN - Discord bot token, obtained from the discord developer portal.  
HEXAPROGRESS_GUILD_ID - Discord guild id, obtained from the discord client.
### Docker
```bash
docker build -t 'hexa-progress' .
docker run -it --rm -e HEXAPROGRESS_DISCORD_TOKEN -e HEXAPROGRESS_GUILD_ID hexa-progress
```
### Golang
```bash
go build -o hexaprogress ./cmd/hexaprogress
./hexaprogress
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)