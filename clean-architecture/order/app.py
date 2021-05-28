from flask import Flask, abort, jsonify, request
from kafka import KafkaConsumer

import order_bp
app = Flask(__name__)

app.register_blueprint(order_bp.bp)
if __name__ == "__main__":
    app.run()
