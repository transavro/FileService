package main

import (
	"bytes"
	chunker "fileTrial/proto"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

func buildFileName(fullFileUrl string) string {
	fileUrl, err := url.Parse(fullFileUrl)
	if err != nil {
		return "dummy"
	}
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

var (
	downloadUrl *string
	uploadPath  *string
)

func upload(cc chunker.FileServiceClient) {
	stream, err := cc.Uploader(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	file, err := ioutil.ReadFile(*uploadPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.CloseAndRecv()
	buf := bytes.NewBuffer(file)
	var line []byte
	for {
		line, err = buf.ReadBytes('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					log.Println("Uploaded Successfull")
					break
				}
				log.Fatalln(err)
			}
		}
		if err = stream.Send(&chunker.Chunk{Chunk: line}); err != nil {
			log.Fatalln(err)
		}
	}
}

func download(cc chunker.FileServiceClient) {
	client, err := cc.Downloader(context.Background(), &chunker.FileMeta{FilePath: *downloadUrl})
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(buildFileName(*downloadUrl), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	defer client.CloseSend()
	for {
		c, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Downloaded Successfull")
				return
			}
			log.Fatalln(err.Error())
		}
		_, err = file.Write(c.Chunk)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	downloadUrl = flag.String("download", "", "Give a download File Url to download a file.")
	uploadPath = flag.String("upload", "", "Give a file path to get upload on the server.")
	flag.Parse()

	conn, err := grpc.Dial(":5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cc := chunker.NewFileServiceClient(conn)

	if len(*downloadUrl) > 0 {
		download(cc)
	} else if len(*uploadPath) > 0 {
		upload(cc)
	} else {
		log.Fatal("Please provide upload or download flag.")
	}
}
