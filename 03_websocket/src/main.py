#!/usr/bin/python

from flask import Flask, render_template
from flask_socketio import SocketIO
from flask_restful import Resource, Api


app = Flask(__name__)
api = Api(app)



socketio = SocketIO(app, cors_allowed_origins="*")

@app.route('/')
def index():
    return render_template('index.html')


class InternalAPI(Resource):
    def get(self):
        return {'message': 'You\'ve accessed internal API!'}


class Flag(Resource):
    def get(self):
        return {'flag': 'In 50VI37 rUS5I4 vODK@ DRiNKs YOu!!!'}

@socketio.on('my event')
def handle_my_custom_event(json):
    print('received json: ' + str(json))
    pass


if __name__ == '__main__':
    
    api.add_resource(Flag, '/flag')
    api.add_resource(InternalAPI, '/api/internal')

    print("Server running on port 8000")
    socketio.run(app, host='0.0.0.0', port=8000, debug=True) #, host='localhost'
