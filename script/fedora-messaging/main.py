import json
from dataclasses import dataclass
from typing import Dict

from fedora_messaging import api, config

from kafka import KafkaProducer


@dataclass
class Event:
    Key: str
    Header: Dict[str, str]
    Body: object


# 创建生产者
producer = KafkaProducer(bootstrap_servers='7.250.74.126:9092', api_version=(2, 8, 1))

# Kafka 主题
topic = 'eur_build_raw'

config.conf.setup_logging()


def json_marshal(obj):
    # 将对象编码为 JSON 格式的字符串
    json_string = json.dumps(
        obj,
        default=lambda o: o.__dict__,
        sort_keys=True,
        indent=4)
    # 将字符串转换为字节数组
    json_bytes = json_string.encode('utf-8')
    return json_bytes


def print_eur_message(message):
    print("******** Received eur event: ********\n")
    producer.send(topic, json_marshal(message))

api.consume(lambda message: print_eur_message(message))
