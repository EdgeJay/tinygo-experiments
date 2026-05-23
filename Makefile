mount-device:
	@chmod +x ./scripts/mount.sh && ./scripts/mount.sh

build-calculator:
	@echo "Building calculator..."
	@tinygo build -o ./bin/01_calculator.uf2 --target waveshare-rp2040-zero --size short ./01_calculator/

copy-into-device:
	@chmod +x ./scripts/copy.sh && ./scripts/copy.sh