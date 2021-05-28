from typing import List

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from repository import Repository
from Order import Order


class OrderRepository(Repository):
	def __init__(self):
		self.engine = create_engine('sqlite:///Order_table.db', echo=True)
		
		Order.__table__.create(bind=self.engine, checkfirst=True)


	def _create_order_object(self , results: List[Order]) -> List[Order]:
		return [
			Order(
                id=q.id, 
                productId=q.productId, 
                qty=q.qty, 
                productName=q.productName, 
                )
			for q in results
		]


	def list(self):
		DBsession = sessionmaker(bind=self.engine)

		session = DBsession()
		query = session.query(Order)

		return self._create_order_object(query.all())

	def save(self, event_object):
		DBsession = sessionmaker(bind=self.engine)

		session = DBsession()
		

		entity_obj = Order(
                productId = event_object.productId,
                qty = event_object.qty,
                productName = event_object.productName,
			)
		
		
		session.add(entity_obj)
		session.commit()

	def get_by_id(self, id:int):
		DBsession = sessionmaker(bind=self.engine)
		session = DBsession()
		
		try:
			query = session.query(Order).filter_by(id=id).one()
		except Exception as err:
			print(err)
		return self._create_order_object([query])

	def update(self, value, orm_object):
		DBsession = sessionmake(bind=self.engine)
		session = DBsession()

		orm_object.status = value

		session.commit()

	def remove(self, id):
		DBsession = sessionmaker(bind=self.engine)
		session = DBsession()

		try:
			query = session.query(Order).filter_by(id=id).one()
			session.delete(query)
		except Exception as err:
			print('error:',err)

		session.commit()


class NotExistedDataError(Exception):
    pass		


