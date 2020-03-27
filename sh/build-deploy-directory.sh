#!/bin/bash

echo 'building /deploy directory'

sh/build-executable.sh

echo "чистим ./deploy/"

rm -rf deploy/message-scheduler
rm -rf deploy/configs_example
rm -rf deploy/certificates



echo "копируем файлы в ./deploy/"

mv message-scheduler   deploy/message-scheduler
cp -r configs           deploy/configs_example
cp -r certificates      deploy/certificates



echo "осторожно копируем файлы в ./deploy/configs/ "

mkdir -p deploy/configs
cp -f configs/configs.yaml  deploy/configs/configs.yaml
cp -f configs/signature.yaml  deploy/configs/signature.yaml
