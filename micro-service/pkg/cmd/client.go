package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sachinpuranik/samples/micro-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type ToDo struct {
	Id          int64
	Title       string
	Description string
	Reminder    time.Time
}

func (client *Client) ConnectGRPCServer(address string) {
	// Set up a connection to the server.
	cred := insecure.NewCredentials()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client.conn = conn

	client.cl = v1.NewToDoServiceClient(conn)
}

func (client *Client) DisconnectGRPCServer() {
	client.conn.Close()
}

func (client *Client) helloHandler(c *gin.Context) {
	w := c.Writer
	w.Write([]byte("hello"))
	w.WriteHeader(http.StatusOK)
}

func (client *Client) createTodo(c *gin.Context) {
	var p ToDo

	w := c.Writer
	r := c.Request

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)

	// Call Create
	bufReq := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       p.Title,
			Description: p.Description,
			Reminder:    reminder,
		},
	}
	res, err := client.cl.Create(ctx, &bufReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Create result: <%+v>\n\n", res)

	p.Id = res.Id

	jsonData, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

func (client *Client) readTodo(c *gin.Context) {
	var p ToDo

	w := c.Writer
	r := c.Request

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Read
	bufReq := v1.ReadRequest{
		Api: apiVersion,
		Id:  p.Id,
	}
	res, err := client.cl.Read(ctx, &bufReq)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res)

	p.Id = res.ToDo.Id
	p.Title = res.ToDo.Title
	p.Description = res.ToDo.Description
	p.Reminder = res.ToDo.Reminder.AsTime()

	jsonData, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

func (client *Client) updateTodo(c *gin.Context) {
	var p ToDo

	w := c.Writer
	r := c.Request

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	// Update
	bufReq := v1.UpdateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Description,
			Reminder:    reminder,
		},
	}
	res, err := client.cl.Update(ctx, &bufReq)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res)

	jsonData, err := json.Marshal(res.Updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

func (client *Client) readAllTodo(c *gin.Context) {
	w := c.Writer

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	// defer cancel()
	ctx := context.Background()
	// Call ReadAll
	bufReq := v1.ReadAllRequest{
		Api: apiVersion,
	}
	res, err := client.cl.ReadAll(ctx, &bufReq)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res)

	resp := []ToDo{}
	for _, item := range res.ToDos {
		e := ToDo{}
		e.Id = item.Id
		e.Title = item.Title
		e.Description = item.Description
		e.Reminder = item.Reminder.AsTime()
		resp = append(resp, e)
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

func (client *Client) deleteTodo(c *gin.Context) {
	var p ToDo

	w := c.Writer
	r := c.Request

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Delete
	bufReq := v1.DeleteRequest{
		Api: apiVersion,
		Id:  p.Id,
	}
	res, err := client.cl.Delete(ctx, &bufReq)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res)

	jsonData, err := json.Marshal(res.Deleted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

type Client struct {
	cl   v1.ToDoServiceClient
	conn *grpc.ClientConn
}

type ClientConnector interface {
	RunClientService(cfg HttpServerConfig)
	ConnectGRPCServer(address string)
}

func NewClient() ClientConnector {
	return &Client{}
}

type HttpServerConfig struct {
	Port int
}

func (client *Client) RunClientService(cfg HttpServerConfig) {
	root := gin.New()

	var grp gin.IRoutes = root.Group("/")

	grp.POST("/create", client.createTodo)
	grp.GET("/read", client.readTodo)
	grp.GET("/read-all", client.readAllTodo)
	grp.PUT("/update", client.updateTodo)
	grp.DELETE("/delete", client.deleteTodo)

	// basicAuth := gin.BasicAuth(gin.Accounts{
	// 	"sachin": "test",
	// })
	// grp.Use(basicAuth)

	grp.GET("/", client.helloHandler)

	go http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", cfg.Port), root)
}
