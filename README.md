# discord-crypto-gas-price

⚠️This project has been merged with [discord stock tickers](https://github.com/rssnyder/discord-stock-ticker).⚠️

a discord bot to display the current reccomended gas amount for ethereum, binance smart chain, or polygon transactions

![Ethereum Demo Gif](https://s3.cloud.rileysnyder.org/public/assets/ethgas.gif)
![Binance Smart Chain Demo Gif](https://s3.cloud.rileysnyder.org/public/assets/bscgas.gif)
![Polygon Demo Gif](https://s3.cloud.rileysnyder.org/public/assets/maticgas.gif)

## Support
<a href='https://ko-fi.com/rileysnyder' target='_blank'><img height='35' style='border:0px;height:46px;' src='https://az743702.vo.msecnd.net/cdn/kofi3.png?v=0' border='0' alt='Buy Me a Coffee' />
[![Discord Chat](https://img.shields.io/discord/806606291798982678)](https://discord.gg/CQqnCYEtG7)

## Add to your discord server (click image)

[![Ethereum Invite Link](https://s3.cloud.rileysnyder.org/public/assets/ethgas.png)](https://discord.com/api/oauth2/authorize?client_id=833797002684661821&permissions=0&scope=bot)
[![Binance Smart Chain Invite Link](https://s3.cloud.rileysnyder.org/public/assets/bscgas.png)](https://discord.com/api/oauth2/authorize?client_id=856947934452645898&permissions=0&scope=bot)
[![Polygon Invite Link](https://s3.cloud.rileysnyder.org/public/assets/maticgas.png)](https://discord.com/api/oauth2/authorize?client_id=857023179210096674&permissions=0&scope=bot)

## Self Host

Self hosting allows you to update the price in the bot name rather than the activity on the public versions.

### Binary

Download the latest binary from the [release page](https://github.com/rssnyder/discord-crypto-gas-price/releases).

Or build from source:

```
git clone https://github.com/rssnyder/discord-crypto-gas-price.git

cd discord-crypto-gas-price

go build -o discord-crypto-gas-price
```

Create a bot in the discord dev portal, and grab the token. When you add it to your server, be sure to give it "Change Nickname" permissions.

Run the bot with the nessesary settings:

```
  -frequency int
        seconds between gas price cycles (default 5)
  -network string
        ethereum, binance-smart-chain, or polygon
  -setNickname
        wether to set nickname of bot
  -token string
        discord bot token
```

```
./discord-crypto-gas-price network 'ethereum' -token 'xxxxxxxxxxx' -setNickname
```

You can also use the template `discord-crypto-gas-price.service` file to install it as a systemd service.

Fill in your discord bot token in the file, and add any extra arguments you need.

```
cp discord-crypto-gas-price.service /etc/systemd/system/

mkdir -p /etc/discord-crypto-gas-price

cp discord-crypto-gas-price /etc/discord-crypto-gas-price/

systemctl daemon-reload 

systemctl start discord-crypto-gas-price.service 
```
