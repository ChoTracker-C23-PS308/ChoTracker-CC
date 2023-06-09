from flask import Flask, request, jsonify
from PIL import Image
import numpy as np
import tensorflow as tf
from predictcnn import Predict
from predictregression import RegressionPredictor
from predictchobot import predict_intent
import cv2
import tflearn
import os
import os

app = Flask(__name__)

current_dir = os.path.abspath(__file__)  
parent_dir = os.path.dirname(os.path.dirname(current_dir))  
target_dir = os.path.join(parent_dir, "Chotracker-CC-ModelAPI/model")
target_dir_chobot = os.path.join(parent_dir, "Chotracker-CC-ModelAPI/model_chobot")

# Model path
model_path_regression = os.path.join(target_dir, "regression_model.h5") 
# model_path_cnn = os.path.join(target_dir, "model_weights.h5")  # Size model Besar, Download model di trello

@app.route("/api/v1/predict/chobot", methods=['POST'])
def predict_chatbot():
    if request.method == 'POST':
        if 'text' not in request.json:
            return jsonify({'message': 'Text input not provided'}), 400
        text = request.json['text']
        response = predict_intent(text)
        return jsonify({'message': response}), 200

@app.route("/api/v1/predict/regression", methods=['POST'])
def predict_regression_image():
    if request.method == 'POST':
        if 'file' not in request.files:
            return jsonify({'message': 'There is no file sended'}), 400

        file = request.files['file']

        if not allowed_file(file.filename):
            return jsonify({'message': 'File is not image'}), 400

        image = Image.open(file.stream)

        predictor = RegressionPredictor(model_path_regression)
        prediction_result = predictor.predict(image)
        rounded_value = "{:.2f}".format(prediction_result)
        response = {
            'message': 'File successfully received and processed',
            'prediction': rounded_value
        }
        return jsonify(response), 200

@app.route("/api/v1/predict/cnn", methods=['POST'])
def predict_image():
    if request.method == 'POST':
        if 'file' not in request.files:
            return jsonify({'message': 'There is no file sended'}), 400
        
        file = request.files['file']

        if not allowed_file(file.filename):
            return jsonify({'message': 'File is not image'}), 400

        image = Image.open(file.stream)
        
        predictor = Predict(image, model_path=model_path_cnn)
        prediction_result = predictor.predict()

        response = {
            'message': 'File successfully received and processed',
            'prediction': prediction_result['predicted_class'],
            'probability': prediction_result['probability'],
        }
        return jsonify(**response), 200

def allowed_file(filename):
    ALLOWED_EXTENSIONS = {'jpg', 'jpeg', 'png'}
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

if __name__ == '__main__':
    app.run()
