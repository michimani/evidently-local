import os
import sys
import uuid
from typing import Final

import boto3
from botocore.config import Config

EVIDENTLY_ENDPOINT_URL: Final[str] = os.environ.get("EVIDENTLY_ENDPOINT_URL", "")


def create_client():
    """Create a CloudWatch Evidently client."""
    if len(EVIDENTLY_ENDPOINT_URL) == 0:
        return boto3.client("evidently")

    return boto3.client(
        "evidently",
        endpoint_url=EVIDENTLY_ENDPOINT_URL,
        config=Config(inject_host_prefix=False),
    )


def evaluate_feature(client, project, feature, entity_id):
    try:
        out = client.evaluate_feature(
            project=project,
            feature=feature,
            entityId=entity_id,
        )

        print(out.get("reason"))
        print(out.get("variation"))
        print(out.get("value"))

    except Exception as e:
        print(e)


if __name__ == "__main__":
    PROJECT: Final[str] = "food"
    FEATURE: Final[str] = "sushi"
    entity_id: str = uuid.uuid4().hex

    args = sys.argv
    if len(args) > 1:
        entity_id = args[1]

    client = create_client()
    evaluate_feature(client, PROJECT, FEATURE, entity_id)
