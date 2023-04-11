# 26_Docker

## Docker
Docker berbasis container yang mana memiliki pengertian container adalah bukan sebuah virtual machine, karena container ini berorientasi kepada service yang ingin dibuat. 

## Docker Infrastructure
- Client -> User
- Docker Host -> Untuk menyimpan containerer yang menjalankan suatu image. Aktifitas ini dilakukan oleh docker daemon
- Registry -> Untuk menyimpan docker image yang mana docker image ini nantinya bisa digunakan oleh developer lain.

## Apa yang bisa docker lakukan?
- Dapat membuat image sendiri (Dockerfile). Dalam Dockerfile terdapat beberapa syntax yang akan dijalankan ketika kita membuat image.
- Mengelompokkan beberapa service atau image dalam suatu container

## Docker Syntax
- FROM -> mendapatkan docker image dari registry
- RUN -> eksekusi bash command ketika build container
- ENV -> set env variable
- AND
- COPY -> melakukan copy
- WORKDIR -> workdir dari image yang di build
- ENTRYPOINT
- CMD -> command untuk menjalankan sesuatu