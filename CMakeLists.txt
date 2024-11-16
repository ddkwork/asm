cmake_minimum_required(VERSION 3.28)
project(asm C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_FIND_STATIC_PREFERENCES TRUE)
set(BUILD_SHARED_LIBS OFF)


add_executable(asm main.c)
