# API Documentation

## 1. Create a Tag

- **Endpoint:** `POST /tags`

- **Description:** Create a new tag.

- **Request Body:**
  ```json
  {
    "name": "Wembley Stadium"
  }
- **Response:** `201 Created`


## 2. List All Tags

- **Endpoint:** `GET /tags`

- **Description:** List all current tags.

- **Response:** `200 OK`
    ```json
  [
	{ 
		"id": "1234",
		"name": "Wembley Stadium" 
	} 
  ]

## 3. Create a Media

- **Endpoint:** `POST /media`

- **Description:** Create a new media with associated tags

- **Request Body *Form-Data*:**
  ```json
  {
    "name": "TF1",
    "tags": ["1", "2"],
    "file": binary file
  }
- **Response:** `201 Created`

## 4. Retrieve Media

- **Endpoint:** `GET /media?tag=1243`

- **Description:** Retrieve media by tag, if no tag provided all media are returned 

- **Response:** `200 OK`
    ```json
  [{ 
        "id": "c906cbbf-1a25-4a99-b223-34bcf6e3b8a7",
        "name": "super nice picture",
        "tags": [ "Zinedine Zidane", "Real Madrid", "Champions League" ],
        "fileUrl": "https://some_url.com/file.jpg"
    }]