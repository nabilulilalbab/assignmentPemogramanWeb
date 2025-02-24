package main

import (
	"bufio"
	"fmt"
	"os"
)

func SayHello(name string)  {
  fmt.Print("Enter your name : ")
  // fmt.Scan(&name) Hanya bisa input data tanpa spasi
  scanner := bufio.NewScanner(os.Stdin) // Membuat instance baru baru Sebagai scanner, Scanner Read ( stdin, file) , flow : memecah menjadi potongan-potongan (token) kecil, Membuat buffer untuk input -> NewScanner : Scanner Membuat buffer untuk membuat data yang akan di baca, Bufffer ini berfungsi menyimpan potongan data sementara dari sumber input(stdin / file),memungkinkan membaca banyak data sekaligus, dari pada scanner bawaan fmt -> os.stdin (input dari keyboard)  
  scanner.Scan() // membaca token berikutnya dari input (satu baris / kata), jika input berisi beberapa kata atau spasi, tokenisasi akan terjadi di dalam buffer , dan hasilnya akan di pisahkan sesuai dengan delimiter(spasi atau newline)
  name = scanner.Text() // Mendapatkan string dari hasil generat token
  fmt.Printf("Halo %v!, Selamat belajar Golang", name)

}

func main()  {
  SayHello("nabiel")
}
