package main

import (
	"bytes"
	chunker "fileTrial/proto"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Handler []byte

func (h Handler) Uploader(stream chunker.FileService_UploaderServer) error {
	// Open a new file for writing only
	file, err := os.OpenFile("dummy", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err = file.Write(c.Chunk)
		if err != nil {
			return err
		}
	}
}

const chunkSize = 64 * 1024 // 64 KiB

func (h Handler)Downloader( meta *chunker.FileMeta, stream chunker.FileService_DownloaderServer) (err error) {
	//Getting file
	resp, err := http.Get(meta.GetFilePath())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fileName := buildFileName(meta.GetFilePath())
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	defer os.Remove(fileName)
	_, err = io.Copy(out, resp.Body)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(file)
	var line []byte
	for {
		line, err = buf.ReadBytes('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}
		if err = stream.Send(&chunker.Chunk{Chunk:line}); err != nil {
			return err
		}
	}
}

func buildFileName(fullFileUrl string) string {
	fileUrl, err := url.Parse(fullFileUrl)
	if err != nil{
		return "dummy"
	}
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

func StreamingFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	blob := make([]byte, chunkSize)
	chunker.RegisterFileServiceServer(g, Handler(blob))
	log.Println("Serving on :5555")
	log.Fatalln(g.Serve(lis))
}
