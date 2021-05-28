from sqlalchemy.ext.declarative import declarative_base 
from sqlalchemy import Column, String, Integer
from datetime import datetime

Base = declarative_base()

class Order(Base):
	__tablename__ = 'Order_table'
	id = Column(Integer, primary_key=True)
	productId = Column(Integer)
	qty = Column(Integer)
	productName = Column(String(50))
