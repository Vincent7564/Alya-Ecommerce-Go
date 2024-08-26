package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func load_env() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	load_env()

	URL := os.Getenv("NEXT_PUBLIC_SUPABASE_URL")
	ANON := os.Getenv("NEXT_PUBLIC_SUPABASE_ANON")

	client, err := supabase.NewClient(URL, ANON, &supabase.ClientOptions{})
	fmt.Println(err)
	data, datas, error := client.From("users").Select("*", "exact", false).Execute()
	fmt.Println("Data ", data)
	fmt.Println("Datas ", datas)
	fmt.Println("Error ", error)
}
