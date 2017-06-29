tweet-picker
====
Pick up deleted tweet

## Overview

TLからツイ消しを拾ってくるやつ  
ツイ消しを眺めたいときにどうぞ

## Usage

1. clone this repository

    ```console
    $ git clone https://github.com/bgpat/tweet-picker
    $ cd tweet-picker
    ```
  
2. copy sample `.env` and edit

    ```console
    $ cp .env.sample .env
    $ vi .env
    ```

3. copy sample `docker-compose.yml` and edit

    ```console
    $ cp docker-compose.yml.sample docker-compose.yml
    $ vi docker-compose.yml
    ```
    
4. execute `docker-compose up`

    ```console
    $ docker-compose up -d
    $ docker-compose ps
           Name                      Command               State           Ports          
    -------------------------------------------------------------------------------------
    tweetpicker_api_1     /server                          Up                             
    tweetpicker_cache_1   docker-entrypoint.sh redis ...   Up      6379/tcp               
    tweetpicker_db_1      docker-entrypoint.sh postgres    Up      5432/tcp               
    tweetpicker_web_1     nginx -g daemon off;             Up      0.0.0.0:8080->80/tcp 
    ```
    
5. have fun !!!

    ```console
    $ open http://localhost:8080
    ```
