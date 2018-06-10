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

```javascript
§start // Collects everything, including the latest 100 messages from each channel
```

or

```javascript
§cm // Collects every member
§cu // Collects every user
§cc // Collects all channels
§cr // Collects all roles
§ce // Collects every emoji
§ct <channel id> // Collects the latest 100 messages in a channel
```

**NOTE:** The bot isn't smart enough yet now to know if the DB has been filled already or not, so everytime you run any of the commands above it will always insert the info as new entries.

## Emojis
Emojis are a bit tricky to count, since some emojis actually consist of two or
more emojis combined.

For example the 🏳️‍🌈 emoji is actually a 🏳 and a 🌈 combined together, so the bot will count it as one character.

### Characters table
So far the `characters` table counts emojis like this: 🏳️‍🌈 and 👨‍🏫 as one character.
However this means that a string like this: 🔥🔥🔥🔥 will also be counted as one emoji
which is not the intended outcome. I'm still trying to solve this issue by probably
treating emojis the same way as they are treated in the `emojis` table (see below).

### Emojis table
In the `emojis` table an emoji like this: 🏳️‍🌈 is stored by each of it's parts, like so:
🏳 🌈. This way each emoji will always be counted separately, even if it's part of
another one. (I'll probably use the same approach in the `characters` table in the future)

## Contributing

The code might be a bit messy right now and there's still a lot of things to do. Any commits would be greatly appreciated 😃.

**NOTE:** I might take a while to merge any commits since I'm near my finals.