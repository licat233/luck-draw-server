#!/bin/bash
loading_show=""
loading_index=0
loading_ch=('|' '\' '-' '/')
loading() {
    i=$loading_index
    j=$1
    if [ -z $j ]; then
        j=100
    fi
    str=$loading_show
    while [ $i -le $j ]; do
        let index=i%4
        if [ $i -le 20 ]; then
            let color=44
            let bg=34
        elif [ $i -le 45 ]; then
            let color=43
            let bg=33
        elif [ $i -le 75 ]; then
            let color=41
            let bg=31
        else
            let color=42
            let bg=32
        fi
        printf "\033[${color};${bg}m%-s\033[0m %d %c\r" "$str" "$i" "%" "${loading_ch[$index]}"
        usleep 30000
        let i=i+1
        if [ $(($i % 2)) -eq 0 ]; then
            str+="#"
        fi

    done
    loading_index=$i
    loading_show=$str
    if (($loading_index > 100)); then
        printf "\n"
        loading_index=0
        loading_show=""
    fi
    unset str
}
loading 100