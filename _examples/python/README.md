by Python (AWS SDK for Python (boto3))
===

This is an example of using AWS SDK for Python (boto3).

## Prepare

```bash
python3 -m venv .venv \
&& source .venv/bin/activate \
&& pip install pip --upgrade \
&& pip install -r requirements.lock
```

## Run

### Request to real resource

```bash
python main.py
```

### Request to evidently local

```bash
EVIDENTLY_ENDPOINT_URL='http://localhost:2306' python main.py
```