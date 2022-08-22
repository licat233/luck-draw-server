#!/bin/bash
goctl model mysql datasource -url="luck_draw:123123123@tcp(127.0.0.1:3306)/luck_draw" -table="*"  -dir='.' -cache=true -style goZero