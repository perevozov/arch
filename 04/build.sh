#!/bin/sh

docker build -t perevozov/arch:userservice user/

docker build -t perevozov/arch:emailchecker emailchecker/
