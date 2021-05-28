
from typing import List, Union
from AbstractEvent import AbstractEvent
import json
from datetime import datetime


class OrderPlaced(AbstractEvent):
    id : int
    productId : int
    qty : int
    productName : str

    def __init__(self):
    	super().__init__()
    	self.id = None
    	self.productId = None
    	self.qty = None
    	self.productName = None
		
