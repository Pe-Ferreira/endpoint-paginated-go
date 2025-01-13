## endpoint-paginated-go

##### Goal
My goal with this project was to simply put into practice a few Golang concepts that I'm learning.

##### The project
endpoint-paginated-go is a very simple webserver that fetches data from an API and serve it in two different routes:
- "/" serves data using and HTML template to display it as a table;
- "/paginated" serves data as a json, but, as you the route states, paginated, so ou can play with "page" and "pageSize" query parameters.

##### How to run
Assuming that you have your Go environment all set, download the code and run ```go run main.go```
