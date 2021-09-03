FROM gitpod/workspace-full

# Install zsh
RUN apt-get update && apt-get install -y zsh

# Install .dotfiles
RUN git clone github.com/ngoldack/.dotfiles && cd .dotfiles && ./install

# Install gems
RUN gem install colorls
