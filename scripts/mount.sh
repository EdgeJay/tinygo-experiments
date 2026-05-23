#!/bin/bash

DEVICE=$(dmesg | tail -50 | grep -Po "^(?:\[\s*\d+\.\d+\])\s\s(?:sd[a-z]\:\s)\K(sd[a-z]\d)")
echo "Mounting device: /dev/$DEVICE"
sudo mount /dev/$DEVICE /mnt/rp2040
