###go环境配置
1.安装gvm，vim的包管理器
	[vim的包管理器github](https://github.com/moovweb/gvm)
通过homebrew安装
<br>
2.
```
gvm install go1.4
gvm use go1.4
//export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.8
```
<br>
2.gvm install go1.8 --binary #直接安裝 binary

###go vim开发环境配置
安装vim插件管理器[Vundle.vim](https://github.com/VundleVim/Vundle.vim) (git clone)<br>
配置/usr/share/vim/vimrc文件，配置安装插件<br>
安装[vim－go](https://github.com/fatih/vim-go)