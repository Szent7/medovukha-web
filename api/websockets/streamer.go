package websockets

import (
	"context"
	"io"
	"log"
	"strings"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *Hub) ContainerEventStreamer(ctx context.Context, broadcast chan<- []byte, _ string) error {
	eventCh := make(chan *dockerpb.GetContainerStateResponse)
	errCh := make(chan error)
	go h.api.GetContainerState(ctx, eventCh, errCh)

	// log.Println("ContainerEventStreamer: start")

	for {
		select {
		case <-ctx.Done():
			{
				// log.Println("ContainerEventStreamer: start")
				return nil
			}
		case event := <-eventCh:
			{
				jsonBytes, err := protojson.Marshal(event)
				if err != nil {
					log.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err != nil {
					log.Printf("ContainerEventStreamer error: %s\n", err.Error())
					return err
				}
			}
		}
	}
}

func (h *Hub) ImageEventStreamer(ctx context.Context, broadcast chan<- []byte, _ string) error {
	eventCh := make(chan *dockerpb.GetImageStateResponse)
	errCh := make(chan error)
	go h.api.GetImageState(ctx, eventCh, errCh)

	// log.Println("ImageEventStreamer: start")

	for {
		select {
		case <-ctx.Done():
			{
				// log.Println("ImageEventStreamer: start")
				return nil
			}
		case event := <-eventCh:
			{
				jsonBytes, err := protojson.Marshal(event)
				if err != nil {
					log.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err != nil {
					log.Printf("ImageEventStreamer error: %s\n", err.Error())
					return err
				}
			}
		}
	}
}

func (h *Hub) NetworkEventStreamer(ctx context.Context, broadcast chan<- []byte, _ string) error {
	eventCh := make(chan *dockerpb.GetNetworkStateResponse)
	errCh := make(chan error)
	go h.api.GetNetworkState(ctx, eventCh, errCh)

	// log.Println("NetworkEventStreamer: start")

	for {
		select {
		case <-ctx.Done():
			{
				// log.Println("NetworkEventStreamer: start")
				return nil
			}
		case event := <-eventCh:
			{
				jsonBytes, err := protojson.Marshal(event)
				if err != nil {
					log.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err != nil {
					log.Printf("NetworkEventStreamer error: %s\n", err.Error())
					return err
				}
			}
		}
	}
}

func (h *Hub) VolumeEventStreamer(ctx context.Context, broadcast chan<- []byte, _ string) error {
	eventCh := make(chan *dockerpb.GetVolumeStateResponse)
	errCh := make(chan error)
	go h.api.GetVolumeState(ctx, eventCh, errCh)

	// log.Println("VolumeEventStreamer: start")

	for {
		select {
		case <-ctx.Done():
			{
				// log.Println("VolumeEventStreamer: start")
				return nil
			}
		case event := <-eventCh:
			{
				jsonBytes, err := protojson.Marshal(event)
				if err != nil {
					log.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err != nil {
					log.Printf("VolumeEventStreamer error: %s\n", err.Error())
					return err
				}
			}
		}
	}
}

func (h *Hub) BuildLogStreamer(ctx context.Context, broadcast chan<- []byte, buildID string) error {
	logCh := make(chan *dockerpb.StreamBuildLogsResponse)
	errCh := make(chan error)
	res := strings.TrimPrefix(buildID, "build-container-")
	go h.api.StreamBuildLogs(ctx, res, logCh, errCh)

	// log.Println("BuildLogStreamer: start")

	for {
		select {
		case <-ctx.Done():
			{
				// log.Println("BuildLogStreamer: start")
				return nil
			}
		case line := <-logCh:
			{
				jsonBytes, err := protojson.Marshal(line)
				if err != nil {
					log.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				log.Printf("Received: %v", line.Line)
				broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err == io.EOF {
					return nil
				}
				if err != nil {
					log.Printf("BuildLogStreamer error: %s\n", err.Error())
					return err
				}
			}
		}
	}
}
