from abc import ABCMeta, abstractmethod
from typing import List, Union
from datetime import datetime

class AbstractEvent:
	eventType : str
	timeStamp : str

	def __init__(self):
		self.eventType = self.__class__.__name__
		self.timeStamp = str(datetime.now())