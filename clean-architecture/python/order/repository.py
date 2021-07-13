from abc import ABCMeta, abstractmethod
from typing import List, Union

class Repository(metaclass=ABCMeta):
	
	@abstractmethod
	def list(self):
		pass

	@abstractmethod
	def save(self, event_object):
		pass

	@abstractmethod
	def get_by_id(self, id:int):
		pass

	@abstractmethod
	def update(self, value, orm_object):
		pass