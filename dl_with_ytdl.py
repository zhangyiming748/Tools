def download(url):
    cmd = 'youtube-dl --proxy 127.0.0.1:1080 --write-sub --all-subs -i -f bestvideo+bestaudio ' + url
