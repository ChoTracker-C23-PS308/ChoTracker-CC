from flask import Flask, request, jsonify
from PIL import Image
import numpy as np
import tensorflow as tf
from predictcnn import Predict
import cv2
from predictregression import RegressionPredictor
import os

app = Flask(__name__)

model_path_cnn = "model\model_weights.h5" # Size model Besar, Download model di trello

current_dir = os.path.dirname(os.path.abspath(__file__))
relative_path = "model/regression_model.h5"
model_path_regression = os.path.join(current_dir, relative_path)

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

        response = {
            'message': 'File successfully received and processed',
            'prediction': float(prediction_result)
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
