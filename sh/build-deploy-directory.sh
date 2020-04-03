#!/bin/bash

echo 'building /deploy directory'

sh/build-linux-executable.sh

echo "чистим ./deploy/"

rm -rf deploy/message-scheduler
rm -rf deploy/certificates



echo "копируем файлы в ./deploy/"

mv message-scheduler   deploy/message-scheduler
cp -r certificates      deploy/certificates


