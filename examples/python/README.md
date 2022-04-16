# Python example

## Pre-requisites

* asdf

## Setup

```
asdf plugin add python
asdf plugin add poetry
asdf install

poetry install
```

## Run

```
NTFSVC_URL=(service URL) NTFSVC_APIKEY=(api key) poetry run python3 main.py "(topic)" "(message)"
```
