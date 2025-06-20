package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/go-connections/nat"
)

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(fmt.Errorf("failed to create Docker client: %v", err))
	}
	defer cli.Close()

	imageName := "nginx:latest"
	containerName := "my-go-nginx-container"
	hostPort := "8080"
	containerPort := "80"

	fmt.Printf("--- Pulling image: %s ---\n", imageName)
	if err := pullImage(ctx, cli, imageName); err != nil {
		panic(fmt.Errorf("failed to pull image: %v", err))
	}

	fmt.Printf("--- Checking for existing container: %s ---\n", containerName)
	if err := removeExistingContainer(ctx, cli, containerName); err != nil {
		panic(fmt.Errorf("failed to remove existing container: %v", err))
	}

	fmt.Printf("--- Creating container: %s ---\n", containerName)
	containerID, err := createContainer(ctx, cli, imageName, containerName, hostPort, containerPort)
	if err != nil {
		panic(fmt.Errorf("failed to create container: %v", err))
	}

	fmt.Printf("--- Starting container: %s ---\n", containerName)
	if err := cli.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		panic(fmt.Errorf("failed to start container: %v", err))
	}

	fmt.Printf("Container %s started. Access it at http://localhost:%s\n", containerName, hostPort)

	fmt.Println("Press Enter to stop and remove the container...")
	fmt.Scanln()

	if err := stopAndRemoveContainer(ctx, cli, containerID); err != nil {
		panic(fmt.Errorf("cleanup failed: %v", err))
	}
}

// Pull Docker image
func pullImage(ctx context.Context, cli *client.Client, imageName string) error {
	reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		images, err := cli.ImageList(ctx, image.ListOptions{
			Filters: filters.NewArgs(filters.Arg("reference", imageName)),
		})
		if err != nil {
			return fmt.Errorf("failed to check local images: %v", err)
		}
		if len(images) == 0 {
			return fmt.Errorf("image %s not found locally and pull failed: %v", imageName, err)
		}
		fmt.Printf("Using existing local image: %s\n", imageName)
		return nil
	}
	defer reader.Close()
	if err := jsonmessage.DisplayJSONMessagesStream(reader, os.Stdout, os.Stdout.Fd(), true, nil); err != nil {
		return fmt.Errorf("failed to display pull progress: %v", err)
	}
	fmt.Println("Image pull complete.")
	return nil
}

// Remove existing container if it exists
func removeExistingContainer(ctx context.Context, cli *client.Client, containerName string) error {
	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", containerName)),
	})
	if err != nil {
		return fmt.Errorf("failed to list containers: %v", err)
	}
	for _, c := range containers {
		if err := cli.ContainerRemove(ctx, c.ID, container.RemoveOptions{
			Force: true,
		}); err != nil {
			return fmt.Errorf("failed to remove container %s: %v", c.ID, err)
		}
		fmt.Printf("Removed existing container: %s\n", c.ID)
	}
	return nil
}

// Create container
func createContainer(ctx context.Context, cli *client.Client, imageName, containerName, hostPort, containerPort string) (string, error) {
	portNat, err := nat.NewPort("tcp", containerPort)
	if err != nil {
		return "", fmt.Errorf("invalid container port: %v", err)
	}
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imageName,
			ExposedPorts: nat.PortSet{
				portNat: struct{}{},
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				portNat: []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: hostPort,
					},
				},
			},
		},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %v", err)
	}
	fmt.Printf("Container created with ID: %s\n", resp.ID)
	return resp.ID, nil
}

// Stop and remove container
func stopAndRemoveContainer(ctx context.Context, cli *client.Client, containerID string) error {
	fmt.Printf("--- Stopping container: %s ---\n", containerID)
	if err := cli.ContainerStop(ctx, containerID, container.StopOptions{
		Timeout: nil,
	}); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}
	fmt.Printf("--- Removing container: %s ---\n", containerID)
	if err := cli.ContainerRemove(ctx, containerID, container.RemoveOptions{}); err != nil {
		return fmt.Errorf("failed to remove container: %v", err)
	}
	fmt.Println("Container stopped and removed successfully.")
	return nil
}
