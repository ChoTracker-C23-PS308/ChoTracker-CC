from PIL import Image
import numpy as np
import cv2
from skimage.feature import local_binary_pattern
import joblib

class RegressionPredictor:
    def __init__(self, model_path):
        self.model = self.load_model(model_path)
    
    def load_model(self, model_path):
        model = joblib.load(model_path)
        return model

    def preprocess_image(self, image):
        gray_image = image.convert('L')
        np_image = np.array(gray_image)
        radius = 3
        n_points = 8 * radius
        lbp = local_binary_pattern(np_image, n_points, radius, method='uniform')
        histogram, _ = np.histogram(lbp.ravel(), bins=np.arange(0, n_points + 3), range=(0, n_points + 2))
        feature = np.reshape(histogram, (1, -1))
        return feature

    def predict(self, image):
        feature = self.preprocess_image(image)
        prediction = self.model.predict(feature)
        return prediction[0] 


# Ini contoh kode untuk error ValueError: X has 26 features, but LinearRegression is expecting 50 features as input.

# class RegressionPredictor:
#     def __init__(self, model_path):
#         self.model = self.load_model(model_path)
    
#     def load_model(self, model_path):
#         model = joblib.load(model_path)
#         return model

#     def preprocess_image(self, image):
#         gray_image = image.convert('L')
#         np_image = np.array(gray_image)
#         radius = 3
#         n_points = 8 * radius
#         lbp = local_binary_pattern(np_image, n_points, radius, method='uniform')
#         histogram, _ = np.histogram(lbp.ravel(), bins=np.arange(0, n_points + 3), range=(0, n_points + 2))
#         feature = np.reshape(histogram, (1, -1))
        
#         # Adjust the number of features to match the expected input
#         expected_features = 50
#         if feature.shape[1] < expected_features:
#             padding = np.zeros((1, expected_features - feature.shape[1]))
#             feature = np.concatenate((feature, padding), axis=1)
#         return feature

#     def predict(self, image):
#         feature = self.preprocess_image(image)
#         prediction = self.model.predict(feature)
#         return prediction
