package pipeline

import (
    "os"
    "context"
    "github.com/ThomasMTurner/BookBuddy/converter"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "github.com/fsnotify/fsnotify"

)

func enterNewBookEntry(bookPath string) {
    // Build Book structs for each PDF file in the directory.
    builder := &convert.PDFBuilder{}
    file, err := os.Open(bookPath) 
    if err != nil {
        return 
    }

    defer file.Close() 

    result, err := builder.Build(file)
    if err != nil {
        return 
    }

    // Map Book structs to BSON documents to insert into MongoDB instance at port number 27017.
    client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

    collection := client.Database("library").Collection("books")
    
    _, err = collection.InsertOne(context.TODO(), result)
    if err != nil {
        log.Fatal(err)
    }
    
}


func ListenOnNewBookEntries() {
    // Ensure existing local MongoDB database on generated port number.
    

    // File watchdog to check for new book entries in directory.
    // Create a new watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // Specify the directory to watch
    dirToWatch := "./books"

    // Add the directory to the watcher
    err = watcher.Add(dirToWatch)
    if err != nil {
        log.Fatal(err)
    }

    // Create a channel to receive events
    done := make(chan bool)

    // Start a goroutine to listen for events
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Create == fsnotify.Create {
                    enterNewBookEntry(event.Name)

                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Wait for the user to stop the program - can limit with timeout, currently indefinite.
    <-done

}
