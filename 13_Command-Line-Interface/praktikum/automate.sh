#!/bin/sh

dir="$1 at $(date)"

mkdir "$dir"
mkdir "$dir/about_me"
mkdir "$dir/about_me/personal"
mkdir "$dir/about_me/professional"
mkdir "$dir/my_friends"
mkdir "$dir/my_system_info"

echo "$2" > "$dir/about_me/personal/facebook.txt"
echo "$3" > "$dir/about_me/professional/linkedin.txt"
curl  "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt/" -o "$dir/my_friends/list_of_my_friends.txt"
uname -a > "$dir/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$dir/my_system_info/internet_connection.txt"