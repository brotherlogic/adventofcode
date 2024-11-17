FROM ruby:3.0.1

WORKDIR /solver

COPY Gemfile /solver
COPY Gemfile.lock /solver
COPY server.rb /solver

RUN bundle config --global frozen 1
RUN bundle install --without development test

COPY lib /solver/lib
EXPOSE 8080

ENTRYPOINT [ "bundle", "exec"]
CMD ["ruby", "server.rb"]