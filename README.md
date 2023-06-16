# ChoTracker-CC Model-API

This documentation provides an overview of the available routes and endpoints for the Model-API that we use for ChoTracker Application. This API use for handling machine learning models using Flask Framework provides a convenient and efficient way to deploy and interact with machine learning models. Built on the Flask web framework, this API allows you to serve your trained models and make predictions or perform other operations via HTTP requests.

### Prerequisites
- Python installed on your local machine

### Setup Instructions

1. Clone the repository from the [GitHub repository link]([https://github.com/your-repo-link](https://github.com/ChoTracker-C23-PS308/ChoTracker-CC)) model-api branch.
2. Install the required dependencies by running the following command in the project directory:
   ```
   pip install -r requirements.txt
   ```
3. Run the API locally using the following command:
   ```
   python app.py
   ```
   The API should now be running on `http://localhost:5000`.


**Base Url :** `https://baseurl` + `/api/v1`

## Predict Regression Routes
- **Route**: POST /api/v1/predict/regression
- **Description**: Predict user cholesterol by image.
- **File**: Image file in the request body
- **JSON Response**:
    ```json
   {
    "message": "File successfully received and processed",
    "prediction": "{{predict_result}}"
    }
    ```

## Chobot Response
- **Route**: POST /api/v1/predict/chobot
- **Description**: Process user message.
- **JSON Request**:
    ```json
    {
        "text": "message",
    }
    ```
- **JSON Response**:
    ```json
   {
    "message": "{{chobot_response}}}",
    }
    ```
    

