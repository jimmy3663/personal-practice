from orderRepository import OrderRepository
from Order import Order

import json
from collections import OrderedDict


class CNT:
	def __init__(self, repository):
		self.cnt = len(repository.list())


def GET_id_response(base_url, orm_object:Order):
	response=OrderedDict()
    
    
	response['_links'] = {
		'order' :{ 					
			'herf':base_url
		},
		'self':{
			'herf':base_url
		}
	}
	response['productId'] = orm_object.productId
	response['qty'] = orm_object.qty
	response['productName'] = orm_object.productName
	return response

def POST_response(base_url, event_obj):
	response = OrderedDict()
	url = base_url + '/'+ str(event_obj.id)
	response['_links'] = {
		'order': {
			'herf': url
		},
		'self':{
			'herf': url
		}
	}
	response['productId'] = event_obj.productId
	response['qty'] = event_obj.qty
	response['productName'] = event_obj.productName
	return response

def GET_list_response(base_url,Order_list):
	totalElements = len(Order_list)
	size = 20
	number = totalElements//size
	response = OrderedDict()
	obj_list = []
	for ele in Order_list:
		obj_dict = OrderedDict()
		url = base_url + '/'+str(ele.id)
		obj_dict['_list'] = {
			"order": {
				"herf":url
			},
			"self":{
				"herf":url
			}
		}
		obj_dict['productId'] = ele.productId
		obj_dict['qty'] = ele.qty
		obj_dict['productName'] = ele.productName
		obj_list.append(obj_dict)

	response['_embedded'] = obj_list
	response['_links'] = {
		'profile' : {
			'herf': base_url+"profile/orders"
		},
		'self':{
			'href': base_url+"{?page,size,sort}",
			'templated': True
		}
	}
	response['page']={
		'number': str(number),
		'size': str(size),
		'totalElements':str(totalElements),
		'totalPages': str(number+1)
	}
	return response



