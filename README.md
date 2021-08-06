![Slack logo](assets/1024px-Slack_Technologies_Logo.svg.png)

### Slacker CLI is a Slack API (partial) implementation write in golang.

---

## 💻 OS supported 

 - Linux based distros
 - MacOS
 - Windows WSL

## 📦 Requirement

- Go 1.13 or more

## 💾 Installation

To install this project, just make:

```bash
go get && go build slacker.go -o slacker
```

If you are the laziest person of this world, you can just make:

```bash
make install
```

This command will import all dependancies in a **vendor** folder on root project, build this application (CLI), and finally add an environment variable to use this CLI through any terminal just make `slacker`.

> ⓘ Calm down, this command don't write in any files, just add env using `export` command. So if you restart your PC/Mac (for purist), enrivonment variable will be remove. You can make `make set-env` to set path variable.

---

## Use

### Set slack environment

If you use Slack API, is because you already have a Slack workspace, and generate a token with defined permission.
Then, make:

```bash
slacker set [token-app]
```

And now, your cli is ready to use.

### 💬 Send message

There many way to use direct messages though these CLI.
Common way is, send a message to known email as:

```bash
slacker sendto [email] [message]
```

## 📌 Reference 
- [Slack API](api.slack.com) documentation
- [Slack Go Library](http://github.com/slack-go/slack) is unofficial library who covering most important API use case.