import os
import uuid
from typing import Final

import boto3
from botocore.config import Config

EVIDENTLY_ENDPOINT_URL: Final[str] = os.environ.get("EVIDENTLY_ENDPOINT_URL", "")

# create Evidently client
if len(EVIDENTLY_ENDPOINT_URL) == 0:
    client = boto3.client("evidently")
else:
    client = boto3.client(
        "evidently",
        endpoint_url=EVIDENTLY_ENDPOINT_URL,
        config=Config(inject_host_prefix=False),
    )

# evaluate feature
PROJECT: Final[str] = "food"
FEATURE: Final[str] = "sushi"

entity_id: str = uuid.uuid4().hex

try:
    out = client.evaluate_feature(
        project=PROJECT,
        feature=FEATURE,
        entityId=entity_id,
    )

    print(out.get("reason"))
    print(out.get("variation"))
    print(out.get("value"))

except Exception as e:
    print(e)
