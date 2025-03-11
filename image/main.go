package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {

    new_png_file := "chessboard.png" // output image lives here
    board_num_pixels := 240

    myimage := image.NewRGBA(image.Rect(0, 0, board_num_pixels, board_num_pixels))
    colors := make(map[int]color.RGBA, 2)

    colors[0] = color.RGBA{0, 100, 0, 255}   // green
    colors[1] = color.RGBA{50, 205, 50, 255} // limegreen

    index_color := 0
    size_board := 8
    size_block := int(board_num_pixels / size_board)
    loc_x := 0

    for curr_x := 0; curr_x < size_board; curr_x++ {

        loc_y := 0
        for curr_y := 0; curr_y < size_board; curr_y++ {

            draw.Draw(myimage, image.Rect(loc_x, loc_y, loc_x+size_block, loc_y+size_block),
                &image.Uniform{colors[index_color]}, image.ZP, draw.Src)

            loc_y += size_block
            index_color = 1 - index_color // toggle from 0 to 1 to 0 to 1 to ...
        }
        loc_x += size_block
        index_color = 1 - index_color // toggle from 0 to 1 to 0 to 1 to ...
    }
    myfile, err := os.Create(new_png_file) 
    if err != nil {
        panic(err.Error())
    }
    defer myfile.Close()
    png.Encode(myfile, myimage) // ... save image
    fmt.Println("firefox ", new_png_file) // view image issue : firefox  /tmp/chessboard.png
}