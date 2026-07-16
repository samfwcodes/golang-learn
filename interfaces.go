package main

import (
	"bytes"
	"fmt"
)
// USE OF INTERFACES FOR WRITING DATA INTO CHUNKS
// THE PROBLEM : "I want to write data, but I don't want it printed immediately. I want it stored in memory, and whenever there are 8 characters available, print those 8 characters. When I'm done, print whatever is left."

func main() {
	var wc WriterCloser = NewBufferedWriterCloser() // Create a new BufferedWriterCloser and use it through the WriterCloser interface.WriterCloser interface.
	wc.Write([]byte("This string will be converted into 8 word characters for each line.")) // Write data into the internal buffer.
	wc.Close()  // Flush and print any remaining data in the buffer.
}

type Writer interface {
	Write([]byte) (int, error) // Any type that implements Write([]byte) (int, error) satisfies the Writer interface.
}
type Closer interface {
	Close() error // Any type that implements Close() error satisfies the Closer interface.
}

type WriterCloser interface {
	Writer
	Closer
} // Embedding Writer and Closer interface into one

type BufferedWriterCloser struct {
	buffer *bytes.Buffer // Pointer to a bytes.Buffer used to temporarily store data before it is processed.
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) { // Write appends data to the buffer and prints it in chunks of 8 bytes.
	n, err := bwc.buffer.Write(data) // Write the incoming data into the internal buffer.
	if err != nil {
		return 0, err // error handling
	}
	v := make([]byte, 8) // Allocate a byte slice of length 8.
	// The buffer will copy 8 bytes into this slice when Read() is called.
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v) // Keep reading while there are MORE THAN 8 bytes in the buffer.
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v)) // Convert the byte slice to a string and print it.
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error { // Closer method on BufferedWriterCloser
	for bwc.buffer.Len() > 0 { // Continue until the buffer becomes empty.
		data := bwc.buffer.Next(8) // Remove and return up to the next 8 bytes from the buffer.
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}), // Constructor that creates and returns a new BufferedWriterCloser with an empty buffer.
	}
}
