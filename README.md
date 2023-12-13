### GoCrypto with Terminal

Simple Terminal UI for GoCrypto. This project uses [CoinMarketCap](https://coinmarketcap.com/) API.

---

## Installation

Install go-crypto with the command below;

```bash

go install github.com/root27/go-crypto@latest

```

Go will automatically install it in your $GOPATH/bin directory which should be in your $PATH.

---

## Troubleshooting

If you have an issue with PATH variable or you don't set GOPATH or GOBIN variable. You can use these steps shown below;

Step 1:

    - Edit ~/.zshrc or ~/.bashrc

Step 2:

    - Paste these lines;
        export GOPATH=$HOME/go
        export GOBIN=$GOPATH/bin
        export PATH=${PATH}:$GOBIN

Step 3:

    - source ~/.zshrc or ~/.bashrc

---

## Screenshot of the project

![Screenshot](./assets/getAll.png)

---

## License


[MIT](./LICENSE)

---

## Author

Oguzhan Dogan




