package main

import (
    kafka "github.com/segmentio/kafka-go"
    "os"
    "encoding/json"
    "fmt"
    "strings"
    "context"
    //"github.com/aws/aws-sdk-go-v2/service/rds"
)

type request struct {
    Requestor string
    Operator string
    Param1 int
    Param2 int
}

func main () {
    raw_topic, exists := os.LookupEnv("RAW_TOPIC")
    kafka_brokers := strings.Split(os.Getenv("BROKER_LIST"), ",")
    consumer_group := os.Getenv("CONSUMER_GROUP")

    if ! exists {
        fmt.Printf("REQUIRED env variables : \n\t%s\n\t%s\n\t%s\n\t%s\n",
	    "RAW_TOPIC", "BROKER_LIST", "CONSUMER_GROUP")
	os.Exit(0)
    }

    reader := getKafkaReader(kafka_brokers, raw_topic, consumer_group)
    defer reader.Close()

    fmt.Println("start consuming")

    //operators := []string{"add", "subtract", "multiply", "divide"}
    var req_batch []request
    batch_size := 50
    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
          fmt.Printf("error reading message")
        }
        req := request{}
        json.Unmarshal(msg.Value, &req)
        fmt.Printf("received %s\n", req)

	if len(req_batch) > batch_size {
	    fmt.Printf("insert into xyz")
            req_batch = []request{req}
        } else {
            req_batch = append(req_batch, req)
	}
    }
    //err = writer.WriteMessages(context.Background(), msg_batch...)
    //if err != nil {
        //fmt.Println(err)
    //}
}


func getKafkaReader(kafka_brokers []string, topic, groupID string) *kafka.Reader {
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers:  kafka_brokers,
        GroupID:  groupID,
        Topic:    topic,
        MinBytes: 10e3, // 10KB
        MaxBytes: 10e6, // 10MB
    })
}
