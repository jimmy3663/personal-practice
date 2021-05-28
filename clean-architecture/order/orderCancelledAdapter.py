import json
from dataclasses import asdict, is_dataclass
from datetime import datetime
from typing import Any, List
from OrderCancelled import OrderCancelled
import json

class OrderCancelledAdapter:

    @classmethod
    def request_body_to_write_order_input_dto(cls, request_body: dict):
        #return ProductChangedDto(**request_body)
        orderCancelled_event = OrderCancelled()
        orderCancelled_event.productId = request_body['productId']
        orderCancelled_event.qty = request_body['qty']
        orderCancelled_event.productName = request_body['productName']

        return orderCancelled_event
