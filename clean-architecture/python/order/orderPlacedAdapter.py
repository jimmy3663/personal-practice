import json
from dataclasses import asdict, is_dataclass
from datetime import datetime
from typing import Any, List
from OrderPlaced import OrderPlaced
import json

class OrderPlacedAdapter:

    @classmethod
    def request_body_to_write_order_input_dto(cls, request_body: dict):
        #return ProductChangedDto(**request_body)
        orderPlaced_event = OrderPlaced()
        orderPlaced_event.productId = request_body['productId']
        orderPlaced_event.qty = request_body['qty']
        orderPlaced_event.productName = request_body['productName']

        return orderPlaced_event
