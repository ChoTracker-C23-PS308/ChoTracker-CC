# ChoTracker-CC RESTful API Documentation

This documentation provides an overview of the available routes and endpoints for the RESTful API. The API requires authentication using a Firebase token, which should be included in the `Authorization` header of the request.

**Base Url :** `https://baseurl` + `/api/v1`

## User Routes

### Get User by ID

- **Route**: GET /api/v1/users/:id
- **Description**: Get user information by ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
   {
   "data": {
        "id": "user id",
        "name": "user name",
        "email": "user@gmail.com",
        "birth_date": "22 august 2022",
        "gender": "Laki-Laki/Perempuan",
        "image_url": "https://storage.googleapis.com/dev-chotracker-image/users-pict/{{uid}}}}",
        "created_at": "2023-05-19T01:14:44.862103Z",
        "updated_at": "2023-06-08T21:23:50.475154Z"
    },
    "message": "Get user Succesfuly"
    }
    ```

### Create User

- **Route**: POST /api/v1/users
- **Description**: Create a new user.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Request**:
    ```json
    {
        "id": "{{user_id from firebase}}",
        "name": "User name",
        "email": "user@gmail.com",  
        "phone_number": "08225730xxxx",
        "birth_date": "birthdate",
        "gender": "Laki-Laki/Perempuan"
    }
    ```
- **JSON Response**:
    ```json
    {
       "data": "uid",
       "message": "Create User Succesfuly"
    }
    ```

### Update User

- **Route**: PUT /api/v1/users/:id
- **Description**: Update user information by ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Request**:
    ```json
    {
        "name": "Updated Name",
        "email": "updated_email@example.com",
        "phone_number": "updated phone number",
        "birth_date": "update",
        "gender": "Laki-Laki/Perempuan"
    }
    ```
- **JSON Response**:
  ```json
   {
       "data": "uid",
       "message": "Update User Succesfuly"
   }
  ```

### Update User Image

- **Route**: PUT /api/v1/users/:id/image
- **Description**: Update user image by ID. Use multipart/form-data to send the image file.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **File**: Image file in the request body
- **JSON Response**:
    ```json
    {
      "data": "userid",
      "message": "Update User picture Succesfuly"
    }
    ```

## Article Routes

### Get Article by ID

- **Route**: GET /api/v1/articles/:id
- **Description**: Get article by ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
    {
        "data": "{json data}",
        "message": "Message"
    }
    ```

### Get All Articles

- **Route**: GET /api/v1/articles
- **Description**: Get all articles.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
    {
       "data": [
       {
           "ID": "articleid",
           "AuthorID": "userid",
           "JudulArticle": "Chotracker untuk masa depan",
           "IsiArticle": "Ini adalah bagian dari isi article",
           "Author": "author name",
           "ImageUrl": "https://storage.googleapis.com/dev-chotracker-image/users-pict/...",
           "CreatedAt": "2023-06-07T21:58:53.4652Z",
           "UpdatedAt": "2023-06-07T21:58:53.4652Z"
       },
       {
           "ID": "articleid",
           "AuthorID": "userid",
           "JudulArticle": "Chotracker untuk masa depan",
           "IsiArticle": "Ini adalah bagian dari isi article",
           "Author": "author name",
           "ImageUrl": "https://storage.googleapis.com/dev-chotracker-image/users-pict/...",
           "CreatedAt": "2023-06-07T21:58:53.4652Z",
           "UpdatedAt": "2023-06-07T21:58:53.4652Z"
       }],
       "message": "Get Article Succesfuly"
    }
  ```

### Create Article

- **Route**: POST /api/v1/articles/:id
- **Description**: Create a new article.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Request**:
    ```json
    {
        "author_id": "{{user_id}}",
        "judul_artikel": "Article Title",
        "isi_artikel": "Article Content",
        "author": "Author Name",
        "image_url": "image.url/user.jpeg"
    }
    ```
- **JSON Response**:
    ```json
    {
        "data": "c8b8815b-7b1e-4702-8689-41b1f1c7a793",
        "message": "Data successfully created"
    }
    ```

### Update Article

- **Route**: PUT /api/v1/articles/:id
- **Description**: Update article by ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Request**:
    ```json
    {
         "author_id": "{{user_id}}",
        "judul_artikel": "Article Title",
        "isi_artikel": "Article Content",
        "author": "Author Name",
        "image_url": "image.url/user.jpeg"
    }
    ```
- **JSON Response**:
    ```json
    {
        "data": "{json data}",
        "message": "Message"
    }
    ```

### Delete Article

- **Route**: DELETE /api/v1/articles/:id
- **Description**: Delete article by ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
    {
        "data": "article id",
        "message": "Data successfully updated"
    }
    ```

## History Routes

### Create History

- **Route**: POST /api/v1/history/:id
- **Description**: Create a new history record.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Request**:
    ```json
    {
        "uid": "{{user_id}}",
        "total_kolestrol": "xxx.xxx",
        "tingkat": "high/normal",
        "image_url": "user.history.url/image.jpeg"
    }
    ```
- **JSON Response**:
    ```json
    {
        "data": "history id",
        "message": "Data successfully created"
    }
    ```

### Update History Image

- **Route**: POST /api/v1/history/:id/image
- **Description**: Update history image by ID. Use multipart/form-data to send the image file.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **File**: Image file in the request body
- **JSON Response**:
    ```json
    {
        "data": "https://storage.googleapis.com/dev-chotracker-image/history-pict/...",
        "message": "Image Upload Succesfuly"
    }
    ```

### Delete History

- **Route**: DELETE /api/v1/history/:uid/:id
- **Description**: Delete history record by UID and ID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
    {
        "data": "{{hid}}",
        "message": "History successfully deleted"
    }
    ```

### Get User History

- **Route**: GET /api/v1/history/:uid
- **Description**: Get user history by UID.
- **Headers**:
    - Authorization: Bearer <Firebase Token>
- **JSON Response**:
    ```json
    {
      "data": [
          {
            "id": "history id",
            "uid": "user id",
            "total_olestrol": "xxx.xxxx",
            "tingkat": "high/normal",
            "image_url": "https://storage.googleapis.com/dev-chotracker-image/users-pict/...",
            "created_at": "2023-05-24T16:55:46.409322Z",
            "updated_at": "2023-05-24T16:55:46.409322Z"
          }, 
          {
            "id": "history id",
            "uid": "user id",
            "total_olestrol": "xxx.xxxx",
            "tingkat": "high/normal",
            "image_url": "https://storage.googleapis.com/dev-chotracker-image/users-pict/...",
            "created_at": "2023-05-24T16:55:46.409322Z",
            "updated_at": "2023-05-24T16:55:46.409322Z"
          }, 
  {
            "id": "history id",
            "uid": "user id",
            "total_olestrol": "xxx.xxxx",
            "tingkat": "high/normal",
            "image_url": "https://storage.googleapis.com/dev-chotracker-image/users-pict/...",
            "created_at": "2023-05-24T16:55:46.409322Z",
            "updated_at": "2023-05-24T16:55:46.409322Z"
          }]
      ,"message": "Get Histoies Data successfully"
    }
    ```
---
