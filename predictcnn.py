from PIL import Image
import numpy as np
import tensorflow as tf

class Predict:
    def __init__(self, image, model_path):
        self.image = image
        self.model_path = model_path

    def load_model(self):
        model = tf.keras.models.load_model(self.model_path)
        return model
    
    def preprocess_image(self):
        image = self.image.resize((48, 48))
        image = image.convert("RGB") 
        image = np.array(image) / 255.0
        image = np.expand_dims(image, axis=0)
        return image

    def predict(self):
        model = self.load_model()
        image = self.preprocess_image()
        predictions = model.predict(image)
        predicted_class = np.argmax(predictions)  
        probability = predictions[0][predicted_class]
        result = {
            'predicted_class': float(predicted_class),
            'probability': float(probability)
        }
        return result
