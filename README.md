# Hexa Progress

Hexa Progress is a discord bot for calculating the progress of a Maplestory character's hexa matrix.

## Features
- Calculate the progress of a Maplestory character's hexa matrix.
- Find videos of where Korean Maplestory players are training with a specific class.
- Find videos of how Korean Maplestory players are bossing with a specific class.

## Screenshots
![Hexa Progress](/screenshots/hexaprogress-photo-collage.png)

## Installation
### Environment Variables
HEXAPROGRESS_DISCORD_TOKEN - Discord bot token, obtained from the discord developer portal.  
HEXAPROGRESS_GUILD_ID (optional) - Discord guild id, obtained from the discord client. This token is only needed if you want immediate command updates.
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