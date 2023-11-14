package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "justech.json")
	ctx := context.Background()
	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	// Sets the name for the new bucket.
	bucketName := "justech"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)
	// escrita
	w := bucket.Object("test/xpto.json").NewWriter(ctx)
	f, err := os.Open("justech.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
	}
	w.Close()
	// letura
	reader, err := bucket.Object("test/xpto.json").NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create reader: %v", err)
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}
	reader.Close()
	fmt.Fprintln(os.Stdout, string(b))
}
