" ---- Basic settings ----
set number
set autoindent
set tabstop=4
set smartindent
set encoding=utf-8 	" file encoding
set nocompatible
set shiftwidth=4

" ---- colorscheme ----
colorscheme habamax

" ---- syntax highlighting ----
syntax on

" ---- Autocompletion[ycm] ----
" let g:ycm_global_extra_conf = '~/.vim/.ycm_extra_conf.py'
set wildmenu
set wildmode=list:longest

" ---- filetype-specific settings ----
augroup filetype_settings
	autocmd!

	" Py
	autocmd FileType python setlocal tabstop=4 shiftwidth=4 softtabstop=4
	autocmd FileType python setlocal autoindent

	" js
	autocmd FileType javascript setlocal tabstop=2 shitwidth=2 shiftwidth=2 softtabstop=2
	autocmd FileType javascript setlocal expandtab
	autocmd FileType javascript setlocal autoindent
	autocmd FileType javascript setlocal smartindent
	" c#
	autocmd FileType cs setlocal smartindent
	autocmd FileType cs setlocal relativenumber
	
	" go
	autocmd FileType go setlocal relativenumber
	 
	" space for any other languages I will care about later rust & haskell
	
augroup END

" ---- status line info ----

set laststatus=2
set statusline+=%f\ 	" for filename
set statusline+=\ %y\ 	" filetype display
set statusline+=\ %{&ff}\ 	" ff file format
set statusline+=\ %l,%c\ 	" display line and column number
set statusline+=\ %p%%\ 	" show percentage through file

" ---- plugins ----
call plug#begin()

Plug 'vim-airline/vim-airline'
Plug 'scrooloose/nerdtree'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'ycm-core/YouCompleteMe'
Plug 'rdnetto/YCM-Generator', {'branch': 'stable'}

call plug#end()

let g:ycm_global_ycm_extra_conf = '~/.vim/.ycm_extra_conf.py'
let g:ycm_confirm_extra_conf = 0
nnoremap <leader>g :YcmCompleter GoToDefinitionElseDeclaration<CR>

" ---- key-mappings ----
nnoremap <C-n> :NERDTreeToggle<CR>
nnoremap <C-t> :YcmCompleter GoToDefinitionElseDeclaration<CR>
