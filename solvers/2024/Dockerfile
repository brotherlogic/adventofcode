FROM ruby:3.3.3-bookworm

RUN mkdir /solver
WORKDIR /solver

COPY Gemfile /solver
COPY Gemfile.lock /solver

RUN mkdir /solver/lib

COPY lib /solver/lib
COPY server.rb /solver
COPY *.rb /solver

RUN bundle config --global 
RUN bundle install --without development test

EXPOSE 8080
CMD ["ruby", "server.rb"]