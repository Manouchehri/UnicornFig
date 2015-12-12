if [ ! -d "$HOME/.vim" ]; then
    mkdir $HOME/.vim
fi

if [ ! -d "$HOME/.vim/ftdetect" ]; then
    mkdir $HOME/.vim/ftdetect
fi

if [ ! -d "$HOME/.vim/ftplugin" ]; then
    mkdir $HOME/.vim/ftplugin
fi

if [ ! -d "$HOME/.vim/syntax" ]; then
    mkdir $HOME/.vim/syntax
fi

cp ./ftdetect/fig.vim $HOME/.vim/ftdetect/fig.vim
cp ./ftplugin/fig.vim $HOME/.vim/ftplugin/fig.vim
cp ./syntax/fig.vim $HOME/.vim/syntax/fig.vim
