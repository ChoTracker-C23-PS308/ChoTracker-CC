# ChoTracker-CC Model-API

This documentation provides an overview of the available routes and endpoints for the Model-API that we use for ChoTracker Application.

**Base Url :** `https://baseurl` + `/api/v1`

## Predict Regression Routes
- **Route**: POST /api/v1/predict/regression
- **Description**: Predict user cholesterol by image.
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
- **JSON Response**:
    ```json
   {
    "message": "{{chobot_response}}}",
    }
    ```