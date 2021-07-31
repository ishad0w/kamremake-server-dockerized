#!/bin/bash

for VAR in $(awk -F "=" '/\[.*\]/{found=0} {if(found) print $1} /\[Server\]/{found=1}' "$KAM_CONFIG" | xargs); do
    if [ -n "${!VAR}" ]
        then
            printf "$VAR variable is present. Overriding $VAR with: ${!VAR}.\n"
            sed -i 's%'"$VAR"'=.*%'"$VAR"'='"${!VAR}"'%g' "$KAM_CONFIG"
        else
            printf "$VAR variable is not present. Using default.\n"
    fi
done

printf "\nReady! \nVersion = $KAM_VERSION, Build Type = $KAM_BUILD, MD5 = $(md5sum "/kam/server-${KAM_BUILD}-$KAM_VERSION" | cut -d ' ' -f 1)\n"

/kam/metrics.sh &
/kam/server-${KAM_BUILD}-$KAM_VERSION
