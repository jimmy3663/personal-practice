from kafka import KafkaProducer
import json
from typing import List


class Usecase:
    def __init__(self, repository) -> None:
        self._repository = repository
        self._dest = 'shop'
    def list_execute(self):
        obj_list = self._repository.list()
        return obj_list

    def find_by_id(self, id):
        try:
            ele = self._repository.get_by_id(id)
        except Exception as err:
            print(str(err))
        else:
            return ele

    def PostPersist(self, input_dto, count):
        try:
            self._repository.save(input_dto)
        except Exception as err:
            print("use case ERROR!!")
            print(str(err))
            #return WritePostOutputDto(is_success=False, error_message=str(err))
        else:
            #msg = OrderAdapter._default_json_encoder(input_dto)
            input_dto.id = count
            msg = input_dto.__dict__
            producer = KafkaProducer(acks=0, compression_type='gzip', bootstrap_servers=['localhost:9092'], value_serializer=lambda x: json.dumps(x).encode('utf-8'))
            producer.send(self._dest,value=msg)
            print("use case SUCCESS!!")
            return input_dto

    def PrePersist(self, input_dto, count):
        try:
            msg = input_dto.__dict__
            producer = KafkaProducer(acks=0, compression_type='gzip', bootstrap_servers=['localhost:9092'], value_serializer=lambda x: json.dumps(x).encode('utf-8'))
            producer.send(self._dest,value=msg)
        except Exception as err:
            print('use case ERROR on PrePersist!!!!')
            print(str(err))
            return str(err)

        else:
            self._repository.save(input_dto)
            input_dto.id = count
            return input_dto

    def PreRemove(self, id, event_obj):
        try:
            event_obj.id = id
            msg = event_obj.__dict__
            producer = KafkaProducer(acks=0, compression_type='gzip', bootstrap_servers=['localhost:9092'], value_serializer=lambda x: json.dumps(x).encode('utf-8'))
            producer.send(self._dest,value=msg)
            
            self._repository.remove(id)

        except Exception as err:
            print('print err:', err)
            return str(err)
        else:
            
            print("use case SUCCESS!!")
            return event_obj.__dict__