# Task2 - Imagy
A simple `Golang` application downloads images from internet and stores them in a local storage.

##Image End-Points:
`Image` has three end-points:
* `GET /api/v1/images` returns list of stored images in Imagy storage. user can select image's name and download it with next API. 
* `GET /api/v1//images/:image_name` by calling this API user can `Download` an image by its name, if there is no image with this name user gets a `404 Not Found` status code.
* `POST /api/v1/images` by calling this API user can `Upload` an image. request's content-type must be `image` otherwise you get an error. also there is max upload limit size in `config/config.json` which configurable. max limit is `2MB`.

## How to Run?
To run the `Imagy` run below command:
```bash
make compose-up
```

>NOTE: make sure `make` and `docker` installed in your system. </br>
> I supposed that you will run `Imagy` in docker containers, so if you want to run it in your local machine, you must do some changes in `config/config.json` file, change `http_address` and `db_address` to the addresses of your preference.
> ```json
>   "http_address": "localhost",
>   "db_address": "localhost:5432"
> ```

## Where is Postman Collections?
After running docker containers you can import `Imagy's Postman Collections` which is located in `./doc/postman/T1-Imagy.postman_collection.json` and send your requests.

## What frameworks that I used?
* `Echo` as Http framework.
* `ENT` as ORM.
* `Docker` for deployment environment.
