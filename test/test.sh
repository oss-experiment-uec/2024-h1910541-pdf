#!/bin/bash

go run generate_text.go
if [ $? -ne 0 ]; then
    echo "Error: Failed to generate sample.txt"
    exit 1
fi

enscript sample.txt --no-header -o - | ps2pdf - sample.pdf
if [ $? -ne 0 ]; then
    echo "Error: Failed to convert sample.txt to sample.pdf"
    exit 1
fi

go run main.go sample.pdf
if [ $? -ne 0 ]; then
    echo "Error: Failed to process sample.pdf"
    exit 1
fi

