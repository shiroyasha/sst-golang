![Semaphore CLI](https://raw.githubusercontent.com/shiroyasha/sst-golang/master/logo.png)

## Installation

Before installation:

``` bash
mkdir -p ~/.sst 

echo "<api_token>" > ~/.sst/api_token
echo "<api_domain>" > ~/.sst/api_domain
```

``` bash
mkdir -p ~/bin

wget "https://github.com/shiroyasha/sst-golang/releases/download/0.1.1/semaphore" -O ~/bin/semaphore

chmod +x ~/bin/semaphore
```

Optionally append the following line to your `~/.bash_rc` or `~/.zshrc`:

``` bash
export PATH="$HOME/bin:$PATH"
```
