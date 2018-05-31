# Neptr: Analytics Bot!

**Neptr** is a Discord analytics bot which delivers detailed analytics about a certain guild/server of your choosing. Neptr's primary goal is to help anyone explore and study their server's data, by storing everything in an SQLite database, which can later be easily accessed and modified by many popular programming languages for any purpose.

So far this is not a *self-hosted bot*, but feel free to host it if you wish to do so 😉.

## Setting up

To use the bot you'll need the create two files in the repo: `config.json` and a SQLite database (thjis one can be stored wherever you want).

The `config.json` will host 3 key components of the bot, the **Token**, the **Prefix** and the **DB location**, like so:
```
{
  "Token": "<Your token>",
  "BotPrefix": "§",
  "DB": "./test_db.db"
}
```

## How does it work?

The bot and commands are still a bit messy and will likely be changed in a future commit. Right now, **the** way to use the bot is by writing the following commands:
```
§cm // Collects every member
§cu // Collects every user
§cc // Collects all channels
§cr // Collects all roles
§ce // Collects every emoji
§ct <channel id> // Collects the latest 100 messages in a channel
```

**NOTE:** The bot isn't smart enough yet now to know if the DB has been filled already or not, so everytime you run any of the commands above it will always insert the info as new entries.

## Contributing

The code might be a bit messy right now and there's still a lot of things to do. Any commits would be greatly appreciated 😃.

**NOTE:** I might take a while to merge any commits since I'm near my finals.