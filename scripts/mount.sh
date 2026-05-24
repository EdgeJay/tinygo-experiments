#!/bin/bash

unset DEVICE
DEVICE=$(dmesg | tail -30 | grep -Po "^(?:\[\s*\d+\.\d+\])\s\s(?:sd[a-z]\:\s)\K(sd[a-z]\d)")
echo "Mounting device: /dev/$DEVICE"
sudo mount /dev/$DEVICE /mnt/rp2040
