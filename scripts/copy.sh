#!/bin/bash

# List down contents of ./bin folder and ask user to choose file to copy into /mnt/rp2040
echo "Available files to copy into device:"
ls -al ./bin
read -p "Enter the name of the file to copy into device: " FILE_NAME
if [ -f "./bin/$FILE_NAME" ]; then
    echo "Copying file: $FILE_NAME into device..."
    sudo cp "./bin/$FILE_NAME" /mnt/rp2040/
    echo "File copied successfully!"
else
    echo "File not found: $FILE_NAME"
fi