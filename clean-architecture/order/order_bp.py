
from flask import Blueprint

from flask import Flask, abort, jsonify, request
from orderPlacedAdapter import OrderPlacedAdapter
from orderCancelledAdapter import OrderCancelledAdapter
from UseCase import Usecase
import json
from orderRepository import OrderRepository

import Order_hateoas

bp = Blueprint('main', __name__, url_prefix='/')

repository = OrderRepository()
count = Order_hateoas.CNT(repository)

@bp.route("/orders", methods=["POST"])
def orderPlaced_post():
    count.cnt = count.cnt + 1

    # convert request_body to input_dto
    event = OrderPlacedAdapter.request_body_to_write_order_input_dto(request.json)
    
    # execute use_case
    use_case = Usecase(repository)
    msg = use_case.PostPersist(event, count.cnt)
    response = Order_hateoas.POST_response(request.base_url, msg)

    return response
@bp.route("/orders/<int:id>", methods=["DELETE"])
def orderCancelled_post(id: int):
    
    use_case = Usecase(repository)
    remove_ele = use_case.find_by_id(id)[0]
    event_obj = OrderCancelledAdapter.request_body_to_write_order_input_dto(remove_ele.__dict__)
    
    msg = use_case.PreRemove(id,event_obj)

    return msg
@bp.route("/orders", methods=["GET"])
def get_order():
    use_case = Usecase(repository)
    order_list = use_case.list_execute()
    response = Order_hateoas.GET_list_response(request.base_url,order_list)
    
    return response

@bp.route("/orders/<int:id>", methods=["GET"])
def get_order_by_id(id: int):
    use_case = Usecase(repository)
    ele = use_case.find_by_id(id)
    response = Order_hateoas.GET_id_response(request.base_url, ele[0])
   
    return response

@bp.errorhandler(404)
def resource_not_found(e):
    return jsonify(error=str(e)), 404
