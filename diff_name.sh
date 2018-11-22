#!/bin/sh

output=$(diff -qr $1 $2)
array=()
while read -r line; do
  string=$(echo "$line" | awk '{print $1}');
  if [ $(echo $string | grep -e 'Only') ]; then
    path=$(echo "$line" | awk '{print $3}' | sed -e "s/\:/\//");
    fileName=$(echo "$line" | awk '{print $4}');
    # .DS_SToreはMacのみなので、必要がなければ削除してください
    if [ $fileName != '.DS_Store' ]; then
      echo ${path/$1/$2}$fileName;
      array+=(${path/$1/$2}$fileName);
    fi
  else
    echo $(echo "$line" | awk '{print $4}');
    array+=($(echo "$line" | awk '{print $4}'));
  fi
done <<< "$output"
